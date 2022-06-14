package models

type ClientValidationList struct {
	ValStatus string `json:"valStatus"`
	Message   string `json:"message"`
}
type ServerValidationList struct {
	ValStatus string `json:"valStatus"`
	Message   string `json:"message"`
}
type Validationlist struct {
	ClientValidationList []ClientValidationList `json:"clientValidationList"`
	ServerValidationList []ServerValidationList `json:"serverValidationList"`
	Links                []interface{}          `json:"links"`
}
