package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Permission struct {
	BaseModel
	Name        string `json:"name" label:"权限名称"`
	Description string `json:"description" label:"权限描述"`
}

func (this *Permission) TableName() string {
	return "rabc_permission"
}

//Find one permission by id from database
func (this *Permission) FindById(id int) error {
	if id <= 0 {
		return errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

//Find one permission by name from database
func (this *Permission) FindByName(name string) error {
	if name == "" {
		return errors.New("权限名称不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("name", name).One(this)
}

//insert current permission to database
//not insert if name is exist
func (this *Permission) Insert() (isInsert bool, err error) {
	if this.Name == "" {
		return false, errors.New("权限名称不能为空")
	}

	self := Permission{}
	self.FindByName(this.Name)
	if self.Id > 0 {
		return false, errors.New("权限已经存在")
	}

	this.CreateAt = int32(time.Now().Unix())
	o := orm.NewOrm()

	id, err := o.Insert(this)

	return id > 0, err
}

//update current permission to database
//not update if name is exist
func (this *Permission) Update() (err error) {
	if this.Name == "" {
		return errors.New("权限名称不能为空")
	}

	self := Permission{}
	self.FindByName(this.Name)
	if self.Id > 0 && self.Id != this.Id {
		return errors.New("权限已经存在")
	}

	o := orm.NewOrm()

	_, err = o.Update(this)

	return err
}

//list permissions
func (this Permission) List(pageIndex, pageCount int) ([]*Permission, int, error) {
	var permissions []*Permission
	var total int64
	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Limit(pageCount).Offset(pageCount * (pageIndex - 1)).All(&permissions)

	if err != nil {
		return permissions, int(total), err
	}

	total, err = o.QueryTable(this.TableName()).Count()
	return permissions, int(total), err
}

//retrieve all permissions
func (this Permission) FindAll() ([]*Permission, error) {
	var permissions []*Permission
	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).All(&permissions)

	if err != nil {
		return permissions, err
	}

	return permissions, err
}

func (this *Permission) PrepareDelete() error {
	if this.IsNewRecord() {
		return errors.New("当前对象为空")
	}

	routeAssignmentModel := AssignmentRoute{}
	err := routeAssignmentModel.DeleteByPermissionId(this.Id)

	if err != nil {
		return err
	}

	assignmentPermission := AssignmentPermission{}
	return assignmentPermission.DeleteByPermissionId(this.Id)
}

//remove current name from database
func (this *Permission) Delete() (bool, error) {
	num, err := this.BaseModel.Delete(this)
	return num > 0, err
}
