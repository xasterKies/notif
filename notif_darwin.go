package notif

import "os/exec"


var command = exec.Command

// Send notification function for MacOs
func(n *Notify) Send() error {
	notifCmdName := "terminal-notifier"

	notifCmd, err := exec.LookPath(notifCmdName)
	if err != nil {
		return err
	}

	title := fmt.Sprintf("(%s) %s", n.severity, n.title)

	notifCommand := command(notifCmd, "-title", title, "-message", n.message)
	return notifCommand.Run()
}