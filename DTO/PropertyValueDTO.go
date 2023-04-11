package DTO

type PropertyValue struct {
	BooleanValue bool                      `xml:"BooleanValue,omitempty"`
	DateValue    string                    `xml:"DateValue,omitempty"`
	EnumValue    *EnumPropertyValue        `xml:"EnumValue,omitempty"`
	EnumValues   *ArrayOfEnumPropertyValue `xml:"EnumValues,omitempty"`
	IntValue     int32                     `xml:"IntValue,omitempty"`
	Property     *Property                 `xml:"Property,omitempty"`
	StringValue  string                    `xml:"StringValue,omitempty"`
}
