package DTO

type Customer struct {
	//Xmlna                             string                        `xml:"xmlns:a,attr"`
	//Xmlni                             string                        `xml:"xmlns:i,attr"`
	Categories                        *ArrayOfCategory              `xml:"Categories,omitempty"`
	IsAddedBonusForMobileRegistration bool                          `xml:"IsAddedBonusForMobileRegistration,omitempty"`
	IsAddedBonusForRegistration       bool                          `xml:"IsAddedBonusForRegistration,omitempty"`
	IsDeleted                         bool                          `xml:"IsDeleted,omitempty"`
	IsFillAllProperties               bool                          `xml:"IsFillAllProperties,omitempty"`
	PhoneIsChecked                    bool                          `xml:"PhoneIsChecked,omitempty"`
	Properties                        *ArrayOfCustomerPropertyValue `xml:"Properties,omitempty"`
	*CustomerRef
}
