package methods

type RSLoyaltyWebService struct {
	Url        string `json:"url"`      // Url Loyalty end point
	PrivateKey string `json:"api_key "` // PrivateKey Api key
	Timeout    uint8  `json:"timeout"`  // Timeout operation
}
