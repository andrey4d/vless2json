/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package configTypes

type Log struct { // Log json:"log"
	Access   string `json:"access"`
	Error    string `json:"error"`
	Loglevel string `json:"loglevel"`
	DNSLog   bool   `json:"dnsLog"`
}

func NewLog() *Log {
	return &Log{
		Access:   "",
		Error:    "",
		Loglevel: "info",
		DNSLog:   false,
	}
}
