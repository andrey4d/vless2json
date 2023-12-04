/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package configTypes

import (
	"xray2json/internal/configTypes/inbounds"
	"xray2json/internal/configTypes/outbounds"
	"xray2json/internal/configTypes/routing"
)

type Config struct {
	Comment   Comment              `json:"_comment"`
	Log       Log                  `json:"log"`
	Inbounds  []inbounds.Inbound   `json:"inbounds"`
	Outbounds []outbounds.Outbound `json:"outbounds"`
	DNS       Dns                  `json:"dns"`
	Routing   routing.Routing      `json:"routing"`
}
