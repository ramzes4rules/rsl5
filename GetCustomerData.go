package methods

import (
	"encoding/xml"
	"fmt"
	"unibot/RSLoyatyWebService/DTO"
)

type GetCustomerData struct {
	XMLName xml.Name `xml:"GetCustomerData"`
	Token   string   `xml:"token,omitempty"`
}

type GetCustomerDataResponse struct {
	XMLName               xml.Name      `xml:"GetCustomerDataResponse"`
	GetCustomerDataResult *DTO.Customer `xml:"GetCustomerDataResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetCustomerData(config RSLoyaltyWebService, channel string, identifier int64, token string) *DTO.Customer {

	rsl.Init(config, channel, identifier)

	var request = &GetCustomerData{Token: token}
	var response = &GetCustomerDataResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetCustomerData", response)
	if err != nil {
		nestor.Error("", fmt.Sprintf("Ошибка: %v", err))
		return nil
	}

	return response.GetCustomerDataResult
}
