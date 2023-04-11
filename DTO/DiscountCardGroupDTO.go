package DTO

type DiscountCardGroupDTO struct {
	DiscountCardGroupID int64  //	ID группы дисконтных карт
	Name                string //	Наименование группы карты
	Color               string //	HTML код цвета.
	IsDeleted           bool   //	Флаг признака удаления записи
	CanPayBonuses       bool   //	Флаг признака возможности оплаты бонусами
	CanAccrualBonuses   bool   //	Флаг признака возможности накопления бонусов
	ValidityPeriod      string // Срок действия карты (подарочной)
	IsGiftCard          bool   //	Флаг признака подарочной карты
	OneTimeUse          bool   // 	Флаг признака одноразового использования карты
}
