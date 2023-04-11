package DTO

type AccountDTO struct {
	AccountID int64    //ID счета
	Customer  Customer //Ссылка на объект клиента.
}
