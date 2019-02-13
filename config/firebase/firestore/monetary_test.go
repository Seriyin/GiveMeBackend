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

func TestParseFromJSON(t *testing.T) {
	ex := "{\"amountCents\":{\"integerValue\":\"52\"},\"amountUnit\":{\"integerValue\":\"432\"},\"confirmedFrom\":{\"booleanValue\":false},\"confirmedTo\":{\"booleanValue\":false},\"currency\":{\"stringValue\":\"€\"},\"date\":{\"timestampValue\":\"2019-02-13T00:21:13.036Z\"},\"desc\":{\"stringValue\":\"Woop\"},\"from\":{\"stringValue\":\"+351345345345\"},\"groupId\":{\"integerValue\":\"-1\"},\"recurrentId\":{\"integerValue\":\"-1\"},\"snowflake\":{\"stringValue\":\"3zwUD2mxrrsAs4QsDRP4\"},\"to\":{\"stringValue\":\"+351366366366\"}}"
	var mon monetaryRequest
	err := json.Unmarshal([]byte(ex), &mon)
	if err != nil {
		t.Errorf("err: %v", err)
	} else {
		bytes, err := json.Marshal(mon)
		log.Printf("mon: %v", string(bytes))
		log.Printf("err: %v", err)
	}
	return
}
