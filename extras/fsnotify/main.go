package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}
	defer watcher.Close()
	MarshalConfig("config.json")
	log.Printf("1Config: %+v\n", config)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("Event: %s\n", event)
				MarshalConfig("config.json")
				fmt.Printf("modified file: %s\n", event.Name)
				fmt.Println(config)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Error: %v\n", err)
			}
		}
	}()

	// Watch the config file for changes
	err = watcher.Add("config.json")
	if err != nil {
		log.Fatalf("Error watching config file: %v", err)
	}

	// Wait for a signal to stop
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error marshaling config: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error writing config to file: %v", err)
	}
}
