// Code generated by qtc from "trouble.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/trouble.sql:2
package ddl

//line queries/ddl/trouble.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/trouble.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/trouble.sql:2
func StreamTroubleDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/trouble.sql:2
	qw422016.N().S(`
drop table if exists "trouble_selectcol";
drop table if exists "trouble";
-- `)
//line queries/ddl/trouble.sql:5
}

//line queries/ddl/trouble.sql:5
func WriteTroubleDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/trouble.sql:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/trouble.sql:5
	StreamTroubleDrop(qw422016)
//line queries/ddl/trouble.sql:5
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/trouble.sql:5
}

//line queries/ddl/trouble.sql:5
func TroubleDrop() string {
//line queries/ddl/trouble.sql:5
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/trouble.sql:5
	WriteTroubleDrop(qb422016)
//line queries/ddl/trouble.sql:5
	qs422016 := string(qb422016.B)
//line queries/ddl/trouble.sql:5
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/trouble.sql:5
	return qs422016
//line queries/ddl/trouble.sql:5
}

// --

//line queries/ddl/trouble.sql:7
func StreamTroubleCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/trouble.sql:7
	qw422016.N().S(`
create table if not exists "trouble" (
  "from" text not null,
  "where" int not null,
  "current_selectcol" int not null default 1,
  "limit" text not null,
  "delete" timestamp default now(),
  primary key ("from", "where")
);

create index if not exists "trouble__from_idx" on "trouble" ("from");

create index if not exists "trouble__where_idx" on "trouble" ("where");

create table if not exists "trouble_selectcol" (
  "trouble_from" text not null,
  "trouble_where" int not null,
  "selectcol" int not null default 1,
  "group" text not null,
  foreign key ("trouble_from", "trouble_where") references "trouble" ("from", "where"),
  primary key ("trouble_from", "trouble_where", "selectcol")
);

create index if not exists "trouble_selectcol__trouble_from_trouble_where_idx" on "trouble_selectcol" ("trouble_from", "trouble_where");

create index if not exists "trouble_selectcol__trouble_from_idx" on "trouble_selectcol" ("trouble_from");

create index if not exists "trouble_selectcol__trouble_where_idx" on "trouble_selectcol" ("trouble_where");
-- `)
//line queries/ddl/trouble.sql:35
}

//line queries/ddl/trouble.sql:35
func WriteTroubleCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/trouble.sql:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/trouble.sql:35
	StreamTroubleCreate(qw422016)
//line queries/ddl/trouble.sql:35
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/trouble.sql:35
}

//line queries/ddl/trouble.sql:35
func TroubleCreate() string {
//line queries/ddl/trouble.sql:35
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/trouble.sql:35
	WriteTroubleCreate(qb422016)
//line queries/ddl/trouble.sql:35
	qs422016 := string(qb422016.B)
//line queries/ddl/trouble.sql:35
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/trouble.sql:35
	return qs422016
//line queries/ddl/trouble.sql:35
}