package libs

type Breadcrumb struct {
	Label string
	Url   string
}
type Breadcrumbs struct {
	Items []Breadcrumb
}

//set Breadcrumbs
func (this *Breadcrumbs) AddBreadcrumbs(label, url string) {
	if this.Items == nil {
		this.Items = make([]Breadcrumb, 0)
	}

	breadcrumb := Breadcrumb{}
	breadcrumb.Label = label
	breadcrumb.Url = url

	this.Items = append(this.Items, breadcrumb)
}
