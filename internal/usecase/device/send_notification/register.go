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

	deviceFound, err := u.repository.FindByMonitorID(input.MonitorID)
	if err != nil {
		return "", err
	}

	if deviceFound.Token == "" {
		return "", nil
	}

	response, err := fcmClient.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: config.NOTIFICATION_TITLE,
			Body:  config.NOTIFICATION_BODY,
		},
		Token: deviceFound.Token,
	})
	if err != nil {
		return "", err
	}
	return response, nil

}
