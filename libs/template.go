package libs

import (
	"os"
	"path/filepath"
)

//存储用户自定义的layout内容
var Template GrabcTemplate

func init() {
	Template = GrabcTemplate{}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err.Error())
	}

	Template.ViewPath = filepath.Dir(dir) + "/github.com/codyi/grabc/views/"
	Template.Layout = filepath.Dir(dir) + "/github.com/codyi/grabc/views/layout/main.html"
	Template.Data = make(map[string]interface{}, 0)

}

type GrabcTemplate struct {
	Layout   string
	ViewPath string
	Data     map[string]interface{}
}

func (this *GrabcTemplate) GlobalCss() string {
	return `
<style type="text/css">
    .row_operate{
        width: 15%;
    }
    /*route page css*/
    table.route_warp tr td{
        border: none;
    }
    table.route_warp tr td select{
        width: 80%;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    table.route_warp tr td:nth-child(1){
        text-align: right;
    }
    table.route_warp tr td:nth-child(2){
        text-align: center;
    }
    table.route_warp tr td:nth-child(3){
        text-align: left;
    }
    select[multiple]{
        height: 500px;
    }
</style>
`
}

func (this *GrabcTemplate) GlobalJs() string {
	return `
<script>
$(function(){
    $("[data-confirm]").click(function () {
        if (confirm($(this).attr("data-confirm"))) {
            var oThis = this;
            $.post($(this).attr("href"), function(response){
                if (response.Code == 200) {
                    $(oThis).parent().parent().remove();
                } else {
                    alert(response.Message);
                }
            });
        }

        return false;
    });

    //显示select显示的数据
    $.showSelectOption = function (select_id, items) {
        $(select_id).empty()
        items = items.sort()
        $(items).each(function(index, value){
            $(select_id).append("<option value='" + value + "'>" + value + "</option>")
        });
    };

    //将item从items中删除
    $.removeItem = function(item, items){
        var new_items = new Array();
        for (var i in items) {
            if (items[i] != item) {
                new_items.push(items[i]);
            }
        }

        return new_items;
    }
});
</script>
`
}
