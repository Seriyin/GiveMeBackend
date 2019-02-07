package datastore

//PaypalProvider is a PaymentProvider
var _ PaymentProvider = &PaypalProvider{}

type PaymentProvider interface {
	GetId() string
	GetUserId() string
	GetDetails() interface{}
}

type PaypalProvider struct {
	Id    string `firestore:"id" json:"id"`
	Email string `firestore:"email" json:"email"`
}

func (p *PaypalProvider) GetId() string {
	return p.Id
}

func (p *PaypalProvider) GetUserId() string {
	return p.Email
}

func (p *PaypalProvider) GetDetails() interface{} {
	return nil
}

type DefaultProvider struct {
	Id   string `firestore:"id" json:"id"`
	IBAN string `firestore:"iban"´json:"iban"`
}

func (d *DefaultProvider) GetId() string {
	return d.Id
}

func (d *DefaultProvider) GetUserId() string {
	return d.IBAN
}

func (d *DefaultProvider) GetDetails() interface{} {
	return nil
}
