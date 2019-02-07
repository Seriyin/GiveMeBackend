package request

import (
	"strings"
	"json"
	"context"
	"github.com/Seriyin/GibMe-backend/config/firebase"
	"github.com/Seriyin/GibMe-backend/config/firebase/firestore"
)

var db = firebase.GetDB()

func Request(
	ctx context.Context,
	e firestore.Event,
) {
	var moneteryT MonetaryTransfer //Monetary Structure object
	err := json.Unmarshal(e.Value.Fields, &moneteryT) //Json obejct to Monetary Structure
	debtor := moneteryT.To //Debtor present in the To value on the Monetary Stuct

	/*
	*Use the debotr id to split the path
	*so it gets the path next to the credor ID 
	*to be replaced by the debtor ID, so it inserts on the correct collection
	*/
	splitvalue := "/"+moneteryT.From+"/"
	path := strings.Split(e.Value.Name, splivalue)[1]

	db.setMonetaryTransfer(debtor, newmonetaryT, path)
}
