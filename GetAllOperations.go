package methods

import (
	"encoding/xml"
	"unibot/RSLoyatyWebService/DTO"
)

type GetAllOperations struct {
	XMLName xml.Name `xml:"GetAllOperations"`
	Token   string   `xml:"token,omitempty"`
}
type GetAllOperationsResponse struct {
	XMLName                xml.Name       `xml:"GetAllOperationsResponse"`
	GetAllOperationsResult *ArrayOfCheque `xml:"GetAllOperationsResult,omitempty"`
}
type ArrayOfCheque struct {
	Cheque []*DTO.ChequeDTO `xml:"Cheque,omitempty"`
}

func (rsl RSLoyaltyWebService) GetAllOperations(config RSLoyaltyWebService, channel string, identifier int64, token string) []*DTO.ChequeDTO {

	rsl.Init(config, channel, identifier)

	var request = &GetAllOperations{Token: token}
	var response = &GetAllOperationsResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetAllOperations", response)
	if err != nil {
		//log.Println("Ошибка выполнения запроса")
		//return settings
	}

	return response.GetAllOperationsResult.Cheque
}
