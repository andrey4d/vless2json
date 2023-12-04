/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package outbounds

import (
	"strconv"
)

type Welcome9 struct {
	Outbounds []Outbound `json:"outbounds"`
}

type Outbound struct {
	Tag            string          `json:"tag"`
	Protocol       string          `json:"protocol"`
	Settings       Settings        `json:"settings"`
	StreamSettings *streamSettings `json:"streamSettings,omitempty"`
	Mux            *mux            `json:"mux,omitempty"`
}

type mux struct {
	Enabled     bool  `json:"enabled"`
	Concurrency int64 `json:"concurrency"`
}

type Settings struct {
	Vnext          []vnext `json:"vnext,omitempty"`
	DomainStrategy string  `json:"domainStrategy"`
}

type vnext struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Users   []user `json:"users"`
}

type user struct {
	ID         string `json:"id"`
	Security   string `json:"security"`
	Level      int16  `json:"level"`
	Encryption string `json:"encryption"`
	Flow       string `json:"flow"`
}

type streamSettings struct {
	Network     string      `json:"network"`
	Security    string      `json:"security"`
	WsSettings  wsSettings  `json:"wsSettings"`
	TLSSettings tLSSettings `json:"tlsSettings"`
}

type tLSSettings struct {
	AllowInsecure bool   `json:"allowInsecure"`
	ServerName    string `json:"serverName"`
	Show          bool   `json:"show"`
	PublicKey     string `json:"publicKey"`
	ShortID       string `json:"shortId"`
	SpiderX       string `json:"spiderX"`
}

type wsSettings struct {
	Path    string  `json:"path"`
	Headers headers `json:"headers"`
}

type headers struct {
	Host string `json:"Host"`
}

type Parameters struct {
	Protocol      string
	UserId        string
	ServerAddress string
	Port          string
	Path          string
	Comment       string
	StSettings    map[string]string
}

func NewOutbound(parameters Parameters) *Outbound {
	return &Outbound{
		Tag:      "proxy",
		Protocol: parameters.Protocol,
		Settings: Settings{
			Vnext: []vnext{NewVnext(parameters.UserId, parameters.ServerAddress, parameters.Port)},
		},
		StreamSettings: NewStreamSettings(parameters.StSettings),
		Mux: &mux{
			Enabled:     false,
			Concurrency: 8,
		},
	}
}

func NewUser(userId string) user {
	return user{
		ID:         userId,
		Security:   "auto",
		Level:      8,
		Encryption: "none",
		Flow:       "",
	}
}

func NewVnext(userId, serverAddress string, port string) vnext {
	p, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	return vnext{
		Address: serverAddress,
		Port:    p,
		Users:   []user{NewUser(userId)},
	}
}

func NewStreamSettings(sts map[string]string) *streamSettings {
	return &streamSettings{
		Network:  sts["type"],
		Security: sts["security"],
		WsSettings: wsSettings{
			Path: sts["path"],
			Headers: headers{
				Host: sts["sni"],
			},
		},
		TLSSettings: tLSSettings{
			AllowInsecure: true,
			ServerName:    sts["host"],
			Show:          false,
			PublicKey:     "",
			ShortID:       "",
			SpiderX:       "",
		},
	}
}
