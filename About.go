package methods

import (
	"encoding/xml"
	"unibot/RSLoyatyWebService/DTO"
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

func (rsl *RSLoyaltyWebService) About() []DTO.About {

	var request = &AboutRequest{}
	var response = &AboutResponse{}

	var abouts []DTO.About
	_, err := rsl.Soap("", request, "About", response)
	if err != nil {
		//log.Println("Ошибка выполнения запроса")
		return abouts
	}
	//log.Println("Статус выполения запроса:", status)
	//log.Println(len(response.Result.Abouts))
	//log.Println(response.Result.Abouts[0].CompanyName)

	abouts = response.Result.Abouts
	return abouts
}
