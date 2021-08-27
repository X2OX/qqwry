package data

import (
	"testing"
)

func TestFind(t *testing.T) {
	ip := "114.114.114.114"
	cz := New().Find(ip)
	if cz.Error != nil {
		t.Fatal(cz.Error)
	}
	t.Logf("%s: %s %s", ip, cz.Country, cz.Area)

	if New().Find("...").Error == nil {
		t.Fatal(cz.Error)
	}
}
