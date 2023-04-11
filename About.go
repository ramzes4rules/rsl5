package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type AboutRequest struct {
	XMLName xml.Name `xml:"About"`
}
type AboutResponse struct {
	XMLName xml.Name    `xml:"AboutResponse"`
	Result  AboutResult `xml:"AboutResult"`
}
type AboutResult struct {
	XMLName xml.Name    `xml:"AboutResult"`
	Abouts  []DTO.About `xml:"AboutDTO"`
}

func (rsl RSLoyaltyWebService) About() ([]DTO.About, error) {

	var request = &AboutRequest{}
	var response = &AboutResponse{}

	err := rsl.Soap("", request, "About", response)
	if err != nil {
		return []DTO.About{}, err
	}

	return response.Result.Abouts, nil
}
