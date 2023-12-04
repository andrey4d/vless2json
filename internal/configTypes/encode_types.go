/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package configTypes

type protocol struct {
	ProtocolScheme string
	ProtocolName   string
}

var VMESS = protocol{
	ProtocolScheme: "vmess://",
	ProtocolName:   "vmess",
}

var SHADOWSOCKS = protocol{
	ProtocolScheme: "ss://",
	ProtocolName:   "ss",
}
var SOCKS = protocol{
	ProtocolScheme: "socks://",
	ProtocolName:   "socks",
}
var VLESS = protocol{
	ProtocolScheme: "vless://",
	ProtocolName:   "vless",
}
var TROJAN = protocol{
	ProtocolScheme: "trojan://",
	ProtocolName:   "trojan",
}
var WIREGUARD = protocol{
	ProtocolScheme: "wireguard://",
	ProtocolName:   "wireguard",
}
var FREEDOM = protocol{
	ProtocolScheme: "freedom://",
	ProtocolName:   "freedom",
}
var BLACKHOLE = protocol{
	ProtocolScheme: "blackhole://",
	ProtocolName:   "blackhole",
}
