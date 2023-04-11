package methods

import (
	"encoding/xml"
	"fmt"
)

type GetTokenByType struct {
	XMLName   xml.Name `xml:" GetTokenByType"`
	AuthToken string   `xml:"authToken,omitempty"`
	Password  string   `xml:"password,omitempty"`
	//Type_ *PortalAuthType `xml:"type,omitempty"`
	Type_      string `xml:"type,omitempty"`
	IsInternal bool   `xml:"isInternal,omitempty"`
}

type GetTokenByTypeResponse struct {
	XMLName              xml.Name `xml:"GetTokenByTypeResponse"`
	GetTokenByTypeResult string   `xml:"GetTokenByTypeResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetTokenByType(config RSLoyaltyWebService, channel string, identifier int64, login string, password string, loginType string, internal bool) string {

	rsl.Init(config, channel, identifier)

	var name = fmt.Sprintf(nestor.HeaderFormat, channel, identifier, "RSLoyaltyWebService.GetTokenByType")
	nestor.Debug(name, "*** Получение токена авторизации.")

	// создаем экземпляр объекта запроса
	var request = &GetTokenByType{AuthToken: login, Password: password, Type_: loginType, IsInternal: internal}
	var response = &GetTokenByTypeResponse{}
	nestor.Trace(name, "Экземпляры объектов запроса и ответа созданы успешно.")

	// Отправляем запрос в лояльность
	status, err := rsl.Soap2(channel, identifier, request, "GetTokenByType", response)
	nestor.Debug(name, fmt.Sprintf("Статус выполнения запроса: '%s'", status))
	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: '%v'", err))
	}

	//
	switch status {
	case "200 OK":
		{
			nestor.Debug(name, fmt.Sprintf("Получен токен: %s", response.GetTokenByTypeResult))
			return response.GetTokenByTypeResult
		}
	case "503 Service Unavailable":
		{
			nestor.Debug(name, fmt.Sprintf("Лояльность недоступна."))
			return "503"
		}
	case "500 Internal Server Error":
		{
			nestor.Debug(name, fmt.Sprintf("Ошибка авторизации."))
			return "500"
		}
	default:
		//nestor.Debug(name, fmt.Sprintf("Ошибка авторизации."))
		return ""
	}

}
