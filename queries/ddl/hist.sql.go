// Code generated by qtc from "hist.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/hist.sql:2
package ddl

//line queries/ddl/hist.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/hist.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/hist.sql:2
func StreamHistDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/hist.sql:2
	qw422016.N().S(`
drop table if exists "hist_history";
drop table if exists "hist";
-- `)
//line queries/ddl/hist.sql:5
}

//line queries/ddl/hist.sql:5
func WriteHistDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/hist.sql:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/hist.sql:5
	StreamHistDrop(qw422016)
//line queries/ddl/hist.sql:5
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/hist.sql:5
}

//line queries/ddl/hist.sql:5
func HistDrop() string {
//line queries/ddl/hist.sql:5
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/hist.sql:5
	WriteHistDrop(qb422016)
//line queries/ddl/hist.sql:5
	qs422016 := string(qb422016.B)
//line queries/ddl/hist.sql:5
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/hist.sql:5
	return qs422016
//line queries/ddl/hist.sql:5
}

// --

//line queries/ddl/hist.sql:7
func StreamHistCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/hist.sql:7
	qw422016.N().S(`
create table if not exists "hist" (
  "id" text not null,
  "data" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  primary key ("id")
);

create table if not exists "hist_history" (
  "id" uuid,
  "hist_id" text not null,
  "o" jsonb not null,
  "n" jsonb not null,
  "c" jsonb not null,
  "created" timestamp not null default now(),
  foreign key ("hist_id") references "hist" ("id"),
  primary key ("id")
);
-- `)
//line queries/ddl/hist.sql:26
}

//line queries/ddl/hist.sql:26
func WriteHistCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/hist.sql:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/hist.sql:26
	StreamHistCreate(qw422016)
//line queries/ddl/hist.sql:26
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/hist.sql:26
}

//line queries/ddl/hist.sql:26
func HistCreate() string {
//line queries/ddl/hist.sql:26
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/hist.sql:26
	WriteHistCreate(qb422016)
//line queries/ddl/hist.sql:26
	qs422016 := string(qb422016.B)
//line queries/ddl/hist.sql:26
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/hist.sql:26
	return qs422016
//line queries/ddl/hist.sql:26
}