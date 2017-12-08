package permission

import (
	. "github.com/grabc/views/layout"
)

type Form struct {
	BaseTemplate
}

func (this *Form) FormHtml() string {
	html := `
<div class="form-group required">
    <label class="control-label col-sm-2" for="permission_name">名称</label>
    <div class="col-sm-4">
        <input type="text" id="permission_name" class="form-control" name="permission_name" maxlength="20" value="{{.model.Name}}">
    </div>
</div>
<div class="form-group">
    <label class="control-label col-sm-2" for="permission_desc">简介</label>
    <div class="col-sm-4">
        <input type="text" id="permission_desc" class="form-control" name="permission_desc" maxlength="200" value="{{.model.Description}}">
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
