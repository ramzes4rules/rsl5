package DTO

type PropertyGet struct {
	*PropertyRef
	IsDeleted    bool                         `xml:"IsDeleted,omitempty"`
	Pattern      string                       `xml:"Pattern,omitempty"`
	PresetType   string                       `xml:"PresetType,omitempty"`
	UseInFilter  bool                         `xml:"UseInFilter,omitempty"`
	UserEditType *UserEditType                `xml:"UserEditType,omitempty"`
	Values       *ArrayOfEnumPropertyValueGet `xml:"Values,omitempty"`
}
