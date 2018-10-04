package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Record struct {
	id     int
	name   string
	email  string
	mobile string
}

var store RecordAccessor

func main() {
	_, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Println("[WARN] interrupt signal")
		cancel()
	}()
	log.Println("Started server for receiving messages")

	var err error
	store, err = makeDataStore()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/receive", receiveHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func receiveHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("Receive raw message: %v", string(body))
	var record Record
	err = json.Unmarshal(body, &record)
	if err != nil {
		panic(err)
	}
	err = store.saveRecord(record)
	if err != nil {
		log.Printf("Record was not saved: %v", record)
	}
}
