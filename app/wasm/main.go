// Content managed by Project Forge, see [projectforge.md] for details.
package main

import (
	lg "github.com/kyleu/pftest/app/lib/log"
	"github.com/kyleu/pftest/app/util"
)

var _rootLogger util.Logger

func main() {
	l, err := lg.InitLogging(true)
	if err != nil {
		println(err)
	}
	_rootLogger = l

	t := util.TimerStart()
	wireFunctions()
	l.Infof("[%s] started in [%s]", util.AppName, t.EndString())
	<-make(chan struct{})
}
