package notif


var command = exec.Command

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