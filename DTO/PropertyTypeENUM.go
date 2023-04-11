package DTO

type PropertyType string

const (
	PropertyTypeInteger      PropertyType = "Integer"
	PropertyTypeString       PropertyType = "String"
	PropertyTypeDate         PropertyType = "Date"
	PropertyTypeEnum         PropertyType = "Enum"
	PropertyTypeBoolean      PropertyType = "Boolean"
	PropertyTypeEnumMultiple PropertyType = "EnumMultiple"
)
