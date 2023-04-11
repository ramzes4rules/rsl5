package methods

import (
	"encoding/xml"
	"github.com/ramzes4rules/rsl5/DTO"
)

type GetAllOperations struct {
	XMLName xml.Name `xml:"GetAllOperations"`
	Token   string   `xml:"token,omitempty"`
}
type GetAllOperationsResponse struct {
	XMLName                xml.Name       `xml:"GetAllOperationsResponse"`
	GetAllOperationsResult *ArrayOfCheque `xml:"GetAllOperationsResult,omitempty"`
}
type ArrayOfCheque struct {
	Cheque []*DTO.ChequeDTO `xml:"Cheque,omitempty"`
}

func (rsl RSLoyaltyWebService) GetAllOperations(token string) ([]*DTO.ChequeDTO, error) {

	var request = &GetAllOperations{Token: token}
	var response = &GetAllOperationsResponse{}

	err := rsl.Soap("", request, "GetAllOperations", response)
	if err != nil {
		return []*DTO.ChequeDTO{}, err
	}

	return response.GetAllOperationsResult.Cheque, nil
}
