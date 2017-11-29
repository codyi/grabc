package views

func getGlobalJs() string {
	return `
<script>
$(function(){
	$("#add_route").click(function(){
		var selectRoutes = $("#route_select").val();

		if (selectRoutes.length > 0) {
			
		}
	});

	$.setActiveMenu();
});

//设置导航选中状态
$.setActiveMenu = function(){
	var pathname = window.location.pathname;

	if (pathname.substr(0, 1) == "/") {
		pathname = pathname.substr(1, pathname.length)
	}

	var controllerName = pathname.split("/")[0].toLocaleLowerCase(); //字符分割 


	$("#top_menu").children("li").each(function(){
		var url = $(this).children("a").attr("href").toLocaleLowerCase();
		
		if (url.indexOf(controllerName) == 1) {
			$(this).children("a").addClass("active")
		}
	});
}
</script>
`
}
