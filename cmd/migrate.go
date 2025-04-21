/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"medods/database/migration"
	"medods/logger"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(logrus.Level(viper.GetInt("logger.level")))

		options := pg.Options{
			Addr:     viper.GetString("db.address"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Database: viper.GetString("db.database"),
		}

		if err := migration.Up(options, migration.LastVersion); err != nil {
			logger.Log().Errorf("error on up migrate - %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
