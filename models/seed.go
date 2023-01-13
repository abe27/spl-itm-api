package models

type SeedOrderGroup struct {
	Data []ObjJson `json:"data"`
}

type ObjJson struct {
	UserName   string `json:"username"`
	AffCode    string `json:"aff"`
	CustCode   string `json:"custcode"`
	CustName   string `json:"custname"`
	GroupOrder string `json:"grp"`
	SubOrder   string `json:"sub_order"`
}
