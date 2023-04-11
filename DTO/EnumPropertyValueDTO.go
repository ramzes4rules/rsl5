package DTO

type EnumPropertyValue struct {
	EnumPropertyValueID int64  `xml:"EnumPropertyValueID,omitempty"`
	IsDeleted           bool   `xml:"IsDeleted,omitempty"`
	LoadDate            string `xml:"LoadDate,omitempty"`
	Name                string `xml:"Name,omitempty"`
}
