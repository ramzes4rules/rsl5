package methods

import (
	"encoding/xml"
	"unibot/RSLoyatyWebService/DTO"
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

func (rsl RSLoyaltyWebService) GetCustomersProperties() []*DTO.Property {

	//rsl.Init(config, channel, identifier)

	//var name = fmt.Sprintf(lhm, channel, identifier, "RSLoyaltyWebService.GetCustomersProperties")

	var request = &GetCustomersProperties{}
	var response = &GetCustomersPropertiesResponse{}

	status, err := rsl.Soap2("", 0, request, "GetCustomersProperties", response)
	if err != nil {
		//nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
	}

	if status == "200 OK" {
		//nestor.Debug(name, fmt.Sprintf("Получено Properties: %d", len(response.GetCustomersPropertiesResult.Property)))
		return response.GetCustomersPropertiesResult.Property

	}
	return nil
}
