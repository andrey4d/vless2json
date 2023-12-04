/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package routing

import "xray2json/internal/domain"

type Welcome7 struct {
	Routing Routing `json:"routing"`
}

type Routing struct {
	DomainStrategy string   `json:"domainStrategy"`
	Rules          []Rule   `json:"rules"`
	Balancers      []string `json:"balancers"`
}

type Rule struct {
	Type        string   `json:"type"`
	Domain      []string `json:"domain",omitempty`
	OutboundTag string   `json:"outboundTag"`
	IP          []string `json:"ip",omitempty`
}

func NewRouting() *Routing {
	return &Routing{
		DomainStrategy: domain.DomainStrategy.IpIfNonMatch,
		Rules: []Rule{
			{
				Type: "field",
				Domain: []string{
					"domain:ru",
					"geosite:category-gov-ru",
					"geosite:yandex",
					"geosite:mailru",
				},
				OutboundTag: "Direct",
			},
			{
				Type: "field",
				IP: []string{
					"geoip:ru",
					"geoip:private",
				},
				OutboundTag: "Direct",
			},
			{
				Type:        "field",
				Domain:      []string{"geosite:category-ads-all"},
				OutboundTag: "Direct",
			},
		},
		Balancers: []string{},
	}
}
