package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Menu struct {
	BaseModel
	Name   string `json:"name" label:"菜单名称"`
	Parent int    `json:"parent" label:"父级菜单ID"`
	Route  string `json:"route" label:"菜单地址"`
	Order  int    `json:"order" label:"菜单排序"`
}
