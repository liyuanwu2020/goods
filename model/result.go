package model

type Result struct {
	Code int
	Msg  string
	Data any
}

type Goods struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
