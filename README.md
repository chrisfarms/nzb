nzb.go
======

Little helper package for reading [NZB files](http://en.wikipedia.org/wiki/NZB) in Go (golang)

Install
-------

```
go get github.com/chrisfarms/nzb
```


Godoc
-----

```go
type Nzb struct {
    Meta  map[string]string
    Files []*NzbFile
}
```


```go
func New(buf []byte) (*Nzb, error)
```

```go
type NzbFile struct {
    Groups   []string     `xml:"groups>group"`
    Segments []NzbSegment `xml:"segments>segment"`
    Poster   string       `xml:"poster,attr"`
    Date     int          `xml:"date,attr"`
    Subject  string       `xml:"subject,attr"`
    Part     int
}
```

```go
type NzbSegment struct {
    XMLName xml.Name `xml:"segment"`
    Bytes   int      `xml:"bytes,attr"`
    Number  int      `xml:"number,attr"`
    Id      string   `xml:",innerxml"`
}
```


