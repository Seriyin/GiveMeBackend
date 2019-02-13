package firestore

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestTimeParsing(t *testing.T) {
	mon := monetaryRequest{
		From: "a",
		To:   "b",
		Date: "2019-02-12T23:02:20.215Z",
	}
	date, err := time.Parse(
		"2006-01-02T15:04:05",
		strings.SplitN(mon.Date, ".", -1)[0],
	)
	if err != nil {
		t.Errorf("err: %v", err)
	} else {
		log.Printf("date: %v", date)
	}
	return
}
