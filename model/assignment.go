package model

type Assignment struct {
	User_id   int    `json:"user_id" label:"用户ID"`
	Item_name string `jsob:"item_name" label:"权限名称"`
	Create_at int    `json:"create_at" label:"创建时间"`
}
