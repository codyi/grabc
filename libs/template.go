package libs

//存储用户自定义的layout内容
var Template GrabcTemplate

type GrabcTemplate struct {
	Layout string
	Data   map[string]interface{}
}
