// Code generated by qtc from "types.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/types.sql:2
package ddl

//line queries/ddl/types.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/types.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/types.sql:2
func StreamTypesDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/types.sql:2
	qw422016.N().S(`
drop type if exists "foo";
-- `)
//line queries/ddl/types.sql:4
}

//line queries/ddl/types.sql:4
func WriteTypesDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/types.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/types.sql:4
	StreamTypesDrop(qw422016)
//line queries/ddl/types.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/types.sql:4
}

//line queries/ddl/types.sql:4
func TypesDrop() string {
//line queries/ddl/types.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/types.sql:4
	WriteTypesDrop(qb422016)
//line queries/ddl/types.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/types.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/types.sql:4
	return qs422016
//line queries/ddl/types.sql:4
}

// --

//line queries/ddl/types.sql:6
func StreamTypesCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/types.sql:6
	qw422016.N().S(`
create type "foo" as enum ('a', 'b', 'c', 'd');
-- `)
//line queries/ddl/types.sql:8
}

//line queries/ddl/types.sql:8
func WriteTypesCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/types.sql:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/types.sql:8
	StreamTypesCreate(qw422016)
//line queries/ddl/types.sql:8
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/types.sql:8
}

//line queries/ddl/types.sql:8
func TypesCreate() string {
//line queries/ddl/types.sql:8
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/types.sql:8
	WriteTypesCreate(qb422016)
//line queries/ddl/types.sql:8
	qs422016 := string(qb422016.B)
//line queries/ddl/types.sql:8
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/types.sql:8
	return qs422016
//line queries/ddl/types.sql:8
}