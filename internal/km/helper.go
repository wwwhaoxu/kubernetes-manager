package km

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultConfigName  = "km.yaml"
	recommendedHomeDir = ".km"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserConfigDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		viper.AddConfigPath(".")

		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()
	// 读取环境变量的前缀为 KM，如果是 km，将自动转变为大写。
	viper.SetEnvPrefix("KM")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintln(os.Stdout, "Used config file:", viper.ConfigFileUsed())
}
