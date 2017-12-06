package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Role struct {
	BaseModel
	Name        string `json:"name" label:"角色名称"`
	Description string `json:"description" label:"角色描述"`
}

func (this *Role) TableName() string {
	return "rabc_role"
}

//Find one Role by id from database
func (this *Role) FindById(id int) error {
	if id <= 0 {
		return errors.New("角色ID不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

//Find one Role by name from database
func (this *Role) FindByName(name string) error {
	if name == "" {
		return errors.New("角色名称不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("name", name).One(this)
}

//insert current Role to database
//not insert if name is exist
func (this *Role) Insert() (isInsert bool, err error) {
	if this.Name == "" {
		return false, errors.New("角色名称不能为空")
	}

	self := Role{}
	self.FindByName(this.Name)
	if self.Id > 0 {
		return false, errors.New("角色已经存在")
	}

	this.CreateAt = int32(time.Now().Unix())
	o := orm.NewOrm()

	id, err := o.Insert(this)

	return id > 0, err
}

//update current Role to database
//not update if name is exist
func (this *Role) Update() (err error) {
	if this.Name == "" {
		return errors.New("角色名称不能为空")
	}

	self := Role{}
	self.FindByName(this.Name)
	if self.Id > 0 && self.Id != this.Id {
		return errors.New("角色已经存在")
	}

	o := orm.NewOrm()

	_, err = o.Update(this)

	return err
}

//remove current name from database
func (this Role) DeleteByName(name string) (isDelete bool, err error) {
	if name == "" {
		return false, errors.New("角色名称不能为空")
	}
	o := orm.NewOrm()

	RoleModel := &Role{}
	o.QueryTable(this.TableName()).Filter("name", name).One(RoleModel)

	num, err := o.Delete(RoleModel)

	return num > 0, err
}

//retrieve all Roles
func (this Role) FindAll(pageIndex, pageCount int) ([]*Role, int, error) {
	var Roles []*Role
	var total int64
	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Limit(pageCount).Offset(pageCount * (pageIndex - 1)).All(&Roles)

	if err != nil {
		return Roles, int(total), err
	}

	total, err = o.QueryTable(this.TableName()).Count()
	return Roles, int(total), err
}

//remove current name from database
func (this *Role) Delete() (isDelete bool, err error) {
	if this.Id <= 0 {
		return false, errors.New("数据不能为空")
	}

	o := orm.NewOrm()
	num, err := o.Delete(this)

	return num > 0, err
}
