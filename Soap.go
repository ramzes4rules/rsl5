package methods

import (
	"crypto/tls"
	"encoding/xml"
	"github.com/pkg/errors"
	"io"

	"log"
	"net/http"
	"strings"
	"time"
	//"unibot/RSLoyatyWebService/DTO"
)

const header = `<?xml version="1.0" encoding="utf-8"?>`

type Fault struct {
	Faultcode   string `xml:"faultcode,omitempty"`
	Faultstring string `xml:"faultstring,omitempty"`
}

type RequestEnvelope struct {
	XMLName xml.Name    `xml:"soap:Envelope"`
	Xmlns   string      `xml:"xmlns:soap,attr"`
	Body    RequestBody `xml:"soap:Body"`
}
type RequestBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Request any
}

type ResponseEnvelope struct {
	XMLName xml.Name     `xml:"Envelope"`
	Xmlns   string       `xml:"xmlns:soap,attr"`
	Body    ResponseBody `xml:"s:Body"`
}
type ResponseBody struct {
	XMLName  xml.Name `xml:"s:Body"`
	Response any      //`xml:"AboutResponse"`
}

func (rsl RSLoyaltyWebService) Soap(identifier string, request any, method string, response any) error {

	//var service = fmt.Sprintf("%s [%s]", "SoapRequest", identifier)
	var fault Fault
	//var mara = rsl.Mara
	//mara.Trace(service, "*** Выполнение soap-запроса.")

	//mara.Trace(service, "Отключаем проверку ssl-сертификата.")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//mara.Trace(service, "Проверка ssl-сертификата отключена.")

	//mara.Trace(service, "Формируем soap-запрос.")
	var Envelope = RequestEnvelope{Xmlns: "http://schemas.xmlsoap.org/soap/envelope/"}
	Envelope.Body.Request = request
	req1, err := xml.Marshal(Envelope)
	if err != nil {
		//mara.Error(service, fmt.Sprintf("%v", err))
	}
	req2 := header + string(req1)
	//fmt.Println(req2)
	//mara.Debug(service, fmt.Sprintf("Запрос для отправки сформирован: %s", req2))

	//mara.Trace(service, "Создаем http-запрос.")
	payload := strings.NewReader(req2)
	req, err := http.NewRequest("POST", rsl.Url, payload)
	if err != nil {
		//mara.Error(service, fmt.Sprintf("Ошибка создания запроса: %v", err))
		return err
	}
	req.Header.Add("soapAction", "urn:IRSLoyaltyClientService/"+method)
	req.Header.Add("Content-Type", "text/xml")

	//
	client := &http.Client{}
	client.Timeout = time.Duration(rsl.Timeout) * time.Second
	res, err := client.Do(req)
	if err != nil {
		//mara.Error(service, fmt.Sprintf("Ошибка выполнения запроса: %v", err))
		return err
	}
	//mara.Trace(service, fmt.Sprintf("Запрос выполнен, статус: %s", res.Status))

	// Читаем ответ
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return err
	}
	//mara.Debug(service, fmt.Sprintf("Ответ лояльности: %s", string(body)))

	var r = strings.Replace(string(body),
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\"><s:Body>", "", -1)
	r = strings.Replace(r, "</s:Body></s:Envelope>", "", -1)
	//mara.Trace(service, fmt.Sprintf("Получен результат: '%s'", r))
	//fmt.Println(service, fmt.Sprintf("Получен результат: '%s'", r))

	// Парсим ответ
	switch res.Status {
	case "500 Internal Server Error":
		err = xml.Unmarshal([]byte(r), &fault)
		if err != nil {
			//mara.Error(service, fmt.Sprintf("Ошибка парсинга объекта fault: '%v'", err))
			return err
		}
		//mara.Error(service, fmt.Sprintf("Получена ошибка выполнения запроса: '%s'", fault.Faultstring))

		return errors.New(fault.Faultstring)
	default:
		err = xml.Unmarshal([]byte(r), &response)
		if err != nil {
			//mara.Error(service, fmt.Sprintf("Ошибка парсинга: %v", err))
			return err
		}
		//mara.Trace(service, fmt.Sprintf("Разобран ответ: %s", response))
		return nil
	}

}
