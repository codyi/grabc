package permission

import (
	. "grabc/views/layout"
)

type Add struct {
	BaseTemplate
}

func (this *Add) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <form action="/permission/add" method="post" class="form-horizontal">
            <div class="form-group required">
                <label class="control-label col-sm-2" for="permission_name">名称</label>
                <div class="col-sm-4">
                    <input type="text" id="permission_name" class="form-control" name="permission_name" maxlength="20">
                </div>
            </div>
            <div class="form-group">
                <label class="control-label col-sm-2" for="permission_desc">简介</label>
                <div class="col-sm-4">
                    <input type="text" id="permission_desc" class="form-control" name="permission_desc" maxlength="200">
                </div>
            </div>
            <div class="form-group">
                <label class="control-label col-sm-2">&nbsp;</label>
                <div class="col-sm-4">
                    <button type="submit" class="btn btn-success">创建</button>
                </div>
            </div>
        </form>
    </div>
</div>
`
	return this.DealHtml(html)
}
