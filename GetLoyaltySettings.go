package methods

import (
	"encoding/xml"
	"log"
	"unibot/RSLoyatyWebService/DTO"
)

type GetLoyaltySettingsRequest struct {
	XMLName xml.Name `xml:"GetLoyaltySettings"`
}
type GetLoyaltySettingsResponse struct {
	XMLName                  xml.Name                 `xml:"GetLoyaltySettingsResponse"`
	GetLoyaltySettingsResult GetLoyaltySettingsResult `xml:"GetLoyaltySettingsResult"`
}
type GetLoyaltySettingsResult struct {
	XMLName xml.Name                `xml:"GetLoyaltySettingsResult"`
	Xmlnsa  string                  `xml:"xmlns:a,attr"`
	Xmlnsi  string                  `xml:"xmlns:i,attr"`
	Result  []DTO.LoyaltySettingDTO `xml:"LoyaltySettingDTO"`
}

func (rsl RSLoyaltyWebService) GetLoyaltySettings() []DTO.LoyaltySettingDTO {

	//rsl.Init()

	var request = &GetLoyaltySettingsRequest{}
	var response = &GetLoyaltySettingsResponse{}
	//response.GetLoyaltySettingsResult.Xmlnsa = "http://schemas.datacontract.org/2004/07/RS.Loyalty.WebClientPortal.Core.Model"
	//response.GetLoyaltySettingsResult.Xmlnsi = "http://www.w3.org/2001/XMLSchema-instance"

	var settings []DTO.LoyaltySettingDTO
	status, err := rsl.Soap2("", 0, request, "GetLoyaltySettings", response)
	if err != nil {
		log.Println("Ошибка выполнения запроса")
		return settings
	}

	if status == "200 OK" {
		settings = response.GetLoyaltySettingsResult.Result
	}

	return settings
}
