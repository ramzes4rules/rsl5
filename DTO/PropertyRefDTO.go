package DTO

type PropertyRef struct {
	IsRequired   bool        `xml:"IsRequired,omitempty"`
	IsUnique     bool        `xml:"IsUnique,omitempty"`
	LoadDate     string      `xml:"LoadDate,omitempty"`
	Name         string      `xml:"Name,omitempty"`
	ObjectType   *ObjectType `xml:"ObjectType,omitempty"`
	Position     int32       `xml:"Position,omitempty"`
	PropertyID   int64       `xml:"PropertyID,omitempty"`
	PropertyType string      `xml:"PropertyType,omitempty"`
}
