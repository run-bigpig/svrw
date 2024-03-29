package main

import (
	"encoding/json"
	"github.com/run-bigpig/svrw/internal/parser/pipixia"
	"log"
)

func main() {
	url := "https://h5.pipix.com/s/iFKLQkMD/"
	parser := pipixia.NewParser(url)
	result, err := parser.Parse()
	if err != nil {
		log.Println(err)
		return
	}
	// do something with result
	data, _ := json.Marshal(result)
	log.Println(string(data))
}
