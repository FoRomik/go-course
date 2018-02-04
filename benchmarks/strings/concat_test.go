package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Record struct to manage stats messages
type Record struct {
	EmailID      int64
	EmailRunTime int64
	ContentID    int64
	DomainGroup  string
	LeadStatus   int
	ProfileID    int64
	SiteID       int64
	TemplateID   int64
	Timestamp    int64
}

var (
	sep = "_"
	rec = Record{
		EmailID:      1234567890,
		EmailRunTime: 1234567890,
		ContentID:    1234567890,
		DomainGroup:  "Google",
		LeadStatus:   1234567890,
		ProfileID:    1234567890,
		SiteID:       1234567890,
		TemplateID:   1234567890,
		Timestamp:    1234567890,
	}
)

// Benchmarks string concatenation with fmt package
func BenchmarkConcatFmt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("%v%v%v%v%v%v%v%v%v",
			"prefix",
			sep,
			rec.Timestamp,
			sep,
			rec.DomainGroup,
			sep,
			rec.SiteID,
			sep,
			rec.EmailID)
		_ = fmt.Sprintf("r_%v", key)
	}
}

// Benchmarks string concatenation with bytes package
func BenchmarkConcatBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		buf.WriteString("prefix")
		buf.WriteString(sep)
		buf.WriteString(strconv.Itoa(int(rec.Timestamp)))
		buf.WriteString(sep)
		buf.WriteString(rec.DomainGroup)
		buf.WriteString(sep)
		buf.WriteString(strconv.FormatInt(rec.SiteID, 10))
		buf.WriteString(sep)
		buf.WriteString(strconv.FormatInt(rec.EmailID, 10))

		_ = buf.String()
		buf.Reset()

		buf.WriteString("r")
		buf.WriteString(sep)
		buf.WriteString("prefix")
		buf.WriteString(sep)
		buf.WriteString(strconv.Itoa(int(rec.Timestamp)))
		buf.WriteString(sep)
		buf.WriteString(rec.DomainGroup)
		buf.WriteString(sep)
		buf.WriteString(strconv.FormatInt(rec.SiteID, 10))
		buf.WriteString(sep)
		buf.WriteString(strconv.FormatInt(rec.EmailID, 10))

		_ = buf.String()
	}
}

// Benchmarks string concatenation with strings package
func BenchmarkConcatStrings(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := []string{
			"r",
			"prefix",
			strconv.Itoa(int(rec.Timestamp)),
			rec.DomainGroup,
			strconv.FormatInt(rec.SiteID, 10),
			strconv.FormatInt(rec.EmailID, 10),
		}

		_ = strings.Join(buf, sep)
		_ = strings.Join(buf[1:], sep)
	}
}

/* 2018-01-12 go 1.9.1 linux/amd64

BenchmarkConcatFmt-8             1000000              1308 ns/op             216 B/op         11 allocs/op
BenchmarkConcatBytes-8           2000000               993 ns/op             304 B/op          9 allocs/op
BenchmarkConcatStrings-8         2000000               594 ns/op             240 B/op          7 allocs/op

*/
