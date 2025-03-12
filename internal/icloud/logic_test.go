// internal/icloud/filtrate_ip_ranges_test.go
package icloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiltrateIPRanges(t *testing.T) {
	ipRanges := []IPRange{
		{"2a02:26f7:f6f9:800::/54", "US", "US-NY", "New York"},
		{"2a02:26f7:f6f9:a06a::/64", "US", "US-NY", "New York"},
		{"2a02:26f7:f6fc:800::/54", "US", "US-NY", "New York"},
		{"2a02:26f7:f6fc:a06a::/64", "US", "US-NY", "New York"},
		{"2606:54c0:a620::/45", "US", "US-NY", "New York"},
		{"2a09:bac2:a620::/45", "US", "US-NY", "New York"},
		{"2a09:bac3:a620::/45", "US", "US-NY", "New York"},
		{"104.28.129.23/32", "DE", "DE-BE", "Berlin"},
		{"104.28.129.24/32", "DE", "DE-BE", "Berlin"},
		{"104.28.129.25/32", "DE", "DE-BE", "Berlin"},
		{"104.28.129.26/32", "DE", "DE-BE", "Berlin"},
		{"140.248.17.50/31", "DE", "DE-BE", "Berlin"},
		{"140.248.34.44/31", "DE", "DE-BE", "Berlin"},
		{"140.248.36.56/31", "DE", "DE-BE", "Berlin"},
		{"146.75.166.14/31", "DE", "DE-BE", "Berlin"},
		{"146.75.169.44/31", "DE", "DE-BE", "Berlin"},
		{"172.224.240.128/27", "JP", "JP-13", "Tokyo"},
		{"172.225.46.64/26", "JP", "JP-13", "Tokyo"},
		{"172.225.46.208/28", "JP", "JP-13", "Tokyo"},
		{"2a04:4e41:0030:0007::/64", "JP", "JP-13", "Tokyo"},
		{"2a04:4e41:0035:0006::/64", "JP", "JP-13", "Tokyo"},
		{"2a04:4e41:0064:000b::/64", "JP", "JP-13", "Tokyo"},
	}

	testCases := []struct {
		name          string
		ipType        string
		filterCountry string
		filterState   string
		filterCity    string
		expected      []IPRange
	}{
		{
			name:          "IPv6 US-NY-New York",
			ipType:        "ipv6",
			filterCountry: "US",
			filterState:   "US-NY",
			filterCity:    "New York",
			expected: []IPRange{
				{"2a02:26f7:f6f9:800::/54", "US", "US-NY", "New York"},
				{"2a02:26f7:f6f9:a06a::/64", "US", "US-NY", "New York"},
				{"2a02:26f7:f6fc:800::/54", "US", "US-NY", "New York"},
				{"2a02:26f7:f6fc:a06a::/64", "US", "US-NY", "New York"},
				{"2606:54c0:a620::/45", "US", "US-NY", "New York"},
				{"2a09:bac2:a620::/45", "US", "US-NY", "New York"},
				{"2a09:bac3:a620::/45", "US", "US-NY", "New York"},
			},
		},
		{
			name:          "IPv4 DE-BE-Berlin",
			ipType:        "ipv4",
			filterCountry: "DE",
			filterState:   "DE-BE",
			filterCity:    "Berlin",
			expected: []IPRange{
				{"104.28.129.23/32", "DE", "DE-BE", "Berlin"},
				{"104.28.129.24/32", "DE", "DE-BE", "Berlin"},
				{"104.28.129.25/32", "DE", "DE-BE", "Berlin"},
				{"104.28.129.26/32", "DE", "DE-BE", "Berlin"},
				{"140.248.17.50/31", "DE", "DE-BE", "Berlin"},
				{"140.248.34.44/31", "DE", "DE-BE", "Berlin"},
				{"140.248.36.56/31", "DE", "DE-BE", "Berlin"},
				{"146.75.166.14/31", "DE", "DE-BE", "Berlin"},
				{"146.75.169.44/31", "DE", "DE-BE", "Berlin"},
			},
		},
		{
			name:          "All IP types in Tokyo",
			ipType:        "ipv4",
			filterCountry: "",
			filterState:   "",
			filterCity:    "Tokyo",
			expected: []IPRange{
				{"172.224.240.128/27", "JP", "JP-13", "Tokyo"},
				{"172.225.46.64/26", "JP", "JP-13", "Tokyo"},
				{"172.225.46.208/28", "JP", "JP-13", "Tokyo"},
			},
		},
		{
			name:          "IPv4 DE-BE without city filter",
			ipType:        "ipv4",
			filterCountry: "DE",
			filterState:   "DE-BE",
			filterCity:    "",
			expected: []IPRange{
				{"104.28.129.23/32", "DE", "DE-BE", "Berlin"},
				{"104.28.129.24/32", "DE", "DE-BE", "Berlin"},
				{"104.28.129.25/32", "DE", "DE-BE", "Berlin"},
				{"104.28.129.26/32", "DE", "DE-BE", "Berlin"},
				{"140.248.17.50/31", "DE", "DE-BE", "Berlin"},
				{"140.248.34.44/31", "DE", "DE-BE", "Berlin"},
				{"140.248.36.56/31", "DE", "DE-BE", "Berlin"},
				{"146.75.166.14/31", "DE", "DE-BE", "Berlin"},
				{"146.75.169.44/31", "DE", "DE-BE", "Berlin"},
			},
		},
		{
			name:          "IPv6 US without state and city filter",
			ipType:        "ipv6",
			filterCountry: "US",
			filterState:   "",
			filterCity:    "",
			expected: []IPRange{
				{"2a02:26f7:f6f9:800::/54", "US", "US-NY", "New York"},
				{"2a02:26f7:f6f9:a06a::/64", "US", "US-NY", "New York"},
				{"2a02:26f7:f6fc:800::/54", "US", "US-NY", "New York"},
				{"2a02:26f7:f6fc:a06a::/64", "US", "US-NY", "New York"},
				{"2606:54c0:a620::/45", "US", "US-NY", "New York"},
				{"2a09:bac2:a620::/45", "US", "US-NY", "New York"},
				{"2a09:bac3:a620::/45", "US", "US-NY", "New York"},
			},
		},
		{
			name:          "No filters",
			ipType:        "ipv6",
			filterCountry: "",
			filterState:   "",
			filterCity:    "",
			expected: []IPRange{
				{"2a02:26f7:f6f9:800::/54", "US", "US-NY", "New York"},
				{"2a02:26f7:f6f9:a06a::/64", "US", "US-NY", "New York"},
				{"2a02:26f7:f6fc:800::/54", "US", "US-NY", "New York"},
				{"2a02:26f7:f6fc:a06a::/64", "US", "US-NY", "New York"},
				{"2606:54c0:a620::/45", "US", "US-NY", "New York"},
				{"2a09:bac2:a620::/45", "US", "US-NY", "New York"},
				{"2a09:bac3:a620::/45", "US", "US-NY", "New York"},
				{"2a04:4e41:0030:0007::/64", "JP", "JP-13", "Tokyo"},
				{"2a04:4e41:0035:0006::/64", "JP", "JP-13", "Tokyo"},
				{"2a04:4e41:0064:000b::/64", "JP", "JP-13", "Tokyo"},
			},
		},
		{
			name:          "City Filter Tokyo",
			ipType:        "ipv4",
			filterCountry: "",
			filterState:   "",
			filterCity:    "Tokyo",
			expected: []IPRange{
				{"172.224.240.128/27", "JP", "JP-13", "Tokyo"},
				{"172.225.46.64/26", "JP", "JP-13", "Tokyo"},
				{"172.225.46.208/28", "JP", "JP-13", "Tokyo"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := filtrateIPRanges(ipRanges, tc.ipType, tc.filterCountry, tc.filterState, tc.filterCity)
			assert.Equal(t, tc.expected, result, "Test case '%s' failed", tc.name)
		})
	}
}
