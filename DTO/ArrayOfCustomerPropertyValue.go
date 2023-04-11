package DTO

type ArrayOfCustomerPropertyValue struct {
	CustomerPropertyValue []*CustomerPropertyValue `xml:"a:CustomerPropertyValue,omitempty"`
}
