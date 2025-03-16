package cmd

import (
	"fmt"
	"os"

	"github.com/obaraelijah/cirr/internal/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Get AWS IP ranges.",
	Long:  `Get AWS IPv4 and IPv6 ranges with optional filtering.`,
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

		ipVersion := []string{}

		if viper.GetBool("aws_ipv4") || (!viper.GetBool("aws_ipv4") && !viper.GetBool("aws_ipv6")) {
			ipVersion = append(ipVersion, "aws_ipv4")
		}
		if viper.GetBool("aws_ipv6") || (!viper.GetBool("aws_ipv4") && !viper.GetBool("aws_ipv6")) {
			ipVersion = append(ipVersion, "aws_ipv6")
		}

		filter := viper.GetString("aws-filter")
		filterRegion := viper.GetString("aws-filter-region")
		filterService := viper.GetString("aws-filter-service")
		filterNetworkBorderGroup := viper.GetString("aws-filter-network-border-group")

		if filter != "" && (filterRegion != "" || filterService != "" || filterNetworkBorderGroup != "") {
			fmt.Fprintln(os.Stderr, "--filter flag cannot be used with individual filter flags")
			os.Exit(1)
		}

		var awsFilter string
		if filter != "" {
			awsFilter = filter
		} else {
			awsFilter = fmt.Sprintf("%s,%s,%s", filterRegion, filterService, filterNetworkBorderGroup)
		}

		config := aws.Config{
			IPType:    "",
			Filter:    awsFilter,
			Verbosity: verbosity,
		}

		for _, version := range ipVersion {
			fmt.Println(version)
			config.IPType = version
			aws.GetIPRanges(config)
		}
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.Flags().Bool("ipv4", false, "Get only IPv4 ranges")
	awsCmd.Flags().Bool("ipv6", false, "Get only IPv6 ranges")
	awsCmd.Flags().String("filter", "", "Filter results. Syntax: aws-region-az,SERVICE,network-border-group")

	awsCmd.Flags().String("filter-region", "", "Filter results by AWS region")
	awsCmd.Flags().String("filter-service", "", "Filter results by AWS service")
	awsCmd.Flags().String("filter-network-border-group", "", "Filter results by AWS network border group")

	viper.BindPFlag("aws_ipv4", awsCmd.Flags().Lookup("ipv4"))
	viper.BindPFlag("aws_ipv6", awsCmd.Flags().Lookup("ipv6"))
	viper.BindPFlag("aws-filter", awsCmd.Flags().Lookup("filter"))
	viper.BindPFlag("aws-filter-region", awsCmd.Flags().Lookup("filter-region"))
	viper.BindPFlag("aws-filter-service", awsCmd.Flags().Lookup("filter-service"))
	viper.BindPFlag("aws-filter-network-border-group", awsCmd.Flags().Lookup("filter-network-border-group"))
}
