package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetCustomersProperties struct {
	XMLName          xml.Name `xml:"GetCustomersProperties"`
	NeedIncludeEmail bool     `xml:"needIncludeEmail,omitempty"`
	EditOnly         bool     `xml:"editOnly,omitempty"`
}

type GetCustomersPropertiesResponse struct {
	XMLName                      xml.Name             `xml:"GetCustomersPropertiesResponse"`
	GetCustomersPropertiesResult *DTO.ArrayOfProperty `xml:"GetCustomersPropertiesResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetCustomersProperties() ([]*DTO.PropertyGet, error) {

	var request = &GetCustomersProperties{}
	var response = &GetCustomersPropertiesResponse{}

	err := rsl.Soap("", request, "GetCustomersProperties", response)
	if err != nil {
		return nil, err
	}

	return response.GetCustomersPropertiesResult.Property, nil

}
