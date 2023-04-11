package DTO

type ItemRefDTO struct {
	Articul       string           `xml:"Articul,omitempty"`
	CategoriesIds *ArrayOfint      `xml:"CategoriesIds,omitempty"`
	ItemGroup     *ItemGroupRefDTO `xml:"ItemGroup,omitempty"`
	ItemID        int64            `xml:"ItemID,omitempty"`
	//LoadDate time.Time `xml:"LoadDate,omitempty"`
	LoadDate       string                       `xml:"LoadDate,omitempty"`
	Name           string                       `xml:"Name,omitempty"`
	NoAddBonus     bool                         `xml:"NoAddBonus,omitempty"`
	NoDiscounts    bool                         `xml:"NoDiscounts,omitempty"`
	NoPayBonus     bool                         `xml:"NoPayBonus,omitempty"`
	PropertyValues *ArrayOfKeyValueOfintanyType `xml:"PropertyValues,omitempty"`
}
