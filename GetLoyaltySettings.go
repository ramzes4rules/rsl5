package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
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

func (rsl RSLoyaltyWebService) GetLoyaltySettings() ([]DTO.LoyaltySettingDTO, error) {

	var request = &GetLoyaltySettingsRequest{}
	var response = &GetLoyaltySettingsResponse{}

	var settings []DTO.LoyaltySettingDTO
	err := rsl.Soap("", request, "GetLoyaltySettings", response)
	if err != nil {
		return settings, err
	}

	return response.GetLoyaltySettingsResult.Result, nil
}
