package cloudflare

import (
	"fmt"

	"github.com/obaraelijah/cirr/internal/utils"
)

func GetCloudflareIPv4Ranges() {
	data := utils.GetRawData("https://www.cloudflare.com/ips-v4/")

	fmt.Println(data)
}

func GetCloudflareIPv6Ranges() {
	data := utils.GetRawData("https://www.cloudflare.com/ips-v6/")

	fmt.Println(data)
}
