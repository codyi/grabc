package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Route))
	orm.RegisterModel(new(Permission))
	orm.RegisterModel(new(AssignmentPermission))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(AssignmentRole))
	orm.RegisterModel(new(AssignmentRoute))
}

type BaseModel struct {
	Id       int   `json:"id" label:"Id"`
	CreateAt int32 `json:"create_at" label:"创建时间"`
}

//check current model is new
func (this *BaseModel) IsNewRecord() bool {
	return this.Id <= 0
}
