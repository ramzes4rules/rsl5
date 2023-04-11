package DTO

import "encoding/xml"

type LoyaltySettingDTO struct {
	XMLName xml.Name `xml:"LoyaltySettingDTO"`
	//xmlns:a="http://schemas.datacontract.org/2004/07/RS.Loyalty.WebClientPortal.Core.Model" xmlns:i="http://www.w3.org/2001/XMLSchema-instance
	KeyValue         string `xml:"KeyValue"`
	LoyaltySettingID int    `xml:"LoyaltySettingID"`
	Type             string `xml:"Type"`
	Value            string `xml:"Value"`
}
