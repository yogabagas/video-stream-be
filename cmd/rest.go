package cmd

import (
	"github/yogabagas/video-stream-be/transport/rest"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use: "serve",
	PreRun: func(cmd *cobra.Command, args []string) {
		initModule()
	},
	Run: func(cmd *cobra.Command, args []string) {
		rest := rest.NewRest(&rest.Options{
			Port:         os.Getenv("APP_PORT"),
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		})

		rest.Serve()
	},
}
