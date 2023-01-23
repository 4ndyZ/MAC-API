package main

import "regexp"

// SplitMAC splits an OUI/MAC after every 2 chars and add a dash
func SplitMAC(s string) string {
	for i := 2; i < len(s); i += 3 {
		s = s[:i] + "-" + s[i:]
	}
	return s
}

// RemoveAllNonHex removes all non-HEX chars from a string
func RemoveAllNonHex(s string) string {
	reg := regexp.MustCompile("[^0-9A-F]+")
	return reg.ReplaceAllString(s, "")
}
