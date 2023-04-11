package methods

import (
	"encoding/xml"
	"fmt"
	"log"
)

type SendCodeRequest struct {
	XMLName xml.Name `xml:"SendCode"`
	Phone   string   `xml:"phone"`
}

type SendCodeResponse struct {
	XMLName        xml.Name `xml:"SendCodeResponse"`
	SendCodeResult string   `xml:"SendCodeResult"`
}

func (rsl RSLoyaltyWebService) SendCode(channel string, identifier int64, phone string) string {

	//rsl.Init()
	var name = fmt.Sprintf("%s %d %s", channel, identifier, "SendCode")
	//var debug = true

	var request = &SendCodeRequest{Phone: phone}
	var response = &SendCodeResponse{}

	//var settings []LoyaltySettingDTO
	status, err := rsl.Soap2(channel, identifier, request, "SendCode", response)
	if err != nil {
		log.Printf("%s Ошибка выполнения запроса", name)
		return ""
	}

	if status == "200 OK" {
		return response.SendCodeResult
	} else {
		return ""
	}
}
