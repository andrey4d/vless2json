/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package inbounds

type Welcome7 struct {
	Inbounds []Inbound `json:"inbounds"`
}

type Inbound struct {
	Tag      string   `json:"tag"`
	Port     string   `json:"port"`
	Protocol string   `json:"protocol"`
	Listen   string   `json:"listen"`
	Settings settings `json:"settings"`
	Sniffing sniffing `json:"sniffing"`
}

type settings struct {
	Auth      string `json:"auth"`
	UDP       bool   `json:"udp"`
	UserLevel int64  `json:"userLevel"`
}

type sniffing struct {
	Enabled bool `json:"enabled"`
}

func NewInbound(listenAddress, port string) Inbound {
	return Inbound{
		Tag:      "in_proxy",
		Port:     port,
		Protocol: "socks",
		Listen:   listenAddress,
		Settings: settings{
			Auth:      "noauth",
			UDP:       true,
			UserLevel: 8,
		},
		Sniffing: sniffing{
			Enabled: false,
		},
	}
}
