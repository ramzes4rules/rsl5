package DTO

type ArrayOfKeyValueOfintanyType struct {
	KeyValueOfintanyType []struct {
		Key   int32       `xml:"Key,omitempty"`
		Value interface{} `xml:"Value,omitempty"`
	} `xml:"KeyValueOfintanyType,omitempty"`
}
