package division

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Seriyin/GibMe-backend/config/datastore"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
	"github.com/Seriyin/GibMe-backend/config/firebase/messaging"
)

var db = firebase.GetDB()
var mesClient = firebase.GetMessaging()

func Division(
	ctx context.Context,
	e firestore.Event,
) error {
	var groupT datastore.GroupTransfer
	time := time.Now()
	err := json.Unmarshal(e.Value.Fields, &groupT)
	if err != nil {
		return err
	}
	totalvalue := groupT.AmountUnit*100 + groupT.AmountCents
	var dividedvalue float64
	var newAmountUnit int64
	var newAmountCents int64
	if groupT.Included {
		dividedvalue = float64(totalvalue/(len(groupT.Tos)+1))
		st := fmt.Sprintf("%f", dividedvalue)
		sp := strings.Split(a, ".")
		newAmountUnit = strconv.ParseInt(sp[0], 10, 64)
		newAmountCents = strconv.ParseInt(sp[1], 10, 64)
	}
	else {
		dividedvalue = float64(totalvalue/(len(groupT.Tos))
		st := fmt.Sprintf("%f", dividedvalue)
		sp := strings.Split(a, ".")
		newAmountUnit = strconv.ParseInt(sp[0], 10, 64)
		newAmountCents = strconv.ParseInt(sp[1], 10, 64)
	}
	monetaryTs := make([]*datastore.MonetaryTransfer, len(groupT.Tos)) 
	for _, to := range groupT.Tos {
		m := datastore.MonetaryTransfer{
			From: "IMPLEMENT ME", //missing from grouptransfer
			To: to,
			Desc: "Debt Division",
			Date: time,
			AmountUnit: newAmountUnit,
			AmountCents: newAmountCents,
			Currency: "IMPLEMENT ME", //missing from grouptransfer
			ConfirmedFrom: false,
			ConfirmedTo: false,
			GroupId: groupT.GroupId,
			RecurrentId: ,
		}
		monetaryTs = append(monetaryTs, m)

		path := "CREATE PATH" //IMPLEMENT
		_, err := firestore.SetMonetaryTransfer(to,&m,path)
		if err != nil {
			return err
		}

		//generate notification message
		token := "" //IMPLEMENT ME
		message := messaging.GenerateRequestNotification(
			token,
			m.AmountUnit,
			m.AmountCents,
			m.Currency,
			m.From,
		)

		var str string
		str, err = mesClient.Send(ctx, message)
		if err != nil {
			log.Print(str)
		}
		return err
	} 
	path := "CREATE PATH" //IMPLEMENT
	err := firestore.SetMonetaryTransfers(
		"MISSING FROM from grouptransfer",
		&monetaryTs,
		path,
	)
}
