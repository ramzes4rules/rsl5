package DTO

type MailingToCustomerDTO struct {
	//DateAdded time.Time `xml:"DateAdded,omitempty"`
	DateAdded           string `xml:"DateAdded,omitempty"`
	IsRead              bool   `xml:"IsRead,omitempty"`
	MailingToCustomerID int64  `xml:"MailingToCustomerID,omitempty"`
	Message             string `xml:"Message,omitempty"`
	//SendTime time.Time `xml:"SendTime,omitempty"`
	SendTime string `xml:"SendTime,omitempty"`
	Subject  string `xml:"Subject,omitempty"`
}
