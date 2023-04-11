package methods

import (
	"encoding/xml"
	"log"
	"unibot/RSLoyatyWebService/DTO"
)

type GetRetirementData struct {
	XMLName xml.Name `xml:" GetRetirementData"`
	Token   string   `xml:"token,omitempty"`
}

type GetRetirementDataResponse struct {
	XMLName                 xml.Name                  `xml:" GetRetirementDataResponse"`
	GetRetirementDataResult *ArrayOfRetirementDataDTO `xml:"GetRetirementDataResult,omitempty"`
}

type ArrayOfRetirementDataDTO struct {
	RetirementData []*DTO.RetirementDataDTO `xml:"RetirementDataDTO,omitempty"`
}

func (rsl RSLoyaltyWebService) GetRetirementData(config RSLoyaltyWebService, channel string, identifier int64, token string) []*DTO.RetirementDataDTO {

	rsl.Init(config, channel, identifier)

	var request = &GetRetirementData{Token: token}
	var response = &GetRetirementDataResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetRetirementData", response)
	if err != nil {
		log.Println("Ошибка выполнения запроса", err)
		//return settings
	}

	return response.GetRetirementDataResult.RetirementData
}
