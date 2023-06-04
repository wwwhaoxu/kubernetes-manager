package km

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"kubernetes-manager/internal/km/controller/v1/user"
	"kubernetes-manager/internal/km/store"
	"kubernetes-manager/internal/pkg/known"
	"kubernetes-manager/internal/pkg/log"
	mw "kubernetes-manager/internal/pkg/middleware"
	pb "kubernetes-manager/pkg/proto/km/v1"
	"kubernetes-manager/pkg/token"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfgFile string

func NewKuberManagerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "ku",
		Short:        "A good kubernetes manager project",
		Long:         "A good kubernetes manager project",
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 初始化日志
			log.Init(logOptions())
			// Sync 将缓存中的日志刷新到磁盘文件中
			defer log.Sync()
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	// 以下设置，使得 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the kuber-manager configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {

	// 初始化 store 层
	if err := initStore(); err != nil {
		return err
	}

	// 设置 token 包的签发密钥，用于 token 包 token 的签发和解析
	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)

	// 设置Gin模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建Gin引擎
	g := gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), mw.RequestID(), mw.Cors}
	g.Use(mws...)

	// 安装Routes
	if err := installRoutes(g); err != nil {
		return err
	}

	// 创建并运行 HTTP 服务器
	httpsrv := startInsecureServer(g)

	//// 创建并运行 HTTPS 服务器
	//httpssrv := startSecureServer(g)

	// 创建并运行 GRPC 服务器
	grpcsrv := startGRPCServer()

	// 等待中断信号优雅地关闭服务器（10 秒超时)。
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTTIN, syscall.SIGTERM)
	<-quit
	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
	}
	grpcsrv.GracefulStop()

	log.Infow("Server exiting")

	return nil
}

// startInsecureServer 创建并运行 HTTP 服务器.
func startInsecureServer(g *gin.Engine) *http.Server {
	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	// 运行 HTTP 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	return httpsrv
}

// startGRPCServer 创建并运行 GRPC 服务器.
func startGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatalw("Failed to listen", "err", err)
	}
	// 创建 GRPC Server 实例
	grpcsrv := grpc.NewServer()
	pb.RegisterKmServer(grpcsrv, user.New(store.S, nil))
	// 运行 GRPC 服务器。在 goroutine 中启动服务器，它不会阻止下面的正常关闭处理流程
	// 打印一条日志，用来提示 GRPC 服务已经起来，方便排障
	log.Infow("Start to listening the incoming requests on grpc address", "addr", viper.GetString("grpc.addr"))
	go func() {
		if err := grpcsrv.Serve(lis); err != nil {
			log.Fatalw(err.Error())
		}
	}()
	return grpcsrv
}
