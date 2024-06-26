// Code generated by qtc from "all.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/all.sql:1
package ddl

//line queries/ddl/all.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/all.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/all.sql:1
func StreamDropAll(qw422016 *qt422016.Writer) {
//line queries/ddl/all.sql:1
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:2
	StreamTroubleDrop(qw422016)
//line queries/ddl/all.sql:2
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:3
	StreamTimestampDrop(qw422016)
//line queries/ddl/all.sql:3
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:4
	StreamSoftdelDrop(qw422016)
//line queries/ddl/all.sql:4
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:5
	StreamSeedDrop(qw422016)
//line queries/ddl/all.sql:5
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:6
	StreamRelationDrop(qw422016)
//line queries/ddl/all.sql:6
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:7
	StreamReferenceDrop(qw422016)
//line queries/ddl/all.sql:7
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:8
	StreamPathDrop(qw422016)
//line queries/ddl/all.sql:8
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:9
	StreamMixedCaseDrop(qw422016)
//line queries/ddl/all.sql:9
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:10
	StreamBasicDrop(qw422016)
//line queries/ddl/all.sql:10
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:11
	StreamAuditedDrop(qw422016)
//line queries/ddl/all.sql:11
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:12
	StreamCapitalDrop(qw422016)
//line queries/ddl/all.sql:12
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:13
	StreamTypesDrop(qw422016)
//line queries/ddl/all.sql:13
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:14
}

//line queries/ddl/all.sql:14
func WriteDropAll(qq422016 qtio422016.Writer) {
//line queries/ddl/all.sql:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/all.sql:14
	StreamDropAll(qw422016)
//line queries/ddl/all.sql:14
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/all.sql:14
}

//line queries/ddl/all.sql:14
func DropAll() string {
//line queries/ddl/all.sql:14
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/all.sql:14
	WriteDropAll(qb422016)
//line queries/ddl/all.sql:14
	qs422016 := string(qb422016.B)
//line queries/ddl/all.sql:14
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/all.sql:14
	return qs422016
//line queries/ddl/all.sql:14
}

// --

//line queries/ddl/all.sql:16
func StreamCreateAll(qw422016 *qt422016.Writer) {
//line queries/ddl/all.sql:16
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:17
	StreamTypesCreate(qw422016)
//line queries/ddl/all.sql:17
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:18
	StreamCapitalCreate(qw422016)
//line queries/ddl/all.sql:18
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:19
	StreamAuditedCreate(qw422016)
//line queries/ddl/all.sql:19
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:20
	StreamBasicCreate(qw422016)
//line queries/ddl/all.sql:20
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:21
	StreamMixedCaseCreate(qw422016)
//line queries/ddl/all.sql:21
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:22
	StreamPathCreate(qw422016)
//line queries/ddl/all.sql:22
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:23
	StreamReferenceCreate(qw422016)
//line queries/ddl/all.sql:23
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:24
	StreamRelationCreate(qw422016)
//line queries/ddl/all.sql:24
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:25
	StreamSeedCreate(qw422016)
//line queries/ddl/all.sql:25
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:26
	StreamSoftdelCreate(qw422016)
//line queries/ddl/all.sql:26
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:27
	StreamTimestampCreate(qw422016)
//line queries/ddl/all.sql:27
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:28
	StreamTroubleCreate(qw422016)
//line queries/ddl/all.sql:28
	qw422016.N().S(`
-- `)
//line queries/ddl/all.sql:29
}

//line queries/ddl/all.sql:29
func WriteCreateAll(qq422016 qtio422016.Writer) {
//line queries/ddl/all.sql:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/all.sql:29
	StreamCreateAll(qw422016)
//line queries/ddl/all.sql:29
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/all.sql:29
}

//line queries/ddl/all.sql:29
func CreateAll() string {
//line queries/ddl/all.sql:29
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/all.sql:29
	WriteCreateAll(qb422016)
//line queries/ddl/all.sql:29
	qs422016 := string(qb422016.B)
//line queries/ddl/all.sql:29
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/all.sql:29
	return qs422016
//line queries/ddl/all.sql:29
}
