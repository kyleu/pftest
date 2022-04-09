// Content managed by Project Forge, see [projectforge.md] for details.
package log

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/kyleu/pftest/app/util"
)

type Timer struct {
	Key     string             `json:"key"`
	Log     *zap.SugaredLogger `json:"-"`
	Timer   *util.Timer        `json:"-"`
	initial *util.Timer
	index   int
}

func NewTimer(key string, logger *zap.SugaredLogger) *Timer {
	return &Timer{Key: key, Log: logger, Timer: util.TimerStart(), initial: util.TimerStart()}
}

func (l *Timer) Lap(msg string, args ...any) {
	l.index++
	l.Log.Infof(fmt.Sprintf("[%s::%d] ", l.Key, l.index) + fmt.Sprintf(msg, args...) + " [" + util.MicrosToMillis(l.Timer.End()) + "]")
	l.Timer = util.TimerStart()
}

func (l *Timer) Complete() {
	msg := fmt.Sprintf("completed after [%d] steps in [%s]", l.index, util.MicrosToMillis(l.initial.End()))
	l.Log.Infof(fmt.Sprintf("[%s::%d] ", l.Key, l.index) + msg + " [" + util.MicrosToMillis(l.Timer.End()) + "]")
}
