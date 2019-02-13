package firestore

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func TestTimeParsing(t *testing.T) {
	var date time.Time
	err := json.Unmarshal([]byte("\"2019-02-12T23:02:20.215Z\""), &date)
	mon := monetaryRequest{
		From: StringValue{"a"},
		To:   StringValue{"b"},
		Date: TimestampValue{date},
	}
	if err != nil {
		t.Errorf("err: %v", err)
	} else {
		bytes, err := json.Marshal(mon)
		log.Printf("date: %v", date)
		log.Printf("mon: %v", string(bytes))
		log.Printf("err: %v", err)
	}
	return
}
