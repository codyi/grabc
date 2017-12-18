package layout

func GetGlobalJs() string {
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
