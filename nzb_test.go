package nzb

import (
	"testing"
)

const testNzbXml = `<?xml version="1.0" encoding="utf-8" ?>
<!DOCTYPE nzb PUBLIC "-//newzBin//DTD NZB 1.1//EN" "http://www.newzbin.com/DTD/nzb/nzb-1.1.dtd">
<nzb xmlns="http://www.newzbin.com/DTD/2003/nzb">
 <head>
   <meta type="title">Your File!</meta>
   <meta type="tag">Example</meta>
 </head>
 <file poster="Joe Bloggs &lt;bloggs@nowhere.example&gt;" date="1071674882" subject="abc-mr2a.r01 (1/2)">
   <groups>
     <group>alt.binaries.newzbin</group>
     <group>alt.binaries.mojo</group>
   </groups>
   <segments>
     <segment bytes="102394" number="1">123456789abcdef@news.newzbin.com</segment>
     <segment bytes="4501" number="2">987654321fedbca@news.newzbin.com</segment>
   </segments>
 </file>
</nzb>`

func TestNzbParse(t *testing.T) {
	nzb, err := NewString(testNzbXml)
	if err != nil {
		t.Fatalf("expected to parse nzb data: " + err.String())
	}
	if nzb.Meta["title"] != "Your File!" {
		t.Errorf("expected name data to be 'Your File' got %s", nzb.Meta["title"])
	}
	if len(nzb.Files) != 1 {
		t.Fatalf("expected 1 file got %d", len(nzb.Files))
	}
	f := nzb.Files[0]
	if f.Subject != "abc-mr2a.r01 (1/2)" {
		t.Errorf("expected file subject to be abc-mr2a.r01 (1/2) but got %v", f.Subject)
	}
	if len(f.Groups) != 2 {
		t.Fatalf("expected there to be 2 groups got %d", len(f.Groups))
	}
	if f.Groups[0] != "alt.binaries.newzbin" {
		t.Errorf("expected group[0] to be alt.binaries.newzbin")
	}
	if len(f.Segments) != 2 {
		t.Fatalf("expected there to be 2 segments got %d", len(f.Segments))
	}
	if f.Segments[0].Bytes != 102394 {
		t.Errorf("expected segment[0].Bytes to be 102394 but got %d", f.Segments[0].Bytes)
	}
	if f.Segments[0].Number != 1 {
		t.Errorf("expected segment[0].Bytes to be 1 but got %d", f.Segments[0].Number)
	}
	if f.Segments[0].Id != "123456789abcdef@news.newzbin.com" {
		t.Errorf("expected message id to be 123456789abcdef@news.newzbin.com but got %s", f.Segments[0].Id)
	}
}
