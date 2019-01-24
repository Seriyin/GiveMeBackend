package datastore

type Profile struct {
	UserId   UID      `firestore:"user_id" json:"user_id"`
	Metadata Metadata `firestore:"metadata" json:"metadata"`
}

type UID struct {
	//Id should be phone number.
	Id     string `firestore:"id" json:"id"`
	Device []byte `firestore:"device" json:"device"`
	PubKey []byte `firestore:"pub_key" json:"pub_key"`
	Token  []byte `firestore:"token" json:"token"`
	Name   string `firestore:"name" json:"name"`
	Email  string `firestore:"email" json:"email"`
	//	SignedPreKey io.ReadWriter
	//	PreKeyBundle []io.ReadWriter
}

type Metadata struct {
	PaymentProviders []PaymentProvider `firestore:"payment_providers" json:"payment_providers"`
	NumberPayments   uint64            `firestore:"number_payments" json:"number_payments"`
}
