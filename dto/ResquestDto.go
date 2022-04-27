package dto

type Request struct {
	FirstNumber  int    `json:"firstNumber"`
	SecondNumber int    `json:"secondNumber"`
	Limit        int    `json:"limit"`
	FirstString  string `json:"firstString"`
	SecondString string `json:"secondString"`
}
