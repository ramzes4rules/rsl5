package methods

import (
	"encoding/xml"
)

type HasFreeVirtualDiscountCardsRequest struct {
	XMLName xml.Name `xml:"HasFreeVirtualDiscountCards"`
}

type HasFreeVirtualDiscountCardsResponse struct {
	XMLName xml.Name `xml:"HasFreeVirtualDiscountCardsResponse"`
	Result  bool     `xml:"HasFreeVirtualDiscountCardsResult"`
}

func (rsl RSLoyaltyWebService) HasFreeVirtualDiscountCards() (bool, error) {

	var request = &HasFreeVirtualDiscountCardsRequest{}
	var response = &HasFreeVirtualDiscountCardsResponse{}

	err := rsl.Soap("", request, "HasFreeVirtualDiscountCards", response)
	if err != nil {
		return false, err
	}

	return response.Result, nil

}
