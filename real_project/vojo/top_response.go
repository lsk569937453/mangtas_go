package vojo

type TopResponse struct {
	BaseRes
	ResMessage []TopServiceKV `json:"resMessage"`
}
