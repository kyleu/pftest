// Code generated by qtc from "softdel.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/softdel.sql:1
package ddl

//line queries/ddl/softdel.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/softdel.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/softdel.sql:1
func StreamSoftdelDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/softdel.sql:1
	qw422016.N().S(`
drop table if exists "softdel";
-- `)
//line queries/ddl/softdel.sql:3
}

//line queries/ddl/softdel.sql:3
func WriteSoftdelDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/softdel.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/softdel.sql:3
	StreamSoftdelDrop(qw422016)
//line queries/ddl/softdel.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/softdel.sql:3
}

//line queries/ddl/softdel.sql:3
func SoftdelDrop() string {
//line queries/ddl/softdel.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/softdel.sql:3
	WriteSoftdelDrop(qb422016)
//line queries/ddl/softdel.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/softdel.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/softdel.sql:3
	return qs422016
//line queries/ddl/softdel.sql:3
}

// --

//line queries/ddl/softdel.sql:5
func StreamSoftdelCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/softdel.sql:5
	qw422016.N().S(`
create table if not exists "softdel" (
  "id" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  "deleted" timestamp default now(),
  primary key ("id")
);
-- `)
//line queries/ddl/softdel.sql:13
}

//line queries/ddl/softdel.sql:13
func WriteSoftdelCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/softdel.sql:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/softdel.sql:13
	StreamSoftdelCreate(qw422016)
//line queries/ddl/softdel.sql:13
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/softdel.sql:13
}

//line queries/ddl/softdel.sql:13
func SoftdelCreate() string {
//line queries/ddl/softdel.sql:13
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/softdel.sql:13
	WriteSoftdelCreate(qb422016)
//line queries/ddl/softdel.sql:13
	qs422016 := string(qb422016.B)
//line queries/ddl/softdel.sql:13
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/softdel.sql:13
	return qs422016
//line queries/ddl/softdel.sql:13
}
