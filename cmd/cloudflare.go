package cmd

import (
	"github.com/obaraelijah/cirr/internal/cloudflare"
	"github.com/obaraelijah/cirr/internal/utils"
	"github.com/spf13/cobra"
)

var cloudflareCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "Get Cloudflare ip ranges",
	Long:  `Get Cloudflare ip ranges`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := utils.GetCirrLogger()

		logger.Println("cloudflare called")

		cloudflare.GetCloudflareIPv4Ranges()
		cloudflare.GetCloudflareIPv6Ranges()
	},
}

func init() {
	rootCmd.AddCommand(cloudflareCmd)
}
