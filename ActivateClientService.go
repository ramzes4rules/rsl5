package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type ActivateClientService struct {
	XMLName  xml.Name                `xml:"ActivateClientService"`
	UserData DTO.RegisterCustomerDTO `xml:"userData,omitempty"`
	Info     string                  `xml:"info,omitempty"`
}

type ActivateClientServiceResponse struct {
	XMLName                     xml.Name `xml:"ActivateClientServiceResponse"`
	ActivateClientServiceResult string   `xml:"ActivateClientServiceResult,omitempty"`
	Info                        string   `xml:"info,omitempty"`
}

func (rsl RSLoyaltyWebService) ActivateClientService(userData DTO.RegisterCustomerDTO, info string) (string, string, error) {
	var request = &ActivateClientService{UserData: userData, Info: info}
	var response = &ActivateClientServiceResponse{}
	err := rsl.Soap("", request, "ActivateClientService", response)
	if err != nil {
		return "", "", err
	}
	return response.ActivateClientServiceResult, response.Info, nil
}
