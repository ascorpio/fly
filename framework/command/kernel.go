package command

import "github.com/ascorpio/fly/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// app 命令
	root.AddCommand(initAppCommand())
	// env 命令
	root.AddCommand(initEnvCommand())
	// cron 命令
	root.AddCommand(initCronCommand())
	// config 命令
	root.AddCommand(initConfigCommand())
	// build 命令
	root.AddCommand(initBuildCommand())
	// go build
	root.AddCommand(goCommand)
	// npm build
	root.AddCommand(npmCommand)
	//root.AddCommand(deployCommand)
	//
	// cron
	//// cmd
	//cmdCommand.AddCommand(cmdListCommand)
	//cmdCommand.AddCommand(cmdCreateCommand)
	//root.AddCommand(cmdCommand)
	//
	//
	//// dev
	//root.AddCommand(initDevCommand())
	//
	//// middleware
	//middlewareCommand.AddCommand(middlewareAllCommand)
	//middlewareCommand.AddCommand(middlewareAddCommand)
	//middlewareCommand.AddCommand(middlewareRemoveCommand)
	//root.AddCommand(middlewareCommand)
	//
	//// swagger
	//swagger.IndexCommand.AddCommand(swagger.InitServeCommand())
	//swagger.IndexCommand.AddCommand(swagger.GenCommand)
	//root.AddCommand(swagger.IndexCommand)
	//
	//// provider
	//providerCommand.AddCommand(providerListCommand)
	//providerCommand.AddCommand(providerCreateCommand)
	//root.AddCommand(providerCommand)
	//
	//// new
	//root.AddCommand(initNewCommand())
}
