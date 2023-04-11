package methods

import (
	"encoding/xml"
	"fmt"
	"log"
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

func (rsl RSLoyaltyWebService) SendSmsForgetPasswordInfo(channel string, identifier int64, tokenPart string, loginType string) (string, string) {
	// void SendSmsForgetPasswordInfo(string tokenPart, PortalAuthType type = PortalAuthType.Login, bool isActivation = false);
	//rsl.Init()

	trace := false
	var name = fmt.Sprintf("[%s/%d] %s", channel, identifier, "SendSmsForgetPasswordInfo")
	if trace {
		log.Printf("%s Отправка проверочного кода для восстановления пароля, tokenPart=%s type=%s", name, tokenPart, loginType)
	}

	// Шаг 1. Создаем тело soap запроcа
	var request = &SendSmsForgetPasswordInfoRequest{TokenPart: tokenPart, Type: loginType}
	var response = &SendSmsForgetPasswordInfoResponse{}

	// Шаг 2. Вызываем метод лояльности
	status, err := rsl.Soap2(channel, identifier, request, "SendSmsForgetPasswordInfo", response)
	if err != nil {
		log.Printf("%s Ошибка выполнения запроса, завершаем работу.\n", name)
		return "Error", ""
	}

	if status == "200 OK" {
		if trace {
			log.Printf("%s Запрос выполнен успешно SmsToken=%s, завершаем работу.", name, response.SmsToken)
		}
		return status, response.SmsToken
	} else {
		if trace {
			log.Printf("%s Ошибка отправки кода status=%s, завершаем работу.", name, status)
		}
		return "Error", ""
	}
}
