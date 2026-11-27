package main

import (
	"context"
	"fmt"
	"os"
	"push-service/internal/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return fmt.Errorf("cách dùng: go run ./cmd/send-test <device-token>")
	}

	token := os.Args[1]
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithAuthCredentialsFile(option.ServiceAccount, cfg.FirebaseCredentials))
	
	if err != nil {
		return err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	id, err := client.Send(ctx, &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: "Xin chào",
			Body: "Đây là push đầu tiên từ server của bạn",
		},
		Data: map[string]string{
			"screen": "Home",
		},
	})
	if err != nil {
		return err
	}
	fmt.Println("FCM đã nhận:", id)
	return nil
}

