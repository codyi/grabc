package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type AssignmentRoute struct {
	BaseModel
	RouteId      int `json:"route_id" label:"路由ID"`
	PermissionId int `json:"permission_id" label:"权限ID"`
}

func (this *AssignmentRoute) TableName() string {
	return "rabc_assignment_route"
}

//Find one AssignmentRoute by id from database
func (this *AssignmentRoute) FindById(id int) error {
	if id <= 0 {
		return errors.New("ID不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

//insert current permission to database
//not insert if name is exist
func (this *AssignmentRoute) Insert() (isInsert bool, err error) {
	if this.RouteId <= 0 {
		return false, errors.New("路由ID不能为空")
	}

	if this.PermissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}

	self := &AssignmentRoute{}
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
func (this *AssignmentRoute) Delete(routeId, permissionId int) (isDelete bool, err error) {
	if permissionId <= 0 {
		return false, errors.New("权限ID不能为空")
	}
	if routeId <= 0 {
		return false, errors.New("路由ID不能为空")
	}

	o := orm.NewOrm()
	model := &AssignmentRoute{}
	o.QueryTable(this.TableName()).Filter("route_id", routeId).Filter("permission_id", permissionId).One(model)

	if model.Id <= 0 {
		return false, errors.New("数据不存在")
	}
	num, err := o.Delete(model)

	return num > 0, err
}

//retrieve all AssignmentRoute
func (this AssignmentRoute) FindAllByPermissionId(permissionId int) ([]*AssignmentRoute, error) {
	var AssignmentRoutes []*AssignmentRoute

	if permissionId <= 0 {
		return AssignmentRoutes, errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("permission_id", permissionId).All(&AssignmentRoutes)

	return AssignmentRoutes, err
}

//remove current name from database
func (this *AssignmentRoute) DeleteByPermissionId(permissionId int) (err error) {
	if permissionId <= 0 {
		return errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from "+this.TableName()+" where permission_id=?", permissionId).Exec()
	return err
}

//remove current name from database
func (this *AssignmentRoute) DeleteByRouteId(routeId int) (err error) {
	if routeId <= 0 {
		return errors.New("路由ID不能为空")
	}

	o := orm.NewOrm()
	_, err = o.Raw("delete from "+this.TableName()+" where route_id=?", routeId).Exec()
	return err
}

//获取传入权限id的获取全部的路由ids
func (this AssignmentRoute) FindRouteIdsByPerIds(ids []int) ([]int, error) {
	var ra []*AssignmentRoute
	var routeIds []int
	if len(ids) == 0 {
		return routeIds, errors.New("权限ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("permission_id__in", ids).All(&ra)

	if err != nil {
		return routeIds, err
	}

	for _, r := range ra {
		routeIds = append(routeIds, r.RouteId)
	}

	return routeIds, nil
}
