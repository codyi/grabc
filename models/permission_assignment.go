package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type PermissionAssignment struct {
	BaseModel
	RouteId      int `json:"route_id" label:"路由ID"`
	PermissionId int `json:"permission_id" label:"权限ID"`
}

func (this *PermissionAssignment) TableName() string {
	return "rabc_permission_assignment"
}

//Find one PermissionAssignment by id from database
func (this *PermissionAssignment) FindById(id int) error {
	if id <= 0 {
		return errors.New("ID不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

//insert current permission to database
//not insert if name is exist
func (this *PermissionAssignment) Insert() (isInsert bool, err error) {
	if this.RouteId <= 0 {
		return false, errors.New("路由ID不能为空")
	}

	if this.PermissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}

	self := &PermissionAssignment{}
	o := orm.NewOrm()
	o.QueryTable(this.TableName()).Filter("route_id", this.RouteId).Filter("permission_id", this.PermissionId).One(self)

	if self.Id > 0 {
		return false, errors.New("路由和权限已经绑定")
	}

	this.CreateAt = int32(time.Now().Unix())

	id, err := o.Insert(this)

	return id > 0, err
}

//remove current name from database
func (this *PermissionAssignment) Delete(routeId, permissionId int) (isDelete bool, err error) {
	if permissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}
	if routeId <= 0 {
		return false, errors.New("路由ID不能为空")
	}

	o := orm.NewOrm()
	model := &PermissionAssignment{}
	o.QueryTable(this.TableName()).Filter("route_id", routeId).Filter("permission_id", permissionId).One(model)

	if model.Id <= 0 {
		return false, errors.New("数据不存在")
	}
	num, err := o.Delete(model)

	return num > 0, err
}

//retrieve all PermissionAssignment
func (this PermissionAssignment) FindAllByPermissionId(permissionId int) ([]*PermissionAssignment, error) {
	var permissionAssignments []*PermissionAssignment

	if permissionId <= 0 {
		return permissionAssignments, errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("permission_id", permissionId).All(&permissionAssignments)

	return permissionAssignments, err
}

//remove current name from database
func (this *PermissionAssignment) DeleteByPermissionId(permissionId int) (err error) {
	if permissionId <= 0 {
		return errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from "+this.TableName()+" where permission_id=?", permissionId).Exec()
	return err
}
