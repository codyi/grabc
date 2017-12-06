package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type RoleAssignment struct {
	BaseModel
	UserId int `json:"user_id" label:"用户ID"`
	RoleId int `json:"role_id" label:"角色ID"`
}

func (this *RoleAssignment) TableName() string {
	return "rabc_role_assignment"
}

//insert current data to database
//not insert if name is exist
func (this *RoleAssignment) Insert() (isInsert bool, err error) {
	if this.RoleId <= 0 {
		return false, errors.New("角色ID不能为空")
	}

	if this.UserId <= 0 {
		return false, errors.New("用户ID不能为空")
	}

	self := &RoleAssignment{}
	o := orm.NewOrm()
	o.QueryTable(this.TableName()).Filter("role_id", this.RoleId).Filter("user_id", this.UserId).One(self)

	if self.Id > 0 {
		return false, errors.New("角色和用户已经绑定")
	}

	this.CreateAt = int32(time.Now().Unix())

	id, err := o.Insert(this)

	return id > 0, err
}

//remove current data from database
func (this *RoleAssignment) Delete() (isDelete bool, err error) {
	if this.Id <= 0 {
		return false, errors.New("数据不能为空")
	}

	o := orm.NewOrm()
	num, err := o.Delete(this)

	return num > 0, err
}

//retrieve all route
func (this RoleAssignment) FindRoleIdsByUserId(userId int) ([]int, error) {
	var items []*RoleAssignment
	roleIds := make([]int, 0)

	if userId <= 0 {
		return roleIds, errors.New("用户ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("user_id", userId).All(&items)

	if err != nil {
		return roleIds, err
	}

	for _, roleAssignment := range items {
		roleIds = append(roleIds, roleAssignment.RoleId)
	}

	return roleIds, nil
}
