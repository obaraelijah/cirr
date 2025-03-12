package cmd

import (
	"github.com/obaraelijah/cirr/internal/icloud"
	"github.com/spf13/cobra"
)

var (
	icloudIPv4Flag          bool
	icloudIPv6Flag          bool
	icloudFilterCountryFlag []string
	iCloudFilterStateFlag   []string
	iCloudFilterCityFlag    []string
	iCloudVerbosityFlag     string
)

var icloudCmd = &cobra.Command{
	Use:   "icloud",
	Short: "Get iCloud private relay IP ranges.",
	Long:  `Get iCloud private relay IPv4 and IPv6 ranges.`,
	Run: func(cmd *cobra.Command, args []string) {
		if icloudIPv4Flag || (!icloudIPv4Flag && !icloudIPv6Flag) {
			icloud.GetIPRanges("ipv4", icloudFilterCountryFlag, iCloudFilterStateFlag, iCloudFilterCityFlag, iCloudVerbosityFlag)
		}
		if icloudIPv6Flag || (!icloudIPv4Flag && !icloudIPv6Flag) {
			icloud.GetIPRanges("ipv6", icloudFilterCountryFlag, iCloudFilterStateFlag, iCloudFilterCityFlag, iCloudVerbosityFlag)
		}
	},
}

func init() {
	icloudCmd.Flags().BoolVar(&icloudIPv4Flag, "ipv4", false, "Get only IPv4 ranges")
	icloudCmd.Flags().BoolVar(&icloudIPv6Flag, "ipv6", false, "Get only IPv6 ranges")
	icloudCmd.Flags().StringSliceVar(&icloudFilterCountryFlag, "filter-country", []string{}, "Filter results by country")
	icloudCmd.Flags().StringSliceVar(&iCloudFilterStateFlag, "filter-state", []string{}, "Filter results by state")
	icloudCmd.Flags().StringSliceVar(&iCloudFilterCityFlag, "filter-city", []string{}, "Filter results by city (use quotes for names with spaces, e.g. \"New York\")")
	icloudCmd.Flags().StringVarP(&iCloudVerbosityFlag, "verbosity", "v", "none", "Verbosity level: none, mini, full")
	rootCmd.AddCommand(icloudCmd)
}
