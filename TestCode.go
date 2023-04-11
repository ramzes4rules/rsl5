package methods

import (
	"encoding/xml"
)

type TestCodeRequest struct {
	XMLName xml.Name `xml:"TestCode"`
	Phone   string   `xml:"phone"`
	Code    string   `xml:"code"`
}

type TestCodeResponse struct {
	XMLName        xml.Name `xml:"TestCodeResponse"`
	TestCodeResult string   `xml:"TestCodeResult"`
}

func (rsl RSLoyaltyWebService) TestCode(phone string, code string) (string, error) {
	var request = &TestCodeRequest{Phone: phone, Code: code}
	var response = &TestCodeResponse{}
	err := rsl.Soap("", request, "TestCode", response)
	if err != nil {
		return "", err
	}
	return response.TestCodeResult, nil
}
