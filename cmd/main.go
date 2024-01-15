package main

import (
	"fmt"
	"log"

	"github.com/benmorehouse/traveler/config"
	"github.com/benmorehouse/traveler/pkg/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Traveler - API for travel data for users",
	RunE: func(cmd *cobra.Command, args []string) error {
		config := config.DefaultConfig()
		server := api.InitServer()
		return server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	},
}

func main() {
	// get a database connection here
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
