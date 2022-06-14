package models

type SystemsResponse struct {
	ValidationList Validationlist `json:"validationlist"`
	Status         string         `json:"status"`
	Systems        []Systems      `json:"systems"`
	Valid          bool           `json:"valid"`
}
type Systems struct {
	SystemName        string `json:"systemName"`
	TdpID             string `json:"tdpId"`
	IsEnabled         bool   `json:"isEnabled"`
	IrSupportSource   bool   `json:"irSupportSource"`
	IrSupportTarget   bool   `json:"irSupportTarget"`
	IrSupportOnline   bool   `json:"irSupportOnline"`
	WholeDbcSupport   bool   `json:"wholeDbcSupport"`
	IncludeDbcSupport bool   `json:"includeDbcSupport"`
}
