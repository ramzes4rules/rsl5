package methods

import "encoding/xml"

type DiscountCardIsExist struct {
	XMLName            xml.Name `xml:"DiscountCardIsExist"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type DiscountCardIsExistResponse struct {
	XMLName                   xml.Name `xml:"DiscountCardIsExistResponse"`
	DiscountCardIsExistResult bool     `xml:"DiscountCardIsExistResult,omitempty"`
}

func (rsl RSLoyaltyWebService) DiscountCardIsExist(DiscountCardNumber string) (bool, error) {

	var request = &DiscountCardIsExist{DiscountCardNumber: DiscountCardNumber}
	var response = &DiscountCardIsExistResponse{}

	err := rsl.Soap("", request, "DiscountCardIsExist", response)
	if err != nil {
		return false, err
	}

	return response.DiscountCardIsExistResult, nil
}
