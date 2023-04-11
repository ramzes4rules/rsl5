package DTO

type CategoryDTO struct {
	CategoryID int64  `xml:"CategoryID,omitempty"`
	Color      string `xml:"Color,omitempty"`
	IsDeleted  bool   `xml:"IsDeleted,omitempty"`
	Name       string `xml:"Name,omitempty"`
}
