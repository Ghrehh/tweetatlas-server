package web

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type LocationAggregater interface {
	AddParsedLocation(string)
	ToJSON() []byte
}

type LocationAggregate struct {
	Data map[string]int `json:"location_data"`
	SearchPhrases []string `json:"search_phrases"`
	mutex *sync.Mutex
}

func NewLocationAggregate(searhPhrases []string) *LocationAggregate {
	return &LocationAggregate{
		Data: make(map[string]int),
		SearchPhrases: searhPhrases,
		mutex: &sync.Mutex{},
	}
}

func (la *LocationAggregate) AddParsedLocation(location string) {
	if location == "" {
		location = "unknown"
	}

	la.mutex.Lock()
	la.Data[location] +=1
	la.mutex.Unlock()
}

func (la *LocationAggregate) ToJSON() []byte {
	la.mutex.Lock()
	jsonString, err := json.Marshal(la)
	la.mutex.Unlock()

	if err != nil {
		fmt.Println("Error JSONifying the LocationAggregate")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return jsonString
}
