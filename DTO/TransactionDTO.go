package DTO

import (
	"encoding/xml"
)

type TransactionDTO struct {
	XMLName       xml.Name    `xml:"Transaction"`
	AccrualAmount float64     `xml:"AccrualAmount,omitempty"`
	AccrualBonus  float64     `xml:"AccrualBonus,omitempty"`
	AccrualVisit  int32       `xml:"AccrualVisit,omitempty"`
	ChequeType    *ChequeType `xml:"ChequeType,omitempty"`
	CompanyName   string      `xml:"CompanyName,omitempty"`
	//DelayUntil time.Time `xml:"DelayUntil,omitempty"`
	DelayUntil           string `xml:"DelayUntil,omitempty"`
	DiscountCardID       int32  `xml:"DiscountCardID,omitempty"`
	DiscountCardNumber   string `xml:"DiscountCardNumber,omitempty"`
	RelatedTransactionID int64  `xml:"RelatedTransactionID,omitempty"`
	TransactionID        int64  `xml:"TransactionID,omitempty"`
	//TransactionTime time.Time `xml:"TransactionTime,omitempty"`
	TransactionTime string `xml:"TransactionTime,omitempty"`
	//TransactionType *TransactionType `xml:"TransactionType,omitempty"`
	TransactionType string `xml:"TransactionType,omitempty"`
	//ValidUntil time.Time `xml:"ValidUntil,omitempty"`
	ValidUntil string `xml:"ValidUntil,omitempty"`
}

type TransactionType string

const (
	TransactionTypeCHQ TransactionType = "CHQ"
	TransactionTypeBNS TransactionType = "BNS"
	TransactionTypeRTM TransactionType = "RTM"
	TransactionTypeDEC TransactionType = "DEC"
	TransactionTypeFIX TransactionType = "FIX"
	TransactionTypeGFT TransactionType = "GFT"
)
