package methods

import (
	"encoding/xml"
)

type GetTokenByType struct {
	XMLName    xml.Name       `xml:" GetTokenByType"`
	AuthToken  string         `xml:"authToken,omitempty"`
	Password   string         `xml:"password,omitempty"`
	Type       PortalAuthType `xml:"type,omitempty"`
	IsInternal bool           `xml:"isInternal,omitempty"`
}

type GetTokenByTypeResponse struct {
	XMLName              xml.Name `xml:"GetTokenByTypeResponse"`
	GetTokenByTypeResult string   `xml:"GetTokenByTypeResult,omitempty"`
}

type PortalAuthType string

const (
	PortalAuthTypePhone PortalAuthType = "Phone"
	PortalAuthTypeEmail PortalAuthType = "Email"
	PortalAuthTypeCard  PortalAuthType = "Card"
	PortalAuthTypeLogin PortalAuthType = "Login"
)

func (rsl RSLoyaltyWebService) GetTokenByType(login string, password string, loginType PortalAuthType, internal bool) (string, error) {

	var request = &GetTokenByType{AuthToken: login, Password: password, Type: loginType, IsInternal: internal}
	var response = &GetTokenByTypeResponse{}

	err := rsl.Soap("", request, "GetTokenByType", response)
	if err != nil {
		return "", err
	}

	return response.GetTokenByTypeResult, nil

}
