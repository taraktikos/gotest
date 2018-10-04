package main

import (
	"testing"
)

func TestProcessLine(t *testing.T) {
	input := "1,Kirk,ornare@sedtortor.net,(013890) 37420"
	record, err := processLine(input)
	if err != nil {
		t.Errorf("got=%v", err)
	}
	if record.id != 1 {
		t.Errorf("got=%v", record.id)
	}
	if record.name != "Kirk" {
		t.Errorf("got=%v", record.name)
	}
	if record.email != "ornare@sedtortor.net" {
		t.Errorf("got=%v", record.email)
	}
	if record.mobile != "(013890) 37420" {
		t.Errorf("got=%v", record.mobile)
	}
}

func TestProcessPhone(t *testing.T) {
	//todo
}
