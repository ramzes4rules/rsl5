package DTO

type PropertyValue struct {
	BooleanValue bool                      `xml:"a:BooleanValue,omitempty"`
	DateValue    string                    `xml:"a:DateValue,omitempty"`
	EnumValue    *EnumPropertyValue        `xml:"a:EnumValue,omitempty"`
	EnumValues   *ArrayOfEnumPropertyValue `xml:"a:EnumValues,omitempty"`
	IntValue     int32                     `xml:"a:IntValue,omitempty"`
	Property     *Property                 `xml:"a:Property,omitempty"`
	StringValue  string                    `xml:"a:StringValue,omitempty"`
}
