package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Permission struct {
	BaseModel
	Id          int    `json:"id" label:"Id"`
	Name        string `json:"name" label:"权限名称"`
	Description string `json:"description" label:"权限描述"`
	CreateAt    int32  `json:"create_at" label:"创建时间"`
}

func (this *Permission) TableName() string {
	return "rabc_permission"
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

//remove current name from database
func (this Permission) DeleteByName(name string) (isDelete bool, err error) {
	if name == "" {
		return false, errors.New("权限名称不能为空")
	}
	o := orm.NewOrm()

	permissionModel := &Permission{}
	o.QueryTable(this.TableName()).Filter("name", name).One(permissionModel)

	num, err := o.Delete(permissionModel)

	return num > 0, err
}

//retrieve all permissions
func (this Permission) FindAll(pageIndex, pageCount int) ([]*Permission, int, error) {
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
