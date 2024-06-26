// Code generated by qtc from "UUID.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/view/UUID.html:1
package view

//line views/components/view/UUID.html:1
import "github.com/google/uuid"

//line views/components/view/UUID.html:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/UUID.html:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/UUID.html:3
func StreamUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/view/UUID.html:4
	if value != nil {
//line views/components/view/UUID.html:5
		qw422016.E().S(value.String())
//line views/components/view/UUID.html:6
	}
//line views/components/view/UUID.html:7
}

//line views/components/view/UUID.html:7
func WriteUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/view/UUID.html:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/UUID.html:7
	StreamUUID(qw422016, value)
//line views/components/view/UUID.html:7
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/UUID.html:7
}

//line views/components/view/UUID.html:7
func UUID(value *uuid.UUID) string {
//line views/components/view/UUID.html:7
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/UUID.html:7
	WriteUUID(qb422016, value)
//line views/components/view/UUID.html:7
	qs422016 := string(qb422016.B)
//line views/components/view/UUID.html:7
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/UUID.html:7
	return qs422016
//line views/components/view/UUID.html:7
}
