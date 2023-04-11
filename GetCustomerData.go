package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetCustomerData struct {
	XMLName xml.Name `xml:"GetCustomerData"`
	Token   string   `xml:"token,omitempty"`
}

type GetCustomerDataResponse struct {
	XMLName               xml.Name      `xml:"GetCustomerDataResponse"`
	GetCustomerDataResult *DTO.Customer `xml:"GetCustomerDataResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetCustomerData(token string) (*DTO.Customer, error) {

	var request = &GetCustomerData{Token: token}
	var response = &GetCustomerDataResponse{}

	err := rsl.Soap("", request, "GetCustomerData", response)
	if err != nil {
		return nil, err
	}

	return response.GetCustomerDataResult, nil
}
