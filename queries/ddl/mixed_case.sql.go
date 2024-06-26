// Code generated by qtc from "mixed_case.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/mixed_case.sql:1
package ddl

//line queries/ddl/mixed_case.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/mixed_case.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/mixed_case.sql:1
func StreamMixedCaseDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/mixed_case.sql:1
	qw422016.N().S(`
drop table if exists "mixed_case";
-- `)
//line queries/ddl/mixed_case.sql:3
}

//line queries/ddl/mixed_case.sql:3
func WriteMixedCaseDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/mixed_case.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/mixed_case.sql:3
	StreamMixedCaseDrop(qw422016)
//line queries/ddl/mixed_case.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/mixed_case.sql:3
}

//line queries/ddl/mixed_case.sql:3
func MixedCaseDrop() string {
//line queries/ddl/mixed_case.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/mixed_case.sql:3
	WriteMixedCaseDrop(qb422016)
//line queries/ddl/mixed_case.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/mixed_case.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/mixed_case.sql:3
	return qs422016
//line queries/ddl/mixed_case.sql:3
}

// --

//line queries/ddl/mixed_case.sql:5
func StreamMixedCaseCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/mixed_case.sql:5
	qw422016.N().S(`
create table if not exists "mixed_case" (
  "id" text not null,
  "test_field" text not null,
  "another_field" text not null,
  primary key ("id")
);
-- `)
//line queries/ddl/mixed_case.sql:12
}

//line queries/ddl/mixed_case.sql:12
func WriteMixedCaseCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/mixed_case.sql:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/mixed_case.sql:12
	StreamMixedCaseCreate(qw422016)
//line queries/ddl/mixed_case.sql:12
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/mixed_case.sql:12
}

//line queries/ddl/mixed_case.sql:12
func MixedCaseCreate() string {
//line queries/ddl/mixed_case.sql:12
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/mixed_case.sql:12
	WriteMixedCaseCreate(qb422016)
//line queries/ddl/mixed_case.sql:12
	qs422016 := string(qb422016.B)
//line queries/ddl/mixed_case.sql:12
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/mixed_case.sql:12
	return qs422016
//line queries/ddl/mixed_case.sql:12
}
