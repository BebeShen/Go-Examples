package expo

import (
	"fmt"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func SendPushNotificationToUser()  {
	pushToken, err := expo.NewExponentPushToken("ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]")
	if err != nil{
		panic(err)
	}

	client := expo.NewPushClient(nil)

	res, err := client.Publish(
		&expo.PushMessage{
			To:	[]expo.ExponentPushToken{pushToken},
			Body: "This is a test notification",
            Data: map[string]string{"withSome": "data"},
            Sound: "default",
            Title: "Notification Title",
            Priority: expo.DefaultPriority,
		},
	)
	if err != nil {
		panic(nil)
	}
	// Validate responses
    if res.ValidateResponse() != nil {
        fmt.Println(res.PushMessage.To, "failed")
    }
}