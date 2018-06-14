package notifier

import (
	"github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	"errors"
)

var client *expo.PushClient

func init() {
	// Create a new Expo SDK client
	client = expo.NewPushClient(nil)
}

func Push(token, message, page string) error {
	// To check the token is valid
	pushToken, err := expo.NewExponentPushToken("ExponentPushToken["+token+"]")
	if err != nil {
		return err
	}

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To: pushToken,
			Body: message,
			Data: map[string]string{"redirect": page},
			Sound: "default",
			Title: "Pustakalaya",
			Priority: expo.DefaultPriority,
		},
	)

	// Check errors
	if err != nil {
		return err
	}
	// Validate responses
	if response.ValidateResponse() != nil {
		return errors.New("failed to publish")
	}
}