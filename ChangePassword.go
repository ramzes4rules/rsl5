package methods

import (
	"encoding/xml"
)

type ChangePassword struct {
	XMLName     xml.Name `xml:"ChangePassword"`
	Token       string   `xml:"token,omitempty"`
	NewPassword string   `xml:"newPassword,omitempty"`
	CheckToken  string   `xml:"checkToken,omitempty"`
	Info        string   `xml:"info,omitempty"`
}

type ChangePasswordResponse struct {
	XMLName              xml.Name `xml:"ChangePasswordResponse"`
	ChangePasswordResult bool     `xml:"ChangePasswordResult,omitempty"`
}

func (rsl RSLoyaltyWebService) ChangePassword(token string, password string, checkToken string, info string) (bool, error) {

	var request = &ChangePassword{Token: token, NewPassword: password, CheckToken: checkToken, Info: info}
	var response = &ChangePasswordResponse{}

	err := rsl.Soap("", request, "ChangePassword", response)
	if err != nil {
		return false, err
	}

	return response.ChangePasswordResult, nil
}
