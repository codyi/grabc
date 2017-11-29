package grbac

import (
	"fmt"
)

type Rbac struct {
}

func (this *Rbac) A(name string) {
	fmt.Println(name)
}
