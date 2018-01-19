package libs

type Alert struct {
	Error_messages   []string
	Success_messages []string
	Info_messages    []string
	Warning_messages []string
}

func (this *Alert) HasError() bool {
	return len(this.Error_messages) > 0
}
func (this *Alert) AddErrorMessage(message string) {
	this.Error_messages = append(this.Error_messages, message)
}

func (this *Alert) AddSuccessMessage(message string) {
	this.Success_messages = append(this.Success_messages, message)
}

func (this *Alert) AddInfoMessage(message string) {
	this.Info_messages = append(this.Info_messages, message)
}

func (this *Alert) AddWarningMessage(message string) {
	this.Warning_messages = append(this.Warning_messages, message)
}

func (this *Alert) ShowAlert() string {
	var html string

	if len(this.Error_messages) > 0 {
		html += "<div class=\"alert alert-danger fade in\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-hidden=\"true\">×</button><i class='icon fa fa-ban'></i>"
		for _, m := range this.Error_messages {
			html += "<div>" + m + "</div>"
		}

		html += "</div>"
	}

	if len(this.Info_messages) > 0 {
		html += "<div class=\"alert alert-info fade in\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-hidden=\"true\">×</button><i class='icon fa fa-info'></i>"
		for _, m := range this.Info_messages {
			html += "<div>" + m + "</div>"
		}

		html += "</div>"
	}

	if len(this.Success_messages) > 0 {
		html += "<div class=\"alert alert-success fade in\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-hidden=\"true\">×</button><i class='icon fa fa-check'></i>"
		for _, m := range this.Success_messages {
			html += "<div>" + m + "</div>"
		}

		html += "</div>"
	}

	if len(this.Warning_messages) > 0 {
		html += "<div class=\"alert alert-warning fade in\"><button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-hidden=\"true\">×</button><i class='icon fa fa-warning'></i>"
		for _, m := range this.Warning_messages {
			html += "<div>" + m + "</div>"
		}

		html += "</div>"
	}

	return html
}
