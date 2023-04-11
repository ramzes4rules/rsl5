package DTO

type ChequeDTO struct {
	AccrualBonus    float64                `xml:"AccrualBonus,omitempty"`
	Amount          float64                `xml:"Amount,omitempty"`
	ChequeID        int64                  `xml:"ChequeID,omitempty"`
	ChequeLines     *ArrayOfChequeLine     `xml:"ChequeLines,omitempty"`
	ChequeType      *ChequeType            `xml:"ChequeType,omitempty"`
	ChequeUID       string                 `xml:"ChequeUID,omitempty"`
	DiscountCard    *DiscountCardDTO       `xml:"DiscountCard,omitempty"`
	Discounts       *ArrayOfChequeDiscount `xml:"Discounts,omitempty"`
	IsDeleted       bool                   `xml:"IsDeleted,omitempty"`
	PositionCount   int32                  `xml:"PositionCount,omitempty"`
	StoreID         int64                  `xml:"StoreID,omitempty"`
	StoreName       string                 `xml:"StoreName,omitempty"`
	SubtractedBonus float64                `xml:"SubtractedBonus,omitempty"`
	//Time time.Time `xml:"Time,omitempty"`
	Time string `xml:"Time,omitempty"`
}

type ArrayOfChequeLine struct {
	ChequeLine []*ChequeLineDTO `xml:"ChequeLineDTO,omitempty"`
}

type ArrayOfChequeDiscount struct {
	ChequeDiscount []*ChequeDiscountDTO `xml:"ChequeDiscountDTO,omitempty"`
}

type ChequeType string

const (
	ChequeTypeSale   ChequeType = "Sale"
	ChequeTypeReturn ChequeType = "Return"
)

type DiscountType string

const (
	DiscountTypePercent  DiscountType = "Percent"
	DiscountTypeAmount   DiscountType = "Amount"
	DiscountTypeTender   DiscountType = "Tender"
	DiscountTypeFixPrice DiscountType = "FixPrice"
)
