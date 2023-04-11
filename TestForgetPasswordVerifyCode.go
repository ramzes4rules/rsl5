package methods

import (
	"encoding/xml"
)

type TestForgetPasswordVerifyCodeRequest struct {
	XMLName  xml.Name `xml:"TestForgetPasswordVerifyCode"`
	SmsToken string   `xml:"smsToken"`
	Code     string   `xml:"code"`
}

type TestForgetPasswordVerifyCodeResponse struct {
	XMLName                            xml.Name `xml:"TestForgetPasswordVerifyCodeResponse"`
	TestForgetPasswordVerifyCodeResult string   `xml:"TestForgetPasswordVerifyCodeResult,omitempty"`
	Token                              string   `xml:"token,omitempty"`
	CheckToken                         string   `xml:"checkToken,omitempty"`
}

func (rsl RSLoyaltyWebService) TestForgetPasswordVerifyCode(smsToken string, code string) (*TestForgetPasswordVerifyCodeResponse, error) {

	var request = &TestForgetPasswordVerifyCodeRequest{SmsToken: smsToken, Code: code}
	var response = &TestForgetPasswordVerifyCodeResponse{}
	err := rsl.Soap("", request, "TestForgetPasswordVerifyCode", response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
