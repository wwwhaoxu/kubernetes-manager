package log

import "go.uber.org/zap/zapcore"

type Options struct {
	// 是否开启 caller，如果开启会在日志中显示调用日志所在的文件和行号
	DisableCaller bool
	// 是否禁止在 panic 及以上级别打印堆栈信息
	DisableStacktrace bool
	// 指定日志级别，可选值：debug, info, warn, error, dpanic, panic, fatal
	Level string
	// 指定日志显示格式，可选值：console, json
	Format string
	// 指定日志输出位置
	OutputPaths []string
}

type Option func(options *Options)

// NewOptions 创建一个带有默认参数的 Options 对象.
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}

func DisableCaller(b bool) Option {
	return func(opts *Options) {
		opts.DisableCaller = b
	}
}

func DisableStacktrace(b bool) Option {
	return func(opts *Options) {
		opts.DisableStacktrace = b
	}
}
