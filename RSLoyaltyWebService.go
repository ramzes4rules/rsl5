package methods

import (
	"fmt"
	"unibot/nestor2"
)

const lhm = "[%s/%d] %s"

type RSLoyaltyWebService struct {
	Url        string          `json:"Точка подключения"`     //
	PrivateKey string          `json:"ключ"`                  //
	Nestor     nestor2.Nestor2 `json:"Параметры логирования"` //
	//FileLog    string
	LogLevel nestor2.LogLevel
	LogFile  string
	PrintLog bool
	WriteLog bool
}

var nestor nestor2.Nestor2

func (rsl *RSLoyaltyWebService) Init(config RSLoyaltyWebService, channel string, identifier int64) {

	rsl.Url = config.Url
	rsl.PrivateKey = config.PrivateKey

	nestor = nestor2.Nestor2{
		Level:        config.Nestor.Level,
		File:         config.Nestor.File,
		Print:        config.Nestor.Print,
		Write:        config.Nestor.Write,
		HeaderFormat: config.Nestor.HeaderFormat,
		DateFormat:   config.Nestor.DateFormat,
	}

	var name = fmt.Sprintf(nestor.HeaderFormat, channel, identifier, "RSLoyaltyWebService.Init")
	nestor.Trace(name, "*** Иницализация экземпляра объекта RSLoyaltyWebService")

}
