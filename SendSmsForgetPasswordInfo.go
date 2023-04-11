package methods

import (
	"encoding/xml"
)

type SendSmsForgetPasswordInfoRequest struct {
	XMLName   xml.Name `xml:"SendSmsForgetPasswordInfo"`
	TokenPart string   `xml:"tokenPart"`
	Type      string   `xml:"type"`
}

type SendSmsForgetPasswordInfoResponse struct {
	XMLName  xml.Name `xml:"SendSmsForgetPasswordInfoResponse"`
	Result   string   `xml:"SendSmsForgetPasswordInfoResult"`
	SmsToken string   `xml:"smsToken"`
}

func (rsl RSLoyaltyWebService) SendSmsForgetPasswordInfo(tokenPart string, loginType string) (string, error) {

	var request = &SendSmsForgetPasswordInfoRequest{TokenPart: tokenPart, Type: loginType}
	var response = &SendSmsForgetPasswordInfoResponse{}

	err := rsl.Soap("", request, "SendSmsForgetPasswordInfo", response)
	if err != nil {
		return "", err
	}

	return response.SmsToken, nil
}
