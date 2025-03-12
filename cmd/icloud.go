package cmd

import (
	"github.com/obaraelijah/cirr/internal/icloud"
	"github.com/spf13/cobra"
)

var (
	icloudIPv4Flag          bool
	icloudIPv6Flag          bool
	icloudFilterCountryFlag string
	iCloudFilterStateFlag   string
	iCloudFilterCityFlag    string
)

var icloudCmd = &cobra.Command{
	Use:   "icloud",
	Short: "Get iCloud private relay IP ranges.",
	Long:  `Get iCloud private relay IPv4 and IPv6 ranges.`,
	Run: func(cmd *cobra.Command, args []string) {
		if icloudIPv4Flag || (!icloudIPv4Flag && !icloudIPv6Flag) {
			icloud.GetIPRanges("ipv4", icloudFilterCountryFlag, iCloudFilterStateFlag, iCloudFilterCityFlag)
		}
		if icloudIPv6Flag || (!icloudIPv4Flag && !icloudIPv6Flag) {
			icloud.GetIPRanges("ipv6", icloudFilterCountryFlag, iCloudFilterStateFlag, iCloudFilterCityFlag)
		}
	},
}

func init() {
	icloudCmd.Flags().BoolVar(&icloudIPv4Flag, "ipv4", false, "Get only IPv4 ranges")
	icloudCmd.Flags().BoolVar(&icloudIPv6Flag, "ipv6", false, "Get only IPv6 ranges")
	icloudCmd.Flags().StringVar(&icloudFilterCountryFlag, "filter-country", "", "Filter results by country")
	icloudCmd.Flags().StringVar(&iCloudFilterStateFlag, "filter-state", "", "Filter results by state")
	icloudCmd.Flags().StringVar(&iCloudFilterCityFlag, "filter-city", "", "Filter results by city (use quotes for names with spaces, e.g. \"New York\")")
	rootCmd.AddCommand(icloudCmd)
}
