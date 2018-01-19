## GRABC 
[![GitHub forks](https://img.shields.io/github/forks/codyi/grabc.svg?style=social&label=Forks)](https://github.com/codyi/grabc/network)
[![GitHub stars](https://img.shields.io/github/stars/codyi/grabc.svg?style=social&label=Starss)](https://github.com/codyi/grabc/stargazers)
[![GitHub last commit](https://img.shields.io/github/last-commit/codyi/grabc.svg)](https://github.com/codyi/grabc)
[![Go Report Card](https://goreportcard.com/badge/github.com/codyi/grabc)](https://goreportcard.com/report/github.com/codyi/grabc)  

GRABC 是一个beego权限管理插件，插件分为路由、权限、角色。将路由分配给权限，权限授给角色，角色授给用户~~
GRABC 目前依赖的前端是adminlte和boostrap

### 安装
    go get github.com/codyi/grabc

### 配置    
第一步：在你项目中的数据库中导入rabc.sql，生成对应数据表

第二步：在项目中引入grabc库（可以在项目中的main.go或router.go中引入）

<pre>
import "github.com/codyi/grabc"
</pre>

引入之后，在引入的router.go或main.go中添加如下配置
<pre>
func init() {
	var c []beego.ControllerInterface
	c = append(c, &controllers.SiteController{}, &controllers.UserController{})
	beego.Router("/", &controllers.SiteController{})
	for _, v := range c {
		//将路由注册到beego
		beego.AutoRouter(v)
		//将路由注册到grabc
		grabc.RegisterController(v)
	}
	//注册用户系统模型到grabc
	grabc.RegisterUserModel(&models.User{})
	//增加忽律权限检查的页面
	grabc.AppendIgnoreRoute("site", "login")

	//设置grabc页面视图路径
	//如果使用默认的，不要设置或置空
	//如果需要对grabc插件进行二次开发，则需要设置这个目录，否则不需要管
	//注意：设置grabc的模板必须在beego.Run()之前设置，如果视图目录在当前项目中，可以使用相对目录，否则需要绝对路径
	// grabc.SetViewPath("views")
	//设置grabc的layout
	grabc.SetLayout("layout/main.html", "views")
}
</pre>

添加好上面的配置之后，剩下就是在controller中增加权限判了，个人建议做一个BaseController，然后每个controller都继承这个base，然后在BaseController中的Prepare方法中增加grabc的权限检查~~
<pre>
//注册当前登录的用户，注意：user需要继承IUserIdentify接口
grabc.RegisterIdentify(user)
//检查权限
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

### 注意
grabc对注册的控制器会进行反射，然后获取每个controller的名称和controller内的公共方法，由于每个controller都继承了beego.Controller，在获取controller下的方法名称时，会将beego.Controller继承的方法也会获取到，所以目前还不能区分出方法名到底是beego和用户自己定义的，所以grabc将beego继承的方法都进行了忽律，如果在route扫描中，没有找到自定义的方法，可以在controller中增加如下方法，进行方法返回~~
<pre>
func (this *SiteController) RABCMethods() []string {
	return []string{"Get", "Post"}
}
</pre>

grabc的详细例子：github.com/codyi/grabc_example

![Image text](http://www.liguosong.com/grabc_1.jpeg)
![Image text](http://www.liguosong.com/grabc_2.jpeg)
![Image text](http://www.liguosong.com/grabc_3.jpeg)
![Image text](http://www.liguosong.com/grabc_4.jpeg)