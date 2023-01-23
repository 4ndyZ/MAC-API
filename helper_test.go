package main

import (
	"testing"
)

func Test_SplitMAC(t *testing.T) {
	// Sample data
	sampleOUI := "FCFC48"
	sampleMAC := "FCFC48CC3E5D"

	// Expected data
	expectedOUI := "FC-FC-48"
	expectedMAC := "FC-FC-48-CC-3E-5D"

	oui := SplitMAC(sampleOUI)
	if oui != expectedOUI {
		t.Errorf("OUI splitted incorrect. Wanted %s but got %s", expectedOUI, oui)
	}

	mac := SplitMAC(sampleMAC)

	if mac != expectedMAC {
		t.Errorf("MAC splitted incorrect. Wanted %s but got %s", expectedMAC, mac)
	}
}

func Test_RemoveAllNonHex(t *testing.T) {
	// Sample data
	sample := "FC-FZC-4T_"

	// Expected data
	expected := "FCFC4"

	// Remove all non-HEX chars
	result := RemoveAllNonHex(sample)

	if result != expected {
		t.Errorf("Non HEX chars not removed. Wanted %s but got %s", expected, result)
	}
}
