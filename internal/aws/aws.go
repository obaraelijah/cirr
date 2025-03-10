package aws

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/obaraelijah/cirr/internal/utils"
)

type IPv4Prefix struct {
	IPAddress          string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

type IPv6Prefix struct {
	IPv6Address        string `json:"ipv6_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

type IPsData struct {
	SyncToken    string       `json:"syncToken"`
	CreateDate   string       `json:"createDate"`
	Prefixes     []IPv4Prefix `json:"prefixes"`
	IPv6Prefixes []IPv6Prefix `json:"ipv6_prefixes"`
}

func GetIPRanges(ipType string, filter string, getReqFunc func(string) string) {
	raw_data := getReqFunc("https://ip-ranges.amazonaws.com/ip-ranges.json")

	filterValues := separateFilters(filter)

	readyIPs := filtrateIPRanges(raw_data, ipType, filterValues)

	printIPRanges(readyIPs)
}

func separateFilters(filterFlagValues string) []string {
	logger := utils.GetCirrLogger()
	var filterSlice []string

	removeFilterWhitespace := strings.ReplaceAll(filterFlagValues, " ", "")
	filterContents := strings.Split(removeFilterWhitespace, ",")

	for _, val := range filterContents {
		if len(val) > 0 {
			filterSlice = append(filterSlice, strings.TrimSpace(val))
		}
	}

	if len(filterSlice) == 0 && strings.Contains(filterFlagValues, ",") {
		logger.Fatalf("Filter flag needs actual values!")
	}

	return filterSlice
}

func filtrateIPRanges(rawData, ipType string, filterSlice []string) []string {
	logger := utils.GetCirrLogger()
	var data IPsData

	result := []string{}

	err := json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		logger.Fatalf("Error unmarshalling JSON: %v", err)
	}

	if ipType == "ipv4" {
		for _, prefix := range data.Prefixes {
			filteredIPv4String := fmt.Sprintf("%s,%s,%s,%s",
				prefix.IPAddress, prefix.Region, prefix.Service, prefix.NetworkBorderGroup)
			switch len(filterSlice) {
			case 0:
				result = append(result, filteredIPv4String)
			case 1:
				if prefix.Region == filterSlice[0] {
					result = append(result, filteredIPv4String)
				}
			case 2:
				if prefix.Region == filterSlice[0] && prefix.Service == filterSlice[1] {
					result = append(result, filteredIPv4String)
				}
			case 3:
				if prefix.Region == filterSlice[0] && prefix.Service == filterSlice[1] && prefix.NetworkBorderGroup == filterSlice[2] {
					result = append(result, filteredIPv4String)
				}
			}
		}
	} else if ipType == "ipv6" {
		for _, ipv6prefix := range data.IPv6Prefixes {
			filteredIPv6String := fmt.Sprintf("%s,%s,%s,%s",
				ipv6prefix.IPv6Address, ipv6prefix.Region, ipv6prefix.Service, ipv6prefix.NetworkBorderGroup)
			switch len(filterSlice) {
			case 0:
				result = append(result, filteredIPv6String)
			case 1:
				if ipv6prefix.Region == filterSlice[0] {
					result = append(result, filteredIPv6String)
				}
			case 2:
				if ipv6prefix.Region == filterSlice[0] && ipv6prefix.Service == filterSlice[1] {
					result = append(result, filteredIPv6String)
				}
			case 3:
				if ipv6prefix.Region == filterSlice[0] && ipv6prefix.Service == filterSlice[1] && ipv6prefix.NetworkBorderGroup == filterSlice[2] {
					result = append(result, filteredIPv6String)
				}
			}
		}
	}

	if len(result) == 0 {
		fmt.Println("Nothing found!")
		return []string{}
	}

	return result
}

func printIPRanges(IPranges []string) {
	for _, val := range IPranges {
		fmt.Println(val)
	}
}
