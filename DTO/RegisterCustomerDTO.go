package DTO

type RegisterCustomerDTO struct {
	Xmlna              string                        `xml:"xmlns:a,attr"`
	Xmlni              string                        `xml:"xmlns:i,attr"`
	CanRegister        bool                          `xml:"a:CanRegister,omitempty"`
	CreateVirtual      bool                          `xml:"a:CreateVirtual,omitempty"`
	DiscountCardNumber string                        `xml:"a:DiscountCardNumber,omitempty"`
	FirstName          string                        `xml:"a:FirstName,omitempty"`
	HasCards           bool                          `xml:"a:HasCards,omitempty"`
	LastName           string                        `xml:"a:LastName,omitempty"`
	Login              string                        `xml:"a:Login,omitempty"`
	NeedFirstName      bool                          `xml:"a:NeedFirstName,omitempty"`
	NeedLastName       bool                          `xml:"a:NeedLastName,omitempty"`
	NeedSecondName     bool                          `xml:"a:NeedSecondName,omitempty"`
	Password           string                        `xml:"a:Password,omitempty"`
	PinCode            string                        `xml:"a:PinCode,omitempty"`
	Properties         *ArrayOfCustomerPropertyValue `xml:"a:Properties,omitempty"`
	SecondName         string                        `xml:"a:SecondName,omitempty"`
	TermsOfUsing       *ArrayOfKeyValueOflongboolean `xml:"a:TermsOfUsing,omitempty"`
}
