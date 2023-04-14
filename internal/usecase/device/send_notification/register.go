package device

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"

	"github.com/eltoncasacio/vantracking/configs"
	vo "github.com/eltoncasacio/vantracking/internal/domain/shared/valueobjects"
)

type SendNotificationUseCase struct {
	repository vo.DeviceRepositoryInterface
}

func NewUseCase(repository vo.DeviceRepositoryInterface) *SendNotificationUseCase {
	return &SendNotificationUseCase{
		repository: repository,
	}
}

func (u *SendNotificationUseCase) SendNotification(input *DeviceInput) (string, error) {
	config, err := configs.LoadConfig(".env")
	if err != nil {
		log.Fatalln("USECASE::SEND NOTIFICATION::Error loading config ", err)
		return "", err
	}
	data, err := os.ReadFile(config.FIREBASE_AUTH_KEY)
	if err != nil {
		log.Fatalln("USECASE::SEND NOTIFICATION:: READ FIREBASE_AUTH_KEY:: ", err)
		return "", err
	}

	opt := option.WithCredentialsJSON([]byte(data))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return "", err
	}

	fcmClient, err := app.Messaging(context.Background())
	if err != nil {
		return "", err
	}

	// TODO: pegar token do monitor na tabela devices atraves do id do monitor
	response, err := fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: config.NOTIFICATION_TITLE,
			Body:  config.NOTIFICATION_BODY,
		},
		Token: "fjB8q3dXThul_qb6uH_gS8:APA91bFcep8Z0PJW16zcyYnG9lAlUo9PYpsiiptSdABozcsDmm1lwYkLFgxNdbOjTb7JpNlrjrjqFl6EbXdN6u5CJ_8u8rnt9QCqbynNmbTmJYqgqSwkfOgBEakF6uxb0Pls6NW7fC_O",
	})
	if err != nil {
		return "", err
	}

	return response, nil
}
