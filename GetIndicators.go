package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
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

func (rsl RSLoyaltyWebService) GetIndicators(token string) ([]*DTO.IndicatorDTO, error) {

	var request = &GetIndicators{Token: token}
	var response = &GetIndicatorsResponse{}

	err := rsl.Soap("", request, "GetIndicators", response)
	if err != nil {
		return nil, err
	}

	return response.GetIndicatorsResult.Indicators, nil

}
