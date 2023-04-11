package DTO

type UserEditType string

const (
	UserEditTypePrivate  UserEditType = "Private"
	UserEditTypeCanEdit  UserEditType = "CanEdit"
	UserEditTypeReadOnly UserEditType = "ReadOnly"
	UserEditTypeEditOne  UserEditType = "EditOne"
)
