package main

import (
	"testing"
)

func Testmain(t *testing.T) {
	work("12", "20", "2020", "KHTX", "", 000212, 000212)
	want := "(+) FETCHING https://noaa-nexrad-level2.s3.amazonaws.com/2020/12/20/KHTX/KHTX20201220_000212_V06"

	if expected != want {
		t.Errorf("got %q want %q", expected, want)
	}
}
