package DTO

type CompanyDTO struct {
	AboutText                 string  `xml:"AboutText,omitempty"`
	BonusRate                 float64 `xml:"BonusRate,omitempty"`
	CompanyID                 int64   `xml:"CompanyID,omitempty"`
	CompanyUID                string  `xml:"CompanyUID,omitempty"`
	CurrencySign              string  `xml:"CurrencySign,omitempty"`
	DiscountAccuracy          int32   `xml:"DiscountAccuracy,omitempty"`
	ItemFavoriteCategoryCount int32   `xml:"ItemFavoriteCategoryCount,omitempty"`
	Name                      string  `xml:"Name,omitempty"`
	TermsOfUseText            string  `xml:"TermsOfUseText,omitempty"`
}
