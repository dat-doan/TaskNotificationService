package fcm

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type Service struct {
	client *messaging.Client
}

func New(ctx context.Context, credentialsFile string) (*Service, error) {
	app, err := firebase.NewApp(ctx, nil, option.WithAuthCredentialsFile(option.ServiceAccount, credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("khởi tạo Firebase: %w", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("khởi tạo Messaging: %w", err)
	}
	return &Service{client: client}, nil
}

func (s *Service) SendOne(ctx context.Context, token, title, body string, data map[string]string) (string, error) {
	return s.client.Send(
		ctx, &messaging.Message{
			Token: token,
			Notification: &messaging.Notification{
				Title: title,
				Body: body,
			},
			Data: data,
		},
	)
}