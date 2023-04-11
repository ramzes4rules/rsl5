package methods

import (
	"encoding/xml"
	"fmt"
	"unibot/RSLoyatyWebService/DTO"
)

type ActivateClientService struct {
	XMLName  xml.Name                `xml:"ActivateClientService"`
	UserData DTO.RegisterCustomerDTO `xml:"userData,omitempty"`
	Info     string                  `xml:"info,omitempty"`
}

type ActivateClientServiceResponse struct {
	XMLName                     xml.Name `xml:"ActivateClientServiceResponse"`
	ActivateClientServiceResult string   `xml:"ActivateClientServiceResult,omitempty"`
	Info                        string   `xml:"info,omitempty"`
}

func (rsl RSLoyaltyWebService) ActivateClientService(channel string, identifier int64, userData DTO.RegisterCustomerDTO, info string) (string, string, string) {

	var name = fmt.Sprintf("[%s/%d] %s", channel, identifier, "RSLoyaltyWebService.ActivateClientService")
	nestor.Debug(name, fmt.Sprintf("Активация пользователя '%s %s'.", userData.FirstName, userData.LastName))

	var request = &ActivateClientService{UserData: userData, Info: info}
	var response = &ActivateClientServiceResponse{}
	nestor.Debug(name, "Сформированы объекты DTO.")
	//nestor.Debug(name, fmt.Sprintf("CreateVirtual: %t", request.UserData.CreateVirtual))

	nestor.Debug(name, "Отправляем запрос в лояльность")
	status, err := rsl.Soap2(channel, identifier, request, "ActivateClientService", response)
	if err != nil {
		nestor.Error(name, "Ошибка выполнения запроса, завершаем работу")
		return "Error", "", ""
	}

	if status == "200 OK" {
		nestor.Debug(name, fmt.Sprintf("Клиент активирован. Привязана карта %s, получено сообщение для клиента '%s'",
			response.ActivateClientServiceResult, response.Info))
		return status, response.ActivateClientServiceResult, response.Info
	} else {
		nestor.Error(name, fmt.Sprintf("Ошибка активации клиента, завершаем работу"))
		return status, "", ""
	}
}
