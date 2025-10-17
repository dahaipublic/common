package util

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SendMessage2() {
	// 使用 Service Account JSON 初始化
	opt := option.WithCredentialsFile("D:\\trball\\etc\\serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v", err)
	}

	// 发送 Topic 消息
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "广播标题",
			Body:  "这里是广播内容",
		},
		Topic: "allUsers", // 之前客户端订阅的 topic
	}

	response, err := client.Send(context.Background(), message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Successfully sent message:", response)
}
