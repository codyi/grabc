package models

//用于定义用户model
type IUserModel interface {
	UserList(pageIndex, pageCount int) (userList map[int]string, totalNum int, err error) //用户实现应该返回可用用户的id和姓名
	FindNameById(id int) string
}

var UserModel *IUserModel
