package cmd

import (
	"github.com/obaraelijah/cirr/internal/aws"
	"github.com/obaraelijah/cirr/internal/utils"
	"github.com/spf13/cobra"
)

var (
	awsIPv4Flag     bool
	awsIPv6Flag     bool
	awsIPFilterFlag string
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Get AWS IP ranges.",
	Long:  `Get AWS IPv4 and IPv6 ranges.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := utils.GetCirrLogger()

		logger.Println("AWS subcommand called")

		if awsIPv4Flag || (!awsIPv4Flag && !awsIPv6Flag) {
			aws.GetIPRanges("ipv4", awsIPFilterFlag)
		}
		if awsIPv6Flag || (!awsIPv4Flag && !awsIPv6Flag) {
			aws.GetIPRanges("ipv6", awsIPFilterFlag)
		}
	},
}

func main() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.Flags().BoolVar(&awsIPv4Flag, "ipv4", false, "Get only IPv4 ranges")
	awsCmd.Flags().BoolVar(&awsIPv6Flag, "ipv6", false, "Get only IPv6 ranges")
	awsCmd.Flags().StringVar(&awsIPFilterFlag, "filter", "", "Filter results. Syntax: aws-region-az,SERVICE,network-border-group")
}
