package methods

import (
	"encoding/xml"
	"fmt"
	"log"
	"unibot/RSLoyatyWebService/DTO"
)

type GetCustomerRegisterDataRequest struct {
	XMLName      xml.Name `xml:"GetCustomerRegisterData"`
	DiscountCard string   `xml:"number"`
}

type GetCustomerRegisterDataResponse struct {
	XMLName                       xml.Name                `xml:"GetCustomerRegisterDataResponse"`
	GetCustomerRegisterDataResult DTO.RegisterCustomerDTO `xml:"GetCustomerRegisterDataResult"`
}

func (rsl RSLoyaltyWebService) GetCustomerRegisterData(channel string, identifier int64, discountCard string) *GetCustomerRegisterDataResponse {

	//rsl.Init()
	var debug = true
	var name = fmt.Sprintf("%s %d %s", channel, identifier, "GetCustomerRegisterData")
	if debug {
		log.Printf("%s Выполнение проверки кода.", name)
	}

	var request = &GetCustomerRegisterDataRequest{DiscountCard: discountCard}
	var response = &GetCustomerRegisterDataResponse{}
	if debug {
		log.Printf("%s Сформированы объекты ДТО.", name)
	}

	if debug {
		log.Printf("%s Отправляем запрос в лояльность", name)
	}
	status, err := rsl.Soap2(channel, identifier, request, "GetCustomerRegisterData", response)
	if err != nil {
		log.Printf("%s Ошибка выполнения запроса", name)
	}

	if status == "200 OK" {
		if debug {
			log.Printf("%s Запрос на получение объекта UserData выполнен успешно.", name)
		}
	}
	return response
}
