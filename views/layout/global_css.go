package layout

func GetGlobalCss() string {
	return `
	<style type="text/css">
/*global css*/
html,body{
    background: #ecf0f5;
    height: 100%;
}
.navbar-inverse .navbar-collapse, .navbar-inverse .navbar-form,.navbar-inverse{
    background-color: #3c8dbc;
    border-color:#3c8dbc;
}
.navbar-inverse .navbar-nav>li>a{
    color: white
}
.navbar-nav>li>a.active{
    background-color: #367fa9;
}
.content_warp{
    padding-top: 80px;
    height: 100%;
}
.pagination_warp{
    text-align: center;
}
.box{
    border-radius: 3px;
    margin-bottom: 20px;
    border-top: 3px solid #d2d6de;
    background-color: #ffffff;
    box-shadow: 0 1px 1px rgba(0,0,0,0.1);
}
.box.box-primary{
    border-top-color:#3c8dbc;
}
.box.box-info{
    border-top-color:#00c0ef;
}
.box.box-info > .box-header {
    background: #f4f4f4;
    background-color: #f4f4f4;
    padding: 10px;
}
.box.box-info > .box-header > span{
    font-weight: bold;
}
.box-body {
    border-top-left-radius: 0;
    border-top-right-radius: 0;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px;
    padding: 10px;
}
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
