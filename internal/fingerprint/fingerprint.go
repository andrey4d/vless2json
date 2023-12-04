/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package fingerprint

type fingerprint struct {
	Randomized        string
	Randomizedalpn    string
	Randomizednoalpn  string
	Firefox_auto      string
	Dhrome_auto       string
	Ios_auto          string
	Android_11_okhttp string
	Edge_auto         string
	Safari_auto       string
	A360_auto         string
	Qq_auto           string
}

var Fingerprint = fingerprint{
	Randomized:        "randomized",
	Randomizedalpn:    "randomizedalpn",
	Randomizednoalpn:  "randomizednoalpn",
	Firefox_auto:      "firefox_auto",
	Dhrome_auto:       "chrome_auto",
	Ios_auto:          "ios_auto",
	Android_11_okhttp: "android_11_okhttp",
	Edge_auto:         "edge_auto",
	Safari_auto:       "safari_auto",
	A360_auto:         "360_auto",
	Qq_auto:           "qq_auto",
}
