PACKAGE

package nzb
import "."


TYPES

type Nzb struct {
    Meta  map[string]string
    Files []NzbFile
}

func New(buf io.Reader) (*Nzb, os.Error)

func NewString(data string) (*Nzb, os.Error)

type NzbFile struct {
    Groups   []string     `xml:"groups>group"`
    Segments []NzbSegment `xml:"segments>segment"`
    Poster   string       `xml:"attr"`
    Date     int          `xml:"attr"`
    Subject  string       `xml:"attr"`
    Part     int
}

type NzbSegment struct {
    XMLName xml.Name `xml:"segment"`
    Bytes   int      `xml:"attr"`
    Number  int      `xml:"attr"`
    Id      string   `xml:"innerxml"`
}


