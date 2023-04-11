package DTO

import (
	"encoding/xml"
)

type IndicatorDTO struct {
	XMLName         xml.Name    `xml:"Indicator"`
	BonusAmount     float64     `xml:"BonusAmount,omitempty"`
	Company         *CompanyDTO `xml:"Company,omitempty"`
	IndicatorID     int64       `xml:"IndicatorID,omitempty"`
	PayBonuses      float64     `xml:"PayBonuses,omitempty"`
	SavedByDiscount float64     `xml:"SavedByDiscount,omitempty"`
	SavedByPayBonus float64     `xml:"SavedByPayBonus,omitempty"`
}
