package main

// SplitMAC splits an OUI/MAC after every 2 chars and add a dash
func SplitMAC(s string) string {
	for i := 2; i < len(s); i += 3 {
		s = s[:i] + "-" + s[i:]
	}
	return s
}
