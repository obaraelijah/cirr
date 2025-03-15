package cmd

import (
	"fmt"
	"os"

	"github.com/obaraelijah/cirr/internal/icloud"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var icloudCmd = &cobra.Command{
	Use:   "icloud",
	Short: "Get iCloud private relay IP ranges.",
	Long:  `Get iCloud private relay IPv4 and IPv6 ranges.`,
	Run: func(cmd *cobra.Command, args []string) {
		var verbosity string
		if cmd.Flags().Changed("verbose-mode") {
			verbosity = viper.GetString("verbose_mode")
		} else if viper.GetBool("verbose") {
			verbosity = "full"
		} else {
			verbosity = "none"
		}

		if !isValidVerbosity(verbosity) {
			fmt.Fprintf(os.Stderr, "Invalid verbosity level: %s. Allowed values are: none, mini, full.\n", verbosity)
			os.Exit(1)
		}

		ipType := "ipv4"
		if viper.GetBool("ipv6") {
			ipType = "ipv6"
		}

		if !viper.GetBool("ipv4") && !viper.GetBool("ipv6") {
			ipType = "both"
		}

		config := icloud.Config{
			IPType: ipType,
			Filters: icloud.Filters{
				Country: viper.GetStringSlice("filter-country"),
				Region:  viper.GetStringSlice("filter-region"),
				City:    viper.GetStringSlice("filter-city"),
			},
			Verbosity: verbosity,
		}

		icloud.GetIPRanges(config)
	},
}

func init() {
	icloudCmd.Flags().Bool("ipv4", false, "Get only IPv4 ranges")
	icloudCmd.Flags().Bool("ipv6", false, "Get only IPv6 ranges")
	icloudCmd.Flags().StringSlice("filter-country", []string{}, "Filter results by country")
	icloudCmd.Flags().StringSlice("filter-region", []string{}, "Filter results by region")
	icloudCmd.Flags().StringSlice("filter-city", []string{}, `Filter results by city (use quotes for names with spaces, e.g. "New York")`)

	viper.BindPFlag("ipv4", icloudCmd.Flags().Lookup("ipv4"))
	viper.BindPFlag("ipv6", icloudCmd.Flags().Lookup("ipv6"))
	viper.BindPFlag("filter-country", icloudCmd.Flags().Lookup("filter-country"))
	viper.BindPFlag("filter-region", icloudCmd.Flags().Lookup("filter-region"))
	viper.BindPFlag("filter-city", icloudCmd.Flags().Lookup("filter-city"))

	rootCmd.AddCommand(icloudCmd)
}

func isValidVerbosity(verbosity string) bool {
	switch verbosity {
	case "none", "mini", "full":
		return true
	default:
		return false
	}
}
