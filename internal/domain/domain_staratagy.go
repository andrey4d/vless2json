/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package domain

type domainStrategy struct {
	AsIs         string
	UseIp        string
	IpIfNonMatch string
	IpOnDemand   string
}

var DomainStrategy = domainStrategy{
	AsIs:         "AsIs",
	UseIp:        "UseIp",
	IpIfNonMatch: "IpIfNonMatch",
	IpOnDemand:   "IpOnDemand",
}
