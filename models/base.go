package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Route))
	orm.RegisterModel(new(Permission))
}

type BaseModel struct {
}
