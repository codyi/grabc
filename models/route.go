package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Route struct {
	BaseModel
	Route string `json:"route" label:"路由地址"`
}

func (this *Route) TableName() string {
	return "rabc_route"
}

//Find one user by phone from database
func (this *Route) FindByRoute(route string) error {
	if route == "" {
		return errors.New("路由地址不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("route", route).One(this)
}

//insert current route to database
//not insert if route is exist
func (this *Route) Insert() (isInsert bool, err error) {
	if this.Route == "" {
		return false, errors.New("路由地址不能为空")
	}

	self := Route{}
	self.FindByRoute(this.Route)
	if self.Id > 0 {
		return false, errors.New("路由地址已经存在")
	}

	this.CreateAt = int32(time.Now().Unix())
	o := orm.NewOrm()

	id, err := o.Insert(this)

	return id > 0, err
}

//remove current route from database
func (this Route) DeleteByRoute(route string) (isDelete bool, err error) {
	if route == "" {
		return false, errors.New("路由地址不能为空")
	}
	o := orm.NewOrm()

	routeModel := &Route{}
	o.QueryTable(this.TableName()).Filter("route", route).One(routeModel)

	num, err := o.Delete(routeModel)

	return num > 0, err
}

//retrieve all route
func (this Route) FindAll() ([]*Route, error) {
	var routes []*Route
	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).All(&routes)

	return routes, err
}
