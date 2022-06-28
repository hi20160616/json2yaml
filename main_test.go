package main

import "testing"

func TestJ2y(t *testing.T) {
	j, y := "./test/test.json", "./test/test.yaml"
	if err := j2y("header.json", y); err != nil {
		t.Fatal(err)
	}
	if err := j2y(j, y); err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}
