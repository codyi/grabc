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
	orm.RegisterModel(new(Menu))
}

type IModel interface {
	PrepareDelete() error
}

type BaseModel struct {
	Id       int   `json:"id" label:"Id"`
	CreateAt int32 `json:"create_at" label:"创建时间"`
}

//check current model is new
func (this *BaseModel) IsNewRecord() bool {
	return this.Id <= 0
}

//prepare delete
func (this *BaseModel) PrepareDelete() error {
	return nil
}

//delete data from database
func (this *BaseModel) Delete(m IModel) (int64, error) {
	if err := m.PrepareDelete(); err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	return o.Delete(m)
}
