package main

import (
	"log"
	"os"
	"time"

	"load-tester/report"
	"load-tester/request"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var url string
var requests int
var concurrency int

func init() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado (obrigatório)")
	rootCmd.Flags().IntVar(&requests, "requests", 100, "Número total de requests")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")

	rootCmd.MarkFlagRequired("url")
}

var rootCmd = &cobra.Command{
	Use:   "load-tester",
	Short: "Ferramenta para testar carga em um serviço web",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			log.Fatal("A URL é obrigatória.")
		}

		// Inicia o teste de carga
		start := time.Now()
		results := request.RunLoadTest(url, requests, concurrency)
		duration := time.Since(start)

		// Gera o relatório
		report.GenerateReport(results, duration)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}
