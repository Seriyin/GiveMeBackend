package division

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Seriyin/GiveMeBackend/config/datastore"
	"github.com/Seriyin/GiveMeBackend/config/firebase"
	"github.com/Seriyin/GiveMeBackend/config/firebase/firestore"
	"github.com/Seriyin/GiveMeBackend/config/firebase/messaging"
	"github.com/Seriyin/GiveMeBackend/config/firebase/paths"
	"log"
	"strconv"
	"strings"
)

var db = firebase.GetDB()
var mesClient = firebase.GetMessaging()

func Division(
	ctx context.Context,
	e firestore.Event,
) error {
	var groupT datastore.GroupRequest
	err := json.Unmarshal(e.Value.Fields, &groupT)

	log.Print("Attempted unmarshal")
	if err != nil {
		return err
	}

	newAmountUnit, newAmountCents, err := calculateResultingAmounts(&groupT)

	log.Print("Attempted division")
	if err != nil {
		return err
	}

	monPath := paths.TransformGroupIntoMonetary(e.Value.Name)

	log.Printf("Transformed path: %v", monPath)
	monetaryTs := extractIndividualTos(
		ctx,
		monPath,
		&groupT,
		newAmountUnit,
		newAmountCents,
	)

	dbPath := paths.ExtractMethodIdAndDatePath(monPath)

	log.Printf("Extracted db path: %v", dbPath)
	err = db.SetMonetaryRequestsByFullPath(
		ctx,
		monetaryTs,
		dbPath,
	)

	return err
}

func extractIndividualTos(
	ctx context.Context,
	networkPath string,
	groupT *datastore.GroupRequest,
	newAmountUnit int64,
	newAmountCents int64,
) []*datastore.MonetaryRequest {
	monetaryTs := make(
		[]*datastore.MonetaryRequest,
		len(groupT.Tos),
	)

	for _, to := range groupT.Tos {
		m := &datastore.MonetaryRequest{
			From:          groupT.From,
			To:            to,
			Desc:          groupT.Desc,
			Date:          groupT.Date,
			AmountUnit:    newAmountUnit,
			AmountCents:   newAmountCents,
			Currency:      "€", //missing from grouptransfer
			ConfirmedFrom: false,
			ConfirmedTo:   false,
			GroupId:       groupT.GroupId,
			RecurrentId:   -1,
		}
		monetaryTs = append(monetaryTs, m)

		// If no profile can be gathered, the user may not exist.
		// Either by network error or profile not existing, must skip.
		profile, err := db.GetProfileByPhoneNumber(ctx, to)
		if err == nil {
			dbPath := paths.ExtractAndReplaceMethodIdAndDatePath(profile.Id, networkPath)

			//Add has side-effects and assigns a LinkedId to monetary transfer which
			//is the first return param.
			_, err = db.AddMonetaryRequestByFullPath(ctx, m, dbPath)

			if err == nil {
				err = produceAndSendNotification(
					ctx,
					profile,
					m,
				)
				if err != nil {
					log.Print(err)
				}
			} else {
				log.Print(err)
			}

		} else {
			log.Print(err)
		}
		//skip otherwise
	}
	return monetaryTs
}

func calculateResultingAmounts(groupT *datastore.GroupRequest) (int64, int64, error) {
	totalValue := groupT.AmountUnit*100 + groupT.AmountCents

	var newAmountUnit int64
	var newAmountCents int64
	var err error
	if groupT.Included {
		newAmountUnit, newAmountCents, err = extractDivide(
			totalValue,
			int64(len(groupT.Tos)+1),
		)
	} else {
		newAmountUnit, newAmountCents, err = extractDivide(
			totalValue,
			int64(len(groupT.Tos)),
		)
	}
	if err != nil {
		return -1, -1, err
	}
	return newAmountUnit, newAmountCents, nil
}

func extractDivide(totalValue int64, groupNum int64) (int64, int64, error) {
	dividedvalue := float64(totalValue) / float64(groupNum)
	st := fmt.Sprintf("%.2f", dividedvalue)
	sp := strings.Split(st, ".")
	newAmountUnit, err := strconv.ParseInt(sp[0], 10, 64)
	if err != nil {
		return -1, -1, err
	}
	newAmountCents, err := strconv.ParseInt(sp[1], 10, 64)
	if err != nil {
		return -1, -1, err
	}
	return newAmountUnit, newAmountCents, err
}

func produceAndSendNotification(
	ctx context.Context,
	profile *datastore.Profile,
	transfer *datastore.MonetaryRequest,
) error {
	//generate notification message
	token := profile.Token
	message := messaging.GenerateRequestNotification(
		token,
		transfer.AmountUnit,
		transfer.AmountCents,
		transfer.Currency,
		transfer.From,
	)

	str, err := mesClient.Send(ctx, message)
	if err != nil {
		log.Print(str)
	}
	return err
}
