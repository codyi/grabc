package libs

type Alert struct {
	Error_messages   []string
	Success_messages []string
	Info_messages    []string
	Warning_messages []string
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
