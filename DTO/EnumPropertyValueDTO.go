package DTO

type EnumPropertyValue struct {
	EnumPropertyValueID int64  `xml:"a:EnumPropertyValueID,omitempty"`
	IsDeleted           bool   `xml:"a:IsDeleted,omitempty"`
	LoadDate            string `xml:"a:LoadDate,omitempty"`
	Name                string `xml:"a:Name,omitempty"`
}
