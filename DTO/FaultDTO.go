package DTO

//<s:Fault>
//	<faultcode>s:Client</faultcode>
//	<faultstring xml:lang="ru-RU">Указан неверный код карты.</faultstring>
//	<detail>
//		<string xmlns="http://schemas.microsoft.com/2003/10/Serialization/">
//			Указан неверный код карты.
//		</string>
//	</detail>
//</s:Fault>'

type Fault struct {
	Faultcode   string `xml:"faultcode,omitempty"`
	Faultstring string `xml:"faultstring,omitempty"`
}
