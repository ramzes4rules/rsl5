package methods

import (
	"encoding/xml"
	"fmt"
	"unibot/RSLoyatyWebService/DTO"
)

type GetIndicators struct {
	XMLName   xml.Name `xml:"GetIndicators"`
	Token     string   `xml:"token,omitempty"`
	CompanyID int64    `xml:"companyID,omitempty"`
}
type GetIndicatorsResponse struct {
	XMLName             xml.Name          `xml:"GetIndicatorsResponse"`
	GetIndicatorsResult *ArrayOfIndicator `xml:"GetIndicatorsResult,omitempty"`
}
type ArrayOfIndicator struct {
	Indicators []*DTO.IndicatorDTO `xml:"Indicator,omitempty"`
}

func (rsl RSLoyaltyWebService) GetIndicators(config RSLoyaltyWebService, channel string, identifier int64, token string) []*DTO.IndicatorDTO {

	rsl.Init(config, channel, identifier)

	var name = fmt.Sprintf(nestor.HeaderFormat, channel, identifier, "RSLoyaltyWebService.GetIndicators")

	var request = &GetIndicators{Token: token}
	var response = &GetIndicatorsResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "GetIndicators", response)
	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
	}

	return response.GetIndicatorsResult.Indicators

}
