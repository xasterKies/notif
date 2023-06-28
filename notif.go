package notif


const (
	SeverityLow = iota
	SeverityNormal
	SeverityUrgent
)

type Severity int

type Notify struct {
	title string
	message string
	severity Severity
}

func New(title, message string, severity Severity) *Notify {
	return &Notify{
		title:	title,
		message: message,
		severity: severity,
	}
}