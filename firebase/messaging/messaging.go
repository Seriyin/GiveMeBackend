package messaging

import (
	"firebase.google.com/go/messaging"
	"fmt"
)

var ()

func generateRequestNotification(
	token string,
	amountUnits int64,
	amountCents int32,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debt Notification",
				Body: fmt.Sprintf(
					"You owe %v.%v to %v",
					amountUnits,
					amountCents,
					deliveredFrom,
				),
				Color: "#161119",
			},
			RestrictedPackageName: "com.giveme.pei.givemeapp",
		},
		Token: token,
	}
}

func generateNewUserRequest(
	toCheckRemind string,
	token string,
	toRemind string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "User Joined Notification",
				Body: fmt.Sprintf(
					"%v, A user who owes you has joined GiveMe. Remind %v them of their debt?",
					toCheckRemind,
					toRemind,
				),
				Color: "#161119",
			},
			RestrictedPackageName: "com.giveme.pei.givemeapp",
		},
		Token: token,
	}
}
