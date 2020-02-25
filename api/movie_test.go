package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJson(t *testing.T){
	m := Movie{Id: "1234", Title: "Joker", Tagline: "Put on a happy face.", Director: "Todd Phillips"}
	actual, _ := json.Marshal(m)
	assert.Equal(t,
		`{"id":"1234","title":"Joker","tagline":"Put on a happy face.","director":"Todd Phillips"}`,
		string(actual),
		"Movie JSON marshalling is wrong")
}

func TestFromJson(t *testing.T){
	json := []byte(`{"id":"1234","title":"Joker","tagline":"Put on a happy face.","director":"Todd Phillips"}`)
	actual, _ := FromJson(json)
	assert.Equal(t,
		Movie{Id: "1234", Title: "Joker", Tagline: "Put on a happy face.", Director: "Todd Phillips"},
		actual,
		"Movie JSON unmarshalling is wrong")
}