package methods

import (
	"encoding/xml"
	"fmt"
	"log"
)

type TestForgetPasswordVerifyCodeRequest struct {
	XMLName  xml.Name `xml:"TestForgetPasswordVerifyCode"`
	SmsToken string   `xml:"smsToken"`
	Code     string   `xml:"code"`
}

type TestForgetPasswordVerifyCodeResponse struct {
	XMLName                            xml.Name `xml:" TestForgetPasswordVerifyCodeResponse"`
	TestForgetPasswordVerifyCodeResult string   `xml:"TestForgetPasswordVerifyCodeResult,omitempty"`
	Token                              string   `xml:"token,omitempty"`
	CheckToken                         string   `xml:"checkToken,omitempty"`
}

func (rsl RSLoyaltyWebService) TestForgetPasswordVerifyCode(channel string, identifier int64, smsToken string, code string) *TestForgetPasswordVerifyCodeResponse {

	//rsl.Init()
	var trace = true
	var name = fmt.Sprintf(lhm, channel, identifier, "Code RestorePassword TestForgetPasswordVerifyCode")
	if trace {
		log.Printf("%s Выполнение проверки кода.", name)
	}

	var request = &TestForgetPasswordVerifyCodeRequest{SmsToken: smsToken, Code: code}
	var response = &TestForgetPasswordVerifyCodeResponse{}
	if trace {
		log.Printf("%s Сформированы объекты ДТО.", name)
	}

	if trace {
		log.Printf("%s Отправляем запрос в лояльность", name)
	}
	status, err := rsl.Soap2(channel, identifier, request, "TestForgetPasswordVerifyCode", response)
	if err != nil {
		log.Printf("%s Ошибка выполнения запроса", name)
		return response
	}

	if status == "200 OK" {
		return response
	} else {
		return response
	}
}
