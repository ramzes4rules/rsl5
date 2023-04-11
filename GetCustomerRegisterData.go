package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetCustomerRegisterDataRequest struct {
	XMLName      xml.Name `xml:"GetCustomerRegisterData"`
	DiscountCard string   `xml:"number"`
}

type GetCustomerRegisterDataResponse struct {
	XMLName                       xml.Name                `xml:"GetCustomerRegisterDataResponse"`
	GetCustomerRegisterDataResult DTO.RegisterCustomerDTO `xml:"GetCustomerRegisterDataResult"`
}

func (rsl RSLoyaltyWebService) GetCustomerRegisterData(discountCard string) (*GetCustomerRegisterDataResponse, error) {

	var request = &GetCustomerRegisterDataRequest{DiscountCard: discountCard}
	var response = &GetCustomerRegisterDataResponse{}

	err := rsl.Soap("", request, "GetCustomerRegisterData", response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
