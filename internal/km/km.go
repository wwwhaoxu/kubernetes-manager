package km

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewKuberManagerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "ku",
		Short:        "A good kubernetes manager project",
		Long:         "A good kubernetes manager project",
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
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
	return cmd
}

// run 函数是实际的业务代码入口函数.
func run() error {
	fmt.Println("Hello World!")
	return nil
}
