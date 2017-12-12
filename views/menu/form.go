package menu

import (
	. "github.com/codyi/grabc/views/layout"
)

type Form struct {
	BaseTemplate
}

func (this *Form) FormHtml() string {
	html := `
<div class="form-group required">
    <label class="control-label col-sm-2" for="menu_name">菜单名称</label>
    <div class="col-sm-4">
        <input type="text" id="menu_name" class="form-control" name="menu_name" maxlength="20" value="{{.model.Name}}">
    </div>
</div>
<div class="form-group">
    <label class="control-label col-sm-2" for="menu_route">菜单路由</label>
    <div class="col-sm-4">
        <input type="text" id="menu_route" class="form-control" name="menu_route" maxlength="200" value="{{.model.Route}}">
    </div>
</div>
<div class="form-group">
    <label class="control-label col-sm-2" for="menu_order">菜单排序</label>
    <div class="col-sm-4">
        <input type="text" id="menu_order" class="form-control" name="menu_order" maxlength="200" value="{{.model.Order}}">
    </div>
</div>
<div class="form-group">
    <label class="control-label col-sm-2" for="menu_parent">父级菜单</label>
    <div class="col-sm-4">
        <input type="text" id="menu_parent" class="form-control" name="menu_parent" maxlength="200" value="{{.model.Parent}}">
    </div>
</div>
<div class="form-group">
    <label class="control-label col-sm-2">&nbsp;</label>
    <div class="col-sm-4">
        <button type="submit" class="btn btn-success">保存</button>
    </div>
</div>
`
	return html
}
