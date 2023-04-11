package methods

import (
	"encoding/xml"
	"fmt"
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

func (rsl RSLoyaltyWebService) TestCode(channel string, identifier int64, phone string, code string) string {

	var name = fmt.Sprintf("%s %d %s", channel, identifier, "TestCode")
	nestor.Debug(name, "Проверка кода подтверждения.")

	var request = &TestCodeRequest{Phone: phone, Code: code}
	var response = &TestCodeResponse{}

	_, err := rsl.Soap2(channel, identifier, request, "TestCode", response)
	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка: %v", err))
		return fmt.Sprintf("%v", err)
	}

	nestor.Debug(name, fmt.Sprintf("Результат проверки: '%s'", response.TestCodeResult))
	return response.TestCodeResult
}
