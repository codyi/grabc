package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type AssignmentPermission struct {
	BaseModel
	RoleId       int `json:"role_id" label:"角色ID"`
	PermissionId int `json:"permission_id" label:"权限ID"`
}

func (this *AssignmentPermission) TableName() string {
	return "rabc_assignment_permission"
}

//insert current permission to database
//not insert if name is exist
func (this *AssignmentPermission) Insert() (isInsert bool, err error) {
	if this.RoleId <= 0 {
		return false, errors.New("角色ID不能为空")
	}

	if this.PermissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}

	self := &AssignmentPermission{}
	o := orm.NewOrm()
	o.QueryTable(this.TableName()).Filter("role_id", this.RoleId).Filter("permission_id", this.PermissionId).One(self)

	if self.Id > 0 {
		return false, errors.New("角色和权限已经绑定")
	}

	this.CreateAt = int32(time.Now().Unix())

	id, err := o.Insert(this)

	return id > 0, err
}

//remove current name from database
func (this *AssignmentPermission) Delete(roleId, permissionId int) (isDelete bool, err error) {
	if permissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}
	if roleId <= 0 {
		return false, errors.New("角色ID不能为空")
	}

	o := orm.NewOrm()
	model := &AssignmentPermission{}
	o.QueryTable(this.TableName()).Filter("role_id", roleId).Filter("permission_id", permissionId).One(model)

	if model.Id <= 0 {
		return false, errors.New("数据不存在")
	}
	num, err := o.Delete(model)

	return num > 0, err
}

//retrieve all AssignmentPermission
func (this AssignmentPermission) FindAllByRoleId(roleId int) ([]*AssignmentPermission, error) {
	var AssignmentPermissions []*AssignmentPermission

	if roleId <= 0 {
		return AssignmentPermissions, errors.New("角色ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("role_id", roleId).All(&AssignmentPermissions)

	return AssignmentPermissions, err
}

//获取传入角色id的名称获取全部的权限ids
func (this AssignmentPermission) FindPerIdsByRoleIds(ids []int) ([]int, error) {
	var pas []*AssignmentPermission
	var roleIds []int
	if len(ids) == 0 {
		return roleIds, errors.New("角色ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("role_id__in", ids).All(&pas)

	if err != nil {
		return roleIds, err
	}

	for _, r := range pas {
		roleIds = append(roleIds, r.PermissionId)
	}

	return roleIds, nil
}

//remove current name from database
func (this *AssignmentPermission) DeleteByPermissionId(permissionId int) (err error) {
	if permissionId <= 0 {
		return errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from "+this.TableName()+" where permission_id=?", permissionId).Exec()
	return err
}

//remove current  from database
func (this *AssignmentPermission) DeleteByRoleId(roleId int) (err error) {
	if roleId <= 0 {
		return errors.New("角色ID不能为空")
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from "+this.TableName()+" where role_id=?", roleId).Exec()
	return err
}
