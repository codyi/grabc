package controller

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RouteController struct {
	beego.Controller
}

func (this *RouteController) Index() {
	fmt.Println("router")
}
