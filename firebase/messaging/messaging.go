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

func generateRequestRefusal(
	token string,
	amountUnits int64,
	amountCents int32,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debtor Refused Debt Payment Request",
				Body: fmt.Sprintf(
					"%v Refused the Debt of %v.%v %v",
					deliveredFrom,
					amountUnits,
					amountCents,
					currency,
				),
				Color: "#161119",
			},
			RestrictedPackageName: "com.giveme.pei.givemeapp",
		},
		Token: token,
	}
}

func generateRequestAcceptance(
	token string,
	amountUnits int64,
	amountCents int32,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debtor Accepted Debt Payment Request",
				Body: fmt.Sprintf(
					"%v Accepted the Debt of %v.%v %v",
					deliveredFrom,
					amountUnits,
					amountCents,
					currency,
				),
				Color: "#161119",
			},
			RestrictedPackageName: "com.giveme.pei.givemeapp",
		},
		Token: token,
	}
}

func generateReminder(
	token string,
	amountUnites int64,
	amountCents int32,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Remind Debtor of a Debt Payment Request",
				Body: fmt.Sprintf(
					"You have %v.%v %v to pay to %v",
					amountUnits,
					amountCents,
					currency
					deliveredFrom,
				),
				Color: "#161119",
			},
			RestrictedPackageName: "com.giveme.pei.givemeapp",
		},
		Token: token,
	}
}
