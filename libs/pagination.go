package libs

import (
	"math"
	"strconv"
)

type Pagination struct {
	PageIndex int
	PageCount int
	PageTotal int
	Url       string
}

//输出分页
func PaginationRender(page Pagination) string {
	pageHtml := `<nav aria-label="Page navigation" class="pagination_warp"><ul class="pagination">`

	if page.PageIndex == 1 {
		pageHtml += `<li><a href="#" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a></li>`
	} else {
		pageHtml += `<li><a href="` + page.Url + `?page_index=` + strconv.Itoa(page.PageIndex-1) + `" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a></li>`
	}

	for i := 1; i <= int(math.Ceil(float64(page.PageTotal)/float64(page.PageCount))); i++ {
		pageHtml += `<li><a href="` + page.Url + `?page_index=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a></li>`
	}

	if page.PageIndex == int(math.Ceil(float64(page.PageTotal)/float64(page.PageCount))) {
		pageHtml += `<li><a href="#" aria-label="Next"><span aria-hidden="true">&raquo;</span></a></li>`
	} else {
		pageHtml += `<li><a href="` + page.Url + `?page_index=` + strconv.Itoa(page.PageIndex+1) + `" aria-label="Next"><span aria-hidden="true">&raquo;</span></a></li>`
	}

	pageHtml += `</ul></nav>`

	return pageHtml
}
