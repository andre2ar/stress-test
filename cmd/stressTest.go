package cmd

import (
	"fmt"
	"github.com/andre2ar/stress-test/internal"
	"github.com/spf13/cobra"
	"sync"
)

// stressTestCmd represents the stressTest command
var stressTestCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Used to performa stress test on services.",
	Long:  `This stress test application can be used to perform stress test on web based services.`,
	Run: func(cmd *cobra.Command, args []string) {
		st := internal.StressTest{Report: sync.Map{}}

		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		fmt.Println("Processing requests, it might take a while...")
		st.Stress(url, requests, concurrency)
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(stressTestCmd)

	stressTestCmd.Flags().String("url", "", "URL which the performance test will be performed against.")
	_ = stressTestCmd.MarkFlagRequired("url")

	stressTestCmd.Flags().Int("requests", 1, "Number of requests.")
	_ = stressTestCmd.MarkFlagRequired("requests")

	stressTestCmd.Flags().Int("concurrency", 1, "How many simultaneous request will be triggered.")
	_ = stressTestCmd.MarkFlagRequired("concurrency")
}
