// Code generated by qtc from "Option.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/fieldview/Option.html:2
package fieldview

//line views/components/fieldview/Option.html:2
import (
	"github.com/kyleu/pftest/app/lib/types"
)

//line views/components/fieldview/Option.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/fieldview/Option.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/fieldview/Option.html:6
func StreamOption(qw422016 *qt422016.Writer, x any, t *types.Option) {
//line views/components/fieldview/Option.html:7
	if x == nil {
//line views/components/fieldview/Option.html:7
		qw422016.N().S(`<em>∅</em>`)
//line views/components/fieldview/Option.html:9
	} else {
//line views/components/fieldview/Option.html:10
		StreamAny(qw422016, x, t.V)
//line views/components/fieldview/Option.html:11
	}
//line views/components/fieldview/Option.html:12
}

//line views/components/fieldview/Option.html:12
func WriteOption(qq422016 qtio422016.Writer, x any, t *types.Option) {
//line views/components/fieldview/Option.html:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/fieldview/Option.html:12
	StreamOption(qw422016, x, t)
//line views/components/fieldview/Option.html:12
	qt422016.ReleaseWriter(qw422016)
//line views/components/fieldview/Option.html:12
}

//line views/components/fieldview/Option.html:12
func Option(x any, t *types.Option) string {
//line views/components/fieldview/Option.html:12
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/fieldview/Option.html:12
	WriteOption(qb422016, x, t)
//line views/components/fieldview/Option.html:12
	qs422016 := string(qb422016.B)
//line views/components/fieldview/Option.html:12
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/fieldview/Option.html:12
	return qs422016
//line views/components/fieldview/Option.html:12
}