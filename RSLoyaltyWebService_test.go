package methods

import (
	"fmt"
	"testing"
)

var rsl = RSLoyaltyWebService{
	Url:        "https://afanasiy-test.retailloyalty.ru/RS.Loyalty.WebClientPortal.Service/RSLoyaltyClientService.svc",
	PrivateKey: "ApiPrivateKey",
	Timeout:    5,
}
var phone = "79132106020"
var password = "qwerty"

//var token = ""

// About
func TestRSLoyaltyWebService_About(t *testing.T) {
	_, err := rsl.About()
	if err != nil {
		t.Errorf("failed to exec About: %v", err)
	}
}

// GetLoyaltySettings
func TestRSLoyaltyWebService_GetLoyaltySettings(t *testing.T) {
	_, err := rsl.GetLoyaltySettings()
	if err != nil {
		t.Errorf("failed to get settings: %v\n", err)
	}
}

// GetCustomerProperties
func TestRSLoyaltyWebService_GetCustomersProperties(t *testing.T) {
	_, err := rsl.GetCustomersProperties()
	if err != nil {
		t.Errorf("failed to get customer properties: %v", err)
	}
}

// DiscountCardIsExist
func TestRSLoyaltyWebService_DiscountCardIsExist(t *testing.T) {
	_, err := rsl.DiscountCardIsExist("01800")
	if err != nil {
		t.Errorf("failed to exec DiscountCardIsExist: %v", err)
	}
}

// DiscountCardIsPersonalized
func TestRSLoyaltyWebService_DiscountCardIsPersonalized(t *testing.T) {
	_, err := rsl.DiscountCardIsPersonalized("01800")
	if err != nil {
		t.Errorf("failed to exec DiscountCardIsExist: %v", err)
	}

	_, err = rsl.DiscountCardIsPersonalized("123456")
	if err != nil {
		t.Errorf("failed to exec DiscountCardIsExist: %v", err)
	}
}

// GetTokenByType run test
func TestRSLoyaltyWebService_GetTokenByType(t *testing.T) {
	var err error
	token, err := rsl.GetTokenByType(phone, password, PortalAuthTypePhone, true)
	if err != nil {
		if err.Error() == "Ошибка авторизации." {
			//t.Fatalf("failed to exec GetTokenBtType: %v", err)
			return
		}
		if err.Error() != "Превышено количество попыток ввода данных для входа." {
			//t.Fatalf("failed to exec GetTokenBtType: %v", err)
			return
		}
		t.Fatalf("failed to exec GetTokenBtType: %v", err)
	}
	fmt.Printf("Token='%s'", token)
}

func TestHasFreeVirtualCards(t *testing.T) {
	has, err := rsl.HasFreeVirtualDiscountCards()
	if err != nil {
		t.Fatalf("failed to exec HasFreeVirtualCards: %v", err)
	}
	fmt.Printf("\tHas free virtual cards: %t\n", has)
}

func TestRSLoyaltyWebService_GetIndicators(t *testing.T) {

	token, err := rsl.GetTokenByType(phone, password, PortalAuthTypePhone, true)
	if err != nil {
		t.Fatalf("\tFailed to get token: %v", err)
		return
	}

	indicators, err := rsl.GetIndicators(token)
	if err != nil {
		t.Fatalf("\tFailed to get token: %v", err)
		return
	}

	fmt.Printf("Got indicators: %d", len(indicators))
}
