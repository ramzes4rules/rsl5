package DTO

type ChequeDiscountDTO struct {
	Amount           float64        `xml:"Amount,omitempty"`
	ChequeDiscountID int64          `xml:"ChequeDiscountID,omitempty"`
	ChequeLine       *ChequeLineDTO `xml:"ChequeLineDTO,omitempty"`
	DiscountType     *DiscountType  `xml:"DiscountType,omitempty"`
}
