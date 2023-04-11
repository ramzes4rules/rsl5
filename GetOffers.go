package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetOffers struct {
	XMLName   xml.Name `xml:"GetOffers"`
	Token     string   `xml:"token,omitempty"`
	CompanyID int64    `xml:"companyID,omitempty"`
	AreaID    int64    `xml:"areaID,omitempty"`
}

type GetOffersResponse struct {
	XMLName         xml.Name           `xml:"GetOffersResponse"`
	GetOffersResult *DTO.ArrayOfstring `xml:"GetOffersResult,omitempty"`
}

func (rsl RSLoyaltyWebService) GetOffers(token string, companyid int64, areaid int64) ([]string, error) {

	var request = &GetOffers{Token: token, CompanyID: companyid, AreaID: areaid}
	var response = &GetOffersResponse{}

	err := rsl.Soap("", request, "GetOffers", response)
	if err != nil {
		return nil, err
	}

	return response.GetOffersResult.String, nil

}
