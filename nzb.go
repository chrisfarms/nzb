package nzb

import(
    "xml"
    "os"
    "bytes"
    "io"
)

type Nzb struct {
    Meta map[string]string
    Files []NzbFile
}

func NewString(data string) (*Nzb, os.Error) {
    return New(bytes.NewBufferString(data))
}

func New(buf io.Reader) (*Nzb, os.Error) {
    xnzb := new(xNzb)
    err := xml.Unmarshal(buf, xnzb)
    if err != nil {
        return nil,err
    }
    // convert to nicer format
    nzb := new(Nzb)
    // convert metadata
    nzb.Meta = make(map[string]string)
    for _,md := range xnzb.Metadata {
        nzb.Meta[md.Type] = md.Value
    }
    // copy files
    nzb.Files = xnzb.File
    return nzb,nil
}

// used only for unmarshalling xml
type xNzb struct {
    XMLName xml.Name `xml:"nzb"`
    Metadata []xNzbMeta `xml:"head>meta"`
    File []NzbFile `xml:"file"` // xml:tag name doesn't work?
}

// used only in unmarshalling xml
type xNzbMeta struct {
    Type string `xml:"attr"`
    Value string `xml:"innerxml"`
}

type NzbFile struct {
    Groups []string `xml:"groups>group"`
    Segments []NzbSegment `xml:"segments>segment"`
    Poster string `xml:"attr"`
    Date int `xml:"attr"`
    Subject string `xml:"attr"`
    Part int
}

type NzbSegment struct {
    XMLName xml.Name `xml:"segment"`
    Bytes int `xml:"attr"`
    Number int `xml:"attr"`
    Id string `xml:"innerxml"`
}

