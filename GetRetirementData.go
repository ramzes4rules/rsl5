package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetRetirementData struct {
	XMLName xml.Name `xml:"GetRetirementData"`
	Token   string   `xml:"token,omitempty"`
}

type GetRetirementDataResponse struct {
	XMLName                 xml.Name                  `xml:"GetRetirementDataResponse"`
	GetRetirementDataResult *ArrayOfRetirementDataDTO `xml:"GetRetirementDataResult,omitempty"`
}

type ArrayOfRetirementDataDTO struct {
	RetirementData []*DTO.RetirementDataDTO `xml:"RetirementDataDTO,omitempty"`
}

func (rsl RSLoyaltyWebService) GetRetirementData(token string) ([]*DTO.RetirementDataDTO, error) {

	var request = &GetRetirementData{Token: token}
	var response = &GetRetirementDataResponse{}

	err := rsl.Soap("", request, "GetRetirementData", response)
	if err != nil {
		return nil, err
	}

	return response.GetRetirementDataResult.RetirementData, nil
}
