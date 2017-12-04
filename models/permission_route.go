package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type PermissionRoute struct {
	BaseModel
	RouteId      int `json:"route_id" label:"路由ID"`
	PermissionId int `json:"permission_id" label:"权限ID"`
}

func (this *PermissionRoute) TableName() string {
	return "rabc_permission_route"
}

//Find one PermissionRoute by id from database
func (this *PermissionRoute) FindById(id int) error {
	if id <= 0 {
		return errors.New("ID不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

//insert current permission to database
//not insert if name is exist
func (this *PermissionRoute) Insert() (isInsert bool, err error) {
	if this.RouteId <= 0 {
		return false, errors.New("路由ID不能为空")
	}

	if this.PermissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}

	self := &PermissionRoute{}
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
func (this *PermissionRoute) Delete(routeId, permissionId int) (isDelete bool, err error) {
	if permissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}
	if routeId <= 0 {
		return false, errors.New("路由ID不能为空")
	}

	o := orm.NewOrm()
	model := &PermissionRoute{}
	o.QueryTable(this.TableName()).Filter("route_id", routeId).Filter("permission_id", permissionId).One(model)

	if model.Id <= 0 {
		return false, errors.New("数据不存在")
	}
	num, err := o.Delete(model)

	return num > 0, err
}

//retrieve all PermissionRoute
func (this PermissionRoute) FindAllByPermissionId(permissionId int) ([]*PermissionRoute, error) {
	var permissionRoutes []*PermissionRoute

	if permissionId <= 0 {
		return permissionRoutes, errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("permission_id", permissionId).All(&permissionRoutes)

	return permissionRoutes, err
}

//remove current name from database
func (this *PermissionRoute) DeleteByPermissionId(permissionId int) (err error) {
	if permissionId <= 0 {
		return errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from "+this.TableName()+" where permission_id=?", permissionId).Exec()
	return err
}
