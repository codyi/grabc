package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

type Menu struct {
	BaseModel
	Name       string `json:"name" label:"菜单名称"`
	Parent     int    `json:"parent" label:"父级菜单ID"`
	Route      string `json:"route" label:"菜单地址"`
	Order      int    `json:"order" label:"菜单排序"`
	ParentName string `orm:"-" label:"父级菜单名称"`
}

func (this *Menu) TableName() string {
	return "rabc_menu"
}

//Find one Menu by id from database
func (this *Menu) FindById(id int) error {
	if id <= 0 {
		return errors.New("菜单ID不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("id", id).One(this)
}

func (this *Menu) GetParentName() string {
	if this.Parent > 0 {
		m := Menu{}
		m.FindById(this.Parent)
		return m.Name
	}

	return "-"
}

//Find one Menu by name from database
func (this *Menu) FindByName(name string) error {
	if name == "" {
		return errors.New("菜单名称不能为空")
	}

	o := orm.NewOrm()
	return o.QueryTable(this.TableName()).Filter("name", name).One(this)
}

//insert current Menu to database
//not insert if name is exist
func (this *Menu) Insert() (isInsert bool, err error) {
	if this.Name == "" {
		return false, errors.New("菜单名称不能为空")
	}

	if this.Route == "" {
		return false, errors.New("菜单路由不能为空")
	}

	this.CreateAt = int32(time.Now().Unix())
	o := orm.NewOrm()

	id, err := o.Insert(this)

	return id > 0, err
}

//update current Menu to database
//not update if name is exist
func (this *Menu) Update() (err error) {
	if this.Name == "" {
		return errors.New("菜单名称不能为空")
	}

	if this.Route == "" {
		return errors.New("菜单路由不能为空")
	}

	o := orm.NewOrm()

	_, err = o.Update(this)

	return err
}

//retrieve all route
func (this Menu) FindAllParent() ([]*Menu, error) {
	var menus []*Menu
	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Filter("parent", 0).All(&menus)

	return menus, err
}

//remove current menu from database
func (this *Menu) Delete() (isDelete bool, err error) {
	if this.IsNewRecord() {
		return false, errors.New("删除对象不能为空")
	}

	o := orm.NewOrm()
	MenuModel := &Menu{}
	o.QueryTable(this.TableName()).Filter("parent", this.Id).One(MenuModel)

	if !MenuModel.IsNewRecord() {
		return false, errors.New("存在子菜单，不能删除")
	}

	num, err := o.Delete(this)

	return num > 0, err
}

//retrieve all menus
func (this Menu) List(pageIndex, pageCount int) ([]*Menu, int, error) {
	var menus []*Menu
	var total int64
	o := orm.NewOrm()
	_, err := o.QueryTable(this.TableName()).Limit(pageCount).Offset(pageCount * (pageIndex - 1)).All(&menus)

	if err != nil {
		return menus, int(total), err
	}

	total, err = o.QueryTable(this.TableName()).Count()
	return menus, int(total), err
}
