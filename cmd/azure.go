package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type AzureIPRanges struct {
	Values []struct {
		Name       string `json:"name"`
		Properties struct {
			AddressPrefixes []string `json:"addressPrefixes"`
			Region          string   `json:"region"`
		} `json:"properties"`
	} `json:"values"`
}

var AzureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Fetch Azure IP ranges",
	Run: func(cmd *cobra.Command, args []string) {
		fetchAzureIPRanges()
	},
}

func init() {
	rootCmd.AddCommand(AzureCmd)
}

func fetchAzureIPRanges() {
	url := "https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20250303.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching Azure IP ranges:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var ipRanges AzureIPRanges
	if err := json.Unmarshal(body, &ipRanges); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Azure IP Ranges:")
	for _, value := range ipRanges.Values {
		for _, prefix := range value.Properties.AddressPrefixes {
			fmt.Printf("IP Prefix: %s, Region: %s, Service: %s\n", prefix, value.Properties.Region, value.Name)
		}
	}
}
