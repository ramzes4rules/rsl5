package DTO

import "encoding/xml"

type Property struct {
	XMLName xml.Name `xml:"a:Property,omitempty"`
	PropertyRef
	IsDeleted    bool                      `xml:"a:IsDeleted,omitempty"`
	Pattern      string                    `xml:"a:Pattern,omitempty"`
	PresetType   string                    `xml:"a:PresetType,omitempty"`
	UseInFilter  bool                      `xml:"a:UseInFilter,omitempty"`
	UserEditType *UserEditType             `xml:"a:UserEditType,omitempty"`
	Values       *ArrayOfEnumPropertyValue `xml:"a:Values,omitempty"`
}
