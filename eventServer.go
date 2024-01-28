package main

import (
	"fmt"
	"gopkg.in/antage/eventsource.v1"
	"log"
	redis2 "main/redis"
	"net/http"
	"strconv"
)

func main() {
	redis, ctx, err := redis2.NewRedis(redis2.LoadEnv())
	if err != nil {
		log.Fatalf("Error connecting to redis: %s", err.Error())
	}
	defer redis.Close()

	sub := redis.PSubscribe(ctx, "__key*@*__:xadd")
	ch := sub.Channel()
	defer sub.Close()

	es := eventsource.New(
		eventsource.DefaultSettings(),
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("Access-Control-Allow-Origin: *"),
			}
		},
	)
	defer es.Close()
	http.Handle("/events", es)

	go func() {
		id := 1
		fmt.Println("Waiting for messages")
		for msg := range ch {
			fmt.Println(msg.Payload)
			es.SendEventMessage(msg.Payload, "event-channel", strconv.Itoa(id))
			id++

			messages := redis.XRevRangeN(ctx, msg.Payload, "+", "-", 1).Val()

			for _, message := range messages {
				fmt.Println(message.Values["message"])
				es.SendEventMessage(fmt.Sprintf("%v", message.Values["message"]), "event-message", strconv.Itoa(id))
				id++
			}
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
