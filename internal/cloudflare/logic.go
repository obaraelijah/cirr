package cloudflare

import (
	"fmt"
	"os"
	"strings"

	"github.com/obaraelijah/cirr/internal/utils"
)

type Config struct {
	IPType    string
	Verbosity string
}

func GetCloudflareIPRanges(config Config) {
	var url string
	if config.IPType == "ipv4" {
		url = "https://www.cloudflare.com/ips-v4/"
	} else if config.IPType == "ipv6" {
		url = "https://www.cloudflare.com/ips-v6/"
	} else {
		fmt.Fprintf(os.Stderr, "Unsupported IP type: %s\n", config.IPType)
		return
	}

	rawData := utils.GetRawData(url)
	ipRanges := parseIPRanges(rawData)

	printCloudflareIPRanges(ipRanges, config.Verbosity)
}

func parseIPRanges(rawData string) []string {
	lines := strings.Split(rawData, "\n")
	var ipRanges []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			ipRanges = append(ipRanges, trimmed)
		}
	}
	return ipRanges
}

func printCloudflareIPRanges(ipRanges []string, verbosity string) {
	if len(ipRanges) == 0 {
		fmt.Println("No IP ranges to display.")
		return
	}

	switch verbosity {
	case "none":
		for _, ip := range ipRanges {
			fmt.Println(ip)
		}
	case "mini":
		for _, ip := range ipRanges {
			fmt.Println(ip)
		}
	case "full":
		for _, ip := range ipRanges {
			fmt.Printf("Cloudflare IP: %s\n", ip)
		}
	default:
		for _, ip := range ipRanges {
			fmt.Println(ip)
		}
	}
}
