// Code generated by qtc from "Capital.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/Capital.sql:1
package ddl

//line queries/ddl/Capital.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/Capital.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/Capital.sql:1
func StreamCapitalDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/Capital.sql:1
	qw422016.N().S(`
drop table if exists "Capital";
-- `)
//line queries/ddl/Capital.sql:3
}

//line queries/ddl/Capital.sql:3
func WriteCapitalDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/Capital.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/Capital.sql:3
	StreamCapitalDrop(qw422016)
//line queries/ddl/Capital.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/Capital.sql:3
}

//line queries/ddl/Capital.sql:3
func CapitalDrop() string {
//line queries/ddl/Capital.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/Capital.sql:3
	WriteCapitalDrop(qb422016)
//line queries/ddl/Capital.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/Capital.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/Capital.sql:3
	return qs422016
//line queries/ddl/Capital.sql:3
}

// --

//line queries/ddl/Capital.sql:5
func StreamCapitalCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/Capital.sql:5
	qw422016.N().S(`
create table if not exists "Capital" (
  "ID" text not null,
  "Name" text not null,
  "Birthday" timestamp not null,
  "Deathday" timestamp,
  primary key ("ID")
);
-- `)
//line queries/ddl/Capital.sql:13
}

//line queries/ddl/Capital.sql:13
func WriteCapitalCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/Capital.sql:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/Capital.sql:13
	StreamCapitalCreate(qw422016)
//line queries/ddl/Capital.sql:13
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/Capital.sql:13
}

//line queries/ddl/Capital.sql:13
func CapitalCreate() string {
//line queries/ddl/Capital.sql:13
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/Capital.sql:13
	WriteCapitalCreate(qb422016)
//line queries/ddl/Capital.sql:13
	qs422016 := string(qb422016.B)
//line queries/ddl/Capital.sql:13
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/Capital.sql:13
	return qs422016
//line queries/ddl/Capital.sql:13
}
