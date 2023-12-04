/*
 *   Copyright (c) 2023 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"xray2json/internal/configTypes"
	"xray2json/internal/configTypes/inbounds"
	"xray2json/internal/configTypes/outbounds"
	"xray2json/internal/configTypes/routing"
	"xray2json/internal/domain"
)

func Vless(uri string) {
	// infoLogger := loggers.NewInfoLogger(os.Stdout)

	p, err := parseURI(uri)
	if err != nil {
		log.Fatal(err)
	}

	j := configTypes.Config{
		Comment: configTypes.Comment{
			Remark: p.Comment,
		},
		Log:      *configTypes.NewLog(),
		DNS:      *configTypes.NewDns([]string{"8.8.8.8"}),
		Inbounds: []inbounds.Inbound{inbounds.NewInbound("0.0.0.0", "1080")},
		Outbounds: []outbounds.Outbound{
			*outbounds.NewOutbound(*p),
			{
				Tag:      "direct",
				Protocol: "freedom",
				Settings: outbounds.Settings{
					DomainStrategy: domain.DomainStrategy.UseIp,
				},
			},
			{
				Tag:      "blackhole",
				Protocol: "blackhole",
				Settings: outbounds.Settings{},
			},
		},
		Routing: *routing.NewRouting(),
	}

	js, err := json.Marshal(j)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(js))

}

func parseURI(uri string) (*outbounds.Parameters, error) {
	// decode URL by url.Parse
	parsedURL, err := url.Parse(uri)
	if err != nil {
		return nil, err

	}

	stSettings := make(map[string]string)

	for key, values := range parsedURL.Query() {
		stSettings[key] = values[0]

	}

	p := outbounds.Parameters{
		Protocol:      parsedURL.Scheme,
		UserId:        parsedURL.User.Username(),
		ServerAddress: parsedURL.Hostname(),
		Port:          parsedURL.Port(),
		Path:          parsedURL.Path,
		Comment:       parsedURL.Fragment,
		StSettings:    stSettings,
	}

	return &p, err
}
