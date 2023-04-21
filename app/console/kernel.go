package console

import (
	"github.com/gohade/my-web/app/console/command/demo"
	"github.com/gohade/my-web/framework"
	"github.com/gohade/my-web/framework/cobra"
	"github.com/gohade/my-web/framework/command"
)

// AppAddCommand 绑定业务的命令
func AppAddCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}

// RunCommand  初始化根Command并运行
func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "hade",
		Short: "hade 命令",
		Long:  "提供各种命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCmd.SetContainer(container)
	command.AddKernelCommands(rootCmd)
	AppAddCommand(rootCmd)
	return nil
}
