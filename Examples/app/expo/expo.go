package expo

import (
	"fmt"
	
    . "app/pkg"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func SendPushNotificationToUser(tokens_msg TokensWithMessage)  {
	// pushToken, err := expo.NewExponentPushToken("ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]")
	/*
	client := expo.NewPushClient(nil)
	for _, tm := range tms {
		pushToken, err := expo.NewExponentPushToken(tm.Token)
		if err != nil{
			panic(err)
		}
		res, err := client.Publish(
			&expo.PushMessage{
				To:	[]expo.ExponentPushToken{pushToken},
				Body: tm.Message,
				Data: map[string]string{"withSome": "data"},
				Sound: "default",
				Title: "Ainimal 通知",
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
	*/

	client := expo.NewPushClient(nil)
	var pushTokenList []expo.ExponentPushToken
	for _, t := range tokens_msg.Token {
		pushToken, err := expo.NewExponentPushToken(t)
		if err != nil{
			panic(err)
		}
		pushTokenList = append(pushTokenList, pushToken)
	}
	res, err := client.Publish(
		&expo.PushMessage{
			To:	pushTokenList,
			Body: tokens_msg.Message,
			Data: map[string]string{"withSome": "data"},
			Sound: "default",
			Title: "Ainimal 通知",
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