package datastore

type Profile struct {
	UserId   UID      `firestore:"userid"`
	Metadata Metadata `firestore:"metadata"`
}

type UID struct {
	Id     string `firestore:"id"`
	Device []byte `firestore:"device"`
	PubKey []byte `firestore:"pubkey"`
	Token  []byte `firestore:"token"`
	//	SignedPreKey io.ReadWriter
	//	PreKeyBundle []io.ReadWriter
}

type Metadata struct {
	PaymentProviders []PaymentProvider `firestore:"paymentproviders"`
	NumberPayments   uint64            `firestore:"numberpayments"`
}
