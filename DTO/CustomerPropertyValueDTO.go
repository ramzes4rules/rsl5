package DTO

type CustomerPropertyValue struct {
	*PropertyValue
	CustomerPropertyID int64 `xml:"CustomerPropertyID,omitempty"`
}
