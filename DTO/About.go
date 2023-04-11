package DTO

import "encoding/xml"

type About struct {
	XMLName      xml.Name `xml:"AboutDTO"`
	CompanyAbout string
	CompanyID    int
	CompanyName  string
}
