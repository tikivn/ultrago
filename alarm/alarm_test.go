package alarm

import (
	"fmt"
	"testing"
	"time"
)

func TestNewAlarm(t *testing.T) {
	a := New(16, 10, 24*time.Hour)
	a.Restart()
	for {
		select {
		case t := <-a.T.C:
			fmt.Println(t.String())
			a.Restart()
		}
	}
}
