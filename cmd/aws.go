package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type AWSIPRanges struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPprefix string `json:"ip_prefix"`
		Region   string `json:"region"`
		Service  string `json:"service"`
	} `json:"prefixes"`
}

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Fetch AWS IP ranges",
	Run: func(cmd *cobra.Command, args []string) {
		fetchAWSIPRanges()
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)
}

func fetchAWSIPRanges() {
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching AWS IP ranges:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var ipRanges AWSIPRanges
	if err := json.Unmarshal(body, &ipRanges); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("AWS IP Ranges:")
	for _, prefix := range ipRanges.Prefixes {
		fmt.Printf("IP Prefix: %s, Region: %s, Service: %s\n", prefix.IPprefix, prefix.Region, prefix.Service)
	}
}
