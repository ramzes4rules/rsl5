package DTO

type PropertyRef struct {
	IsRequired   bool        `xml:"a:IsRequired,omitempty"`
	IsUnique     bool        `xml:"a:IsUnique,omitempty"`
	LoadDate     string      `xml:"a:LoadDate,omitempty"`
	Name         string      `xml:"a:Name,omitempty"`
	ObjectType   *ObjectType `xml:"a:ObjectType,omitempty"`
	Position     int32       `xml:"a:Position,omitempty"`
	PropertyID   int64       `xml:"a:PropertyID,omitempty"`
	PropertyType string      `xml:"a:PropertyType,omitempty"`
}
