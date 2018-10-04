package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type Record struct {
	id     int
	name   string
	email  string
	mobile string
}

var workerNumber = 10

func main() {
	_, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Println("[WARN] interrupt signal")
		cancel()
	}()
	log.Println("Start parsing")

	queue := make(chan string)    //will contains each line
	workerDone := make(chan bool) //notify if file was processed

	go func() {
		file, err := os.Open("./data.csv") //todo remove hardcoded
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			queue <- scanner.Text()
		}
		close(queue)
	}()

	for i := 0; i < workerNumber; i++ {
		go startWorker(queue, workerDone)
	}

	for i := 0; i < workerNumber; i++ {
		<-workerDone
	}
}

func startWorker(queue chan string, workerDone chan bool) {
	for line := range queue {
		record, err := processLine(line)
		if err != nil {
			log.Printf("Skip line %s", line)
			continue
		}
		log.Printf("Proccessed: %s", record)
		//pass record to saver
	}

	workerDone <- true
}

func processLine(line string) (*Record, error) {
	splitted := strings.Split(line, ",")
	id, err := strconv.Atoi(splitted[0])
	if err != nil {
		return nil, fmt.Errorf("%q looks like a number.\n", splitted[0])
	}
	return &Record{
		id:     id,
		name:   splitted[1],
		email:  splitted[2],
		mobile: processPhone(splitted[3]),
	}, nil
}

func processPhone(phone string) string {
	//todo
	return phone
}
