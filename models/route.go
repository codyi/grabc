package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Route struct {
	BaseModel
	Url string `json:"url" label:"路由地址"`
}

func (this *Route) TableName() string {
	return "rabc_route"
}

//Find one user by phone from database
func (this *Route) FindById(id int) error {
	if id <= 0 {
		return errors.New("id不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

//Find one user by phone from database
func (this *Route) FindByUrl(url string) error {
	if url == "" {
		return errors.New("路由地址不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("url", url).One(this)
}

//insert current route to database
//not insert if route is exist
func (this *Route) Insert() (isInsert bool, err error) {
	if this.Url == "" {
		return false, errors.New("路由地址不能为空")
	}

	self := Route{}
	self.FindByUrl(this.Url)
	if self.Id > 0 {
		return false, errors.New("路由地址已经存在")
	}

	this.CreateAt = int32(time.Now().Unix())
	o := orm.NewOrm()

	id, err := o.Insert(this)

	return id > 0, err
}

//remove current url from database
func (this Route) DeleteByRoute(url string) (isDelete bool, err error) {
	if url == "" {
		return false, errors.New("路由地址不能为空")
	}
	o := orm.NewOrm()

	routeModel := &Route{}
	o.QueryTable(this.TableName()).Filter("url", url).One(routeModel)

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

//retrieve all route
func (this Route) FindAllByIds(routeIds []int) ([]*Route, error) {
	var routes []*Route

	if len(routeIds) == 0 {
		return routes, errors.New("路由ID不能为空")
	}

	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("id__in", routeIds).All(&routes)

	return routes, err
}

//根据用户ID获取对用的路由
func (this Route) ListByUserId(user_id int) map[string][]string {
	routes := make(map[string][]string, 0)

	if user_id <= 0 {
		return routes
	}

	roleAssignment := AssignmentRole{}
	roleIds, err := roleAssignment.FindRoleIdsByUserId(user_id)

	if err != nil {
		return routes
	}

	permissionAssignment := AssignmentPermission{}
	permissionIds, err := permissionAssignment.FindPerIdsByRoleIds(roleIds)

	if err != nil {
		return routes
	}

	routeAssignment := AssignmentRoute{}
	routeIds, err := routeAssignment.FindRouteIdsByPerIds(permissionIds)

	if err != nil {
		return routes
	}

	rs, err := this.FindAllByIds(routeIds)

	if err != nil {
		return routes
	}

	for _, r := range rs {
		t := strings.Split(r.Url, "/")
		if routes[t[0]] == nil {
			routes[t[0]] = make([]string, 0)
		}
		routes[t[0]] = append(routes[t[0]], t[1])
	}

	return routes
}
