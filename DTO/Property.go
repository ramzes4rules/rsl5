package DTO

import "encoding/xml"

type Property struct {
	XMLName xml.Name `xml:"Property,omitempty"`
	PropertyRef
	IsDeleted    bool                      `xml:"IsDeleted,omitempty"`
	Pattern      string                    `xml:"Pattern,omitempty"`
	PresetType   string                    `xml:"PresetType,omitempty"`
	UseInFilter  bool                      `xml:"UseInFilter,omitempty"`
	UserEditType *UserEditType             `xml:"UserEditType,omitempty"`
	Values       *ArrayOfEnumPropertyValue `xml:"Values,omitempty"`
}
