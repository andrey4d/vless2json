/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package configTypes

type Welcome10 struct {
	DNS Dns `json:"dns"`
}

type Dns struct {
	Servers []string `json:"servers"`
}

func NewDns(servers []string) *Dns {
	return &Dns{
		Servers: servers,
	}
}
