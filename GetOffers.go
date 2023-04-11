package methods

import (
	"encoding/xml"
	"fmt"
	"time"
	"unibot/RSLoyatyWebService/DTO"
)

type GetOffers struct {
	XMLName   xml.Name `xml:"GetOffers"`
	Token     string   `xml:"token,omitempty"`
	CompanyID int64    `xml:"companyID,omitempty"`
	AreaID    int64    `xml:"areaID,omitempty"`
}

type GetOffersResponse struct {
	XMLName         xml.Name           `xml:"GetOffersResponse"`
	GetOffersResult *DTO.ArrayOfstring `xml:"GetOffersResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetOffers(config RSLoyaltyWebService, channel string, identifier int64, token string, companyid int64, areaid int64) []string {

	rsl.Init(config, channel, identifier)

	var name = fmt.Sprintf(nestor.HeaderFormat, channel, identifier, "RSLoyaltyWebService.GetOffers")
	var timer time.Time
	nestor.Debug(name, "*** Процесс получения списка акций в лояльности. \U0001F978")

	timer = time.Now()
	nestor.Trace(name, "Создаем экземпляры объектов запроса и результата.")
	var request = &GetOffers{Token: token, CompanyID: companyid, AreaID: areaid}
	var response = &GetOffersResponse{}
	nestor.Debug(name, fmt.Sprintf("Объекты запроса и результата созданы. Время выполнения: %d мс.",
		time.Since(timer).Milliseconds()))

	timer = time.Now()
	nestor.Trace(name, "Вызываем метод лоляьности.")
	_, err := rsl.Soap2(channel, identifier, request, "GetOffers", response)
	if err != nil {
		nestor.Error(name, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
	}

	nestor.Debug(name, fmt.Sprintf("Ответ получен успешно. Количество строк: %d. Время выполнения: %d мс.",
		len(response.GetOffersResult.String), time.Since(timer).Milliseconds()))
	return response.GetOffersResult.String

}
