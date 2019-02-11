package messaging

import (
	"fmt"

	"firebase.google.com/go/messaging"
)

func GenerateRequestNotification(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debt Notification",
				Body: fmt.Sprintf(
					"You were tagged to pay %v.%v %v to %v",
					amountUnits,
					amountCents,
					currency,
					deliveredFrom,
				),
				Color: "#161119",
			},
			RestrictedPackageName: "com.giveme.pei.givemeapp",
		},
		Token: token,
	}
}

func GenerateNewUserRequest(
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
					"%v, A user who owes you has joined GiveMe. Remind %v of their debt?",
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

func GenerateRequestRefusal(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debtor Refused Debt Payment Request",
				Body: fmt.Sprintf(
					"%v refused the debt of %v.%v %v",
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

func GenerateRequestAcceptance(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debtor Accepted Debt Payment Request",
				Body: fmt.Sprintf(
					"%v accepted the debt of %v.%v %v",
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

func GenerateReminder(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
	deliveredFrom string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Debt Payment Reminder",
				Body: fmt.Sprintf(
					"%v wants to remind you to pay %v.%v %v",
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

func GenerateScheduled(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Scheduled Debt Payment",
				Body: fmt.Sprintf(
					"A debt was scheduled of %v.%v%v",
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

func GenerateConfirmedFromNotification(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
	from string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Creditor confirmed payment",
				Body: fmt.Sprintf(
					"%v received your payment of %v.%v%v",
					from,
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

func GenerateConfirmedToNotification(
	token string,
	amountUnits int64,
	amountCents int64,
	currency string,
	to string,
) *messaging.Message {
	return &messaging.Message{
		Android: &messaging.AndroidConfig{
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: "Creditor confirmed payment",
				Body: fmt.Sprintf(
					"%v payed a debt of %v.%v%v",
					from,
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
