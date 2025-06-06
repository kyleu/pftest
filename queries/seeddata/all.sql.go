// Code generated by qtc from "all.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/all.sql:1
package seeddata

//line queries/seeddata/all.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/all.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/all.sql:1
func StreamSeedDataAll(qw422016 *qt422016.Writer) {
//line queries/seeddata/all.sql:1
	qw422016.N().S(`
-- `)
//line queries/seeddata/all.sql:2
	StreamOddPKSeedData(qw422016)
//line queries/seeddata/all.sql:2
	qw422016.N().S(`
-- `)
//line queries/seeddata/all.sql:3
	StreamOddrelSeedData(qw422016)
//line queries/seeddata/all.sql:3
	qw422016.N().S(`
-- `)
//line queries/seeddata/all.sql:4
	StreamSeedSeedData(qw422016)
//line queries/seeddata/all.sql:4
	qw422016.N().S(`
-- `)
//line queries/seeddata/all.sql:5
}

//line queries/seeddata/all.sql:5
func WriteSeedDataAll(qq422016 qtio422016.Writer) {
//line queries/seeddata/all.sql:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/all.sql:5
	StreamSeedDataAll(qw422016)
//line queries/seeddata/all.sql:5
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/all.sql:5
}

//line queries/seeddata/all.sql:5
func SeedDataAll() string {
//line queries/seeddata/all.sql:5
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/all.sql:5
	WriteSeedDataAll(qb422016)
//line queries/seeddata/all.sql:5
	qs422016 := string(qb422016.B)
//line queries/seeddata/all.sql:5
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/all.sql:5
	return qs422016
//line queries/seeddata/all.sql:5
}
