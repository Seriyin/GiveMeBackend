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
)

var db = firebase.GetDB()

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
			From: "IMPLEMENTE ME",
			To: to,
			Desc: ,
			Date: ,
			AmountUnit: ,
			AmountCents: ,
			Currency: ,
			ConfirmedFrom: ,
			ConfirmedTo: ,
			GroupId: ,
			RecurrentId: ,
		}
		monetaryTs = append(monetaryTs, m)
	} 
}
