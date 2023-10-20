/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/lhps/desafio-01/application/grpc"
	"github.com/lhps/desafio-01/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start a gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB()
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC server port")
}
