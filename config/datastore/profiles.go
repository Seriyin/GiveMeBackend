package datastore

type Profile struct {
	UID
	Metadata
}

type UID struct {
	//Id should be auth token.
	Id     string `firestore:"id" json:"id"`
	Phone  string `firestore:"phone" json:"phone"`
	Device []byte `firestore:"device" json:"device"`
	PubKey []byte `firestore:"pubKey" json:"pubKey"`
	Token  string `firestore:"token" json:"token"`
	Name   string `firestore:"name" json:"name"`
	Email  string `firestore:"email" json:"email"`
	//	SignedPreKey io.ReadWriter
	//	PreKeyBundle []io.ReadWriter
}

type Metadata struct {
	PaymentProviders []PaymentProvider `firestore:"paymentProviders" json:"paymentProviders"`
	NumberPayments   int64             `firestore:"numberPayments" json:"numberPayments"`
}
