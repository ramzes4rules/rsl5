package methods

import (
	"encoding/xml"
)

type SendCodeRequest struct {
	XMLName xml.Name `xml:"SendCode"`
	Phone   string   `xml:"phone"`
}

type SendCodeResponse struct {
	XMLName        xml.Name `xml:"SendCodeResponse"`
	SendCodeResult string   `xml:"SendCodeResult"`
}

func (rsl RSLoyaltyWebService) SendCode(phone string) (string, error) {
	var request = &SendCodeRequest{Phone: phone}
	var response = &SendCodeResponse{}

	err := rsl.Soap("", request, "SendCode", response)
	if err != nil {
		return "", err
	}

	return response.SendCodeResult, nil

}
