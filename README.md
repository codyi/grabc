## GRABC 
[![GitHub forks](https://img.shields.io/github/forks/codyi/grabc.svg?style=social&label=Forks)](https://github.com/codyi/grabc/network)
[![GitHub stars](https://img.shields.io/github/stars/codyi/grabc.svg?style=social&label=Starss)](https://github.com/codyi/grabc/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/codyi/grabc.svg)](https://github.com/codyi/grabc)
[![Go Report Card](https://goreportcard.com/badge/github.com/codyi/grabc)](https://goreportcard.com/report/github.com/codyi/grabc)  

GRABC 是一个beego权限管理插件，插件分为路由、权限、角色。将路由分配给权限，权限授给角色，角色授给用户~~

### 安装
    go get github.com/codyi/grabc

### 配置    
第一步：在你项目中的数据库中导入rabc.sql，生成对应数据表

第二步：在项目中引入grabc库（可以在项目中的main.go或router.go中引入）

<pre>
//引入grabc库
import "github.com/codyi/grabc"
</pre>

引入之后，在引入的router.go或main.go中添加如下配置
<pre>
func init() {
	//将路由注册到grabc，用于反射出对应的网址
	grabc.RegisterController(& controllers.SiteController{})
	grabc.RegisterController(&controllers.UserController{})
	//注册用户系统模型到grabc，用于用户ID和grabc插件绑定
	//注意：注册的这个用户模型，需要实现IUserModel中的方法
	grabc.RegisterUserModel(&models.User{})
	//增加忽律权限检查的页面
	grabc.AppendIgnoreRoute("site", "login")
	//403页面地址注册到grabc中，用于grabc插件禁止权限的页面跳转
	grabc.Http_403("/site/nopermission")
	//设置模板，为了让grabc更具有通用性，可以设置模板
	//目前设置模板只支持传入模板的内容
	grabc.SetLayout(libs.Grabc_layout, nil)
}
</pre>

添加好上面的配置之后，剩下就是在controller中增加权限判了，个人建议做一个BaseController，然后每个controller都继承这个base，然后在BaseController中的Prepare方法中增加grabc的权限检查~~
<pre>
//注册当前登录的用户，注意：user需要继承IUserIdentify接口
grabc.RegisterIdentify(user)

if !grabc.CheckAccess(this.controllerName, this.actionName) {
	this.redirect(this.URLFor("SiteController.NoPermission"))
}
</pre>

到此grabc的功能都加完了，是不是很简单~~~

注意：增加完权限判断之后，会发现很多页面都不能访问了，那么就在忽律权限中增加如下配置
<pre>
grabc.AppendIgnoreRoute("\*", "\*")
</pre>
以上配置将会忽律所有的权限检查，这时候需要去/route/index中增加路由，然后添加权限，角色和用户分配，都配置好之后，就可以将grabc.AppendIgnoreRoute("\*", "\*")代码删掉，然后重启项目~~权限起作用了

### 接口说明    
IUserModel接口
<pre>
//用于定义用户model
type IUserModel interface {
	//用户列表返回可用用户的id和姓名
	//参数：pageIndex 分页的页数
	//参数：pageCount 每页显示的用户数量
	//返回值：userList [用户ID]用户姓名，用户列表展示
	//返回值：totalNum 全部的用户数目，用于计算分页的数量
	//返回值：err
	UserList(pageIndex, pageCount int) (userList map[int]string, totalNum int, err error)
	//根据用户ID获取用户姓名
	FindNameById(id int) string 
}
</pre>

IUserIdentify接口
<pre>
type IUserIdentify interface {
	GetId() int //返回当前登录用户的ID
}
</pre>
