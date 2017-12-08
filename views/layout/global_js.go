package layout

func GetGlobalJs() string {
	return `
<script>
//显示select显示的数据
$.showSelectOption = function (selectId, items) {
    $(selectId).empty()
    items = items.sort()
    $(items).each(function(index, value){
        $(selectId).append("<option value='" + value + "'>" + value + "</option>")
    });
};

//将item从items中删除
$.removeItem = function(item, items){
    var newItems = new Array();
    for (var i in items) {
        if (items[i] != item) {
            newItems.push(items[i]);
        }
    }

    return newItems;
}
</script>
`
}
