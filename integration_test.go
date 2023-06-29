package notif_test

import (
	"testing"

	"github.com/xasterKies/notif"
)

func TestSend(t *testing.T) {
	n := notif.New("test title", "test msg", notif.SeverityNormal)

	err := n.Send()

	if err != nil {
		t.Error(err)
	}
}