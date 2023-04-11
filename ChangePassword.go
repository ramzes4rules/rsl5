package methods

import (
	"encoding/xml"
	"fmt"
	"log"
)

type ChangePassword struct {
	XMLName     xml.Name `xml:"ChangePassword"`
	Token       string   `xml:"token,omitempty"`
	NewPassword string   `xml:"newPassword,omitempty"`
	CheckToken  string   `xml:"checkToken,omitempty"`
	Info        string   `xml:"info,omitempty"`
}

type ChangePasswordResponse struct {
	XMLName              xml.Name `xml:"ChangePasswordResponse"`
	ChangePasswordResult bool     `xml:"ChangePasswordResult,omitempty"`
}

func (rsl RSLoyaltyWebService) ChangePassword(channel string, identifier int64, token string, password string, checkToken string, info string) bool {

	//rsl.Init()
	var trace = false
	var name = fmt.Sprintf(lhm, channel, identifier, "ChangePassword")
	if trace {
		log.Printf("%s Изменение пароля.\n", name)
	}

	var request = &ChangePassword{Token: token, NewPassword: password, CheckToken: checkToken, Info: info}
	var response = &ChangePasswordResponse{}
	if trace {
		log.Printf("%s Сформированы объекты ДТО.", name)
	}

	if trace {
		log.Printf("%s Отправляем запрос в лояльность", name)
	}
	status, err := rsl.Soap2(channel, identifier, request, "ChangePassword", response)
	if err != nil {
		log.Printf("%s Ошибка выполнения запроса", name)
		return false
	}

	if status == "200 OK" {
		return response.ChangePasswordResult
	} else {
		return response.ChangePasswordResult
	}

}
