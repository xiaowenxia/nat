package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xiaowenxia/nat/config"
)

type configCmd struct {
	cmd *cobra.Command
	O   struct {
		ConfigList bool
		ConfigPath bool
	}
}

func (v *configCmd) Command() *cobra.Command {
	if v.cmd != nil {
		return v.cmd
	}

	v.cmd = &cobra.Command{
		Use:           "config",
		Short:         "config --list",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return v.Execute(args)

		},
	}

	v.cmd.Flags().BoolVar(&v.O.ConfigList,
		"list",
		false,
		"show config list")

	v.cmd.Flags().BoolVar(&v.O.ConfigPath,
		"path",
		false,
		"show config path")
	return v.cmd
}

func (v configCmd) Execute(args []string) error {
	var seen = 0
	if v.O.ConfigList {
		seen++
	}

	if v.O.ConfigPath {
		seen++
	}

	if seen == 0 {
		v.cmd.Help()
		return errExecute
	}

	if v.O.ConfigList {
		return v.configList(args...)
	}

	if v.O.ConfigPath {
		return v.configPath(args...)
	}

	return nil
}

func (v configCmd) configList(args ...string) error {
	// 格式化 json 输出
	config, _ := json.MarshalIndent(config.GConfig, "", "  ")
	fmt.Println(string(config))
	return nil
}

func (v configCmd) configPath(args ...string) error {
	// 格式化 json 输出
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	return nil
}

var cc = configCmd{}

func init() {
	rootCmd.AddCommand(cc.Command())
}
