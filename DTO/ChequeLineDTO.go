package DTO

type ChequeLineDTO struct {
	Amount       float64     `xml:"Amount,omitempty"`
	ChequeLineID int64       `xml:"ChequeLineID,omitempty"`
	IsDeleted    bool        `xml:"IsDeleted,omitempty"`
	Item         *ItemRefDTO `xml:"Item,omitempty"`
	Price        float64     `xml:"Price,omitempty"`
	Quantity     float64     `xml:"Quantity,omitempty"`
}
