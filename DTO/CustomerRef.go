package DTO

type CustomerRef struct {
	//XMLName              xml.Name    `xml:"CustomerRef"`
	AdministrativeAreaID int64       `xml:"AdministrativeAreaID,omitempty"`
	CategoriesIds        *ArrayOfint `xml:"CategoriesIds,omitempty"`
	CustomerID           int64       `xml:"CustomerID,omitempty"`
	FirstName            string      `xml:"FirstName,omitempty"`
	FullName             string      `xml:"FullName,omitempty"`
	IsSmsSubscribed      bool        `xml:"IsSmsSubscribed,omitempty"`
	IsSubscribed         bool        `xml:"IsSubscribed,omitempty"`
	LastName             string      `xml:"LastName,omitempty"`
	LocalityID           int64       `xml:"LocalityID,omitempty"`
	SecondName           string      `xml:"SecondName,omitempty"`
}
