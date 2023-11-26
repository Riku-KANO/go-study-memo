package main

import (
	"log"
	"time"
	"fmt"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
)

var (
	topic = "go.micro.topic.foo"
)

func pub1() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub1] failed: %v", err)
		} else {
			fmt.Println("[pub1] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func pub2() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub2] failed: %v", err)
		} else {
			fmt.Println("[pub2] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func sub1() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub1] recieved message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func sub2() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub2] recieved message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker connect error: %v", err)
	}

	go pub1()
	// go pub2()
	go sub1()
	// go sub2()
	<- time.After(time.Second * 10)
}