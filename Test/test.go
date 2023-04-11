package main

import (
	"fmt"
	RSL "github.com/ramzes4rules/rsl5"
	"github.com/ramzes4rules/rsl5/DTO"
	"strconv"
)

func main() {
	//
	fmt.Printf("testing rsl5 methods...\n")

	var rsl = RSL.RSLoyaltyWebService{
		Url:        "https://10.14.5.114/RS.Loyalty.WebClientPortal.Service/RSLoyaltyClientService.svc",
		PrivateKey: "API_PrivateKey",
		Timeout:    10,
	}

	// test About
	about, err := rsl.About()
	if err == nil {
		fmt.Printf("About numbers: %d\n", len(about))
	} else {
		fmt.Printf("Failed to exec About: %v\n", err)
	}

	settings, err := rsl.GetLoyaltySettings()
	if err == nil {
		fmt.Printf("Settings numbers: %d\n", len(settings))
	} else {
		fmt.Printf("Failed to exec GetLoyaltySettings: %v\n", err)
	}

	//
	properties, err := rsl.GetCustomersProperties()
	if err == nil {
		fmt.Printf("Properties numbers: %d\n", len(properties))
	} else {
		fmt.Printf("Failed to exec GetCustomersProperties: %v\n", err)
	}

	//
	cards, err := rsl.GetPhoneDiscountCards("+79132106020")
	if err == nil {
		fmt.Printf("Cards numbers: %d\n", len(cards))
	} else {
		fmt.Printf("Failed to exec GetCustomersProperties: %v\n", err)
	}

	// run test HasFreeVirtualDiscountCardsRequest
	has, err := rsl.HasFreeVirtualDiscountCards()
	if err == nil {
		fmt.Printf("Has free virtual discount cards: %t\n", has)
	} else {
		fmt.Printf("Failed to exec HasFreeVirtualDiscountCards: %v\n", err)
	}

	// run test HasFreeVirtualDiscountCardsRequest
	userData := DTO.RegisterCustomerDTO{
		Xmlna:              "http://schemas.datacontract.org/2004/07/RS.Loyalty.WebClientPortal.Core.Model",
		Xmlni:              "http://www.w3.org/2001/XMLSchema-instance",
		CanRegister:        true,
		CreateVirtual:      true,
		DiscountCardNumber: "",
		FirstName:          "Semen",
		HasCards:           false,
		LastName:           "Gorbunkov",
		Login:              "",
		Password:           "qwerty",
		PinCode:            "",
		Properties:         nil,
		SecondName:         "",
		TermsOfUsing:       nil,
	}
	var terms = DTO.ArrayOfKeyValueOflongboolean{Xmlnb: "http://schemas.microsoft.com/2003/10/Serialization/Arrays"}
	for _, company := range about {
		terms.KeyValueOflongboolean = append(terms.KeyValueOflongboolean, DTO.KeyValueOflongboolean{Key: int64(company.CompanyID), Value: true})
	}
	userData.TermsOfUsing = &terms
	var phonePropertyID int64
	for _, setting := range settings {
		if setting.KeyValue == "OptionPhoneCustomerProperty" {
			phonePropertyID, _ = strconv.ParseInt(setting.Value, 10, 64)
			break
		}
	}
	var phone = DTO.CustomerPropertyValue{
		PropertyValue: &DTO.PropertyValue{
			StringValue: "+79132106029",
			Property: &DTO.Property{
				PropertyRef: DTO.PropertyRef{
					PropertyID:   phonePropertyID,
					PropertyType: "String"},
			},
		},
	}
	userData.Properties = &DTO.ArrayOfCustomerPropertyValue{}
	userData.Properties.CustomerPropertyValue = []*DTO.CustomerPropertyValue{}
	userData.Properties.CustomerPropertyValue = append(userData.Properties.CustomerPropertyValue, &phone)

	card, info, err := rsl.ActivateClientService(userData, "info")
	if err == nil {
		fmt.Printf("card: %s, info: %s\n", card, info)
	} else {
		fmt.Printf("Failed to exec ActivateClientService: %v\n", err)
	}
}
