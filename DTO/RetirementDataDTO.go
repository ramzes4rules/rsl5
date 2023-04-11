package DTO

type RetirementDataDTO struct {
	Bonus   float64     `xml:"Bonus,omitempty"`
	Company *CompanyDTO `xml:"Company,omitempty"`
	//RetirementDate time.Time `xml:"RetirementDate,omitempty"`
	RetirementDate string `xml:"RetirementDate,omitempty"`
	RetirementText string `xml:"RetirementText,omitempty"`
}
