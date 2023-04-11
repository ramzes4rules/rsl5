package DTO

type DiscountCardDTO struct {
	//Account           AccountDTO             // Счет клиента
	IsDeleted bool // Флаг признака удаления записи
	//DiscountCardGroup []DiscountCardGroupDTO // Коллекция объектов типа DiscountCardGroupDTO
	//Store             StoreDTO               // Магазин в котором была приобретена карта

	DiscountCardID int    // ID дисконтной карты
	Number         string // строка	Номер карты
	IsActive       bool   // Флаг активности карты
	IsBlocked      bool   // Флаг блокировки карты
	ReasonBlocking string // Описание причины блокировки
	ActivationDate string // Дата активации карты
	//ActivationDate time.Time `xml:"ActivationDate,omitempty"` // Дата активации карты
	ExpirationDate string // Дата срока действия карты
	//ExpirationDate 			time.Time 	// Дата срока действия карты
	//DiscountCardGroupIds	[]int 	// Коллекция ID групп карт к которым принадлежит данная карта
	CustomerName string // Имя владельца карты.
	StoreName    string // Наименование магазина
	PinCode      string // PIN код карты (4 значный CVV)
	IsVirtual    bool   `xml:"IsVirtual,omitempty"`
}
