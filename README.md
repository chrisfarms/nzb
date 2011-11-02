nzb.go
======

Little helper package for reading [NZB files](http://en.wikipedia.org/wiki/NZB) in Go (golang)

Install
-------

```
goinstall github.com/chrisfarms/nzb
```


Godoc
-----

```go
type Nzb struct {
    Meta  map[string]string
    Files []NzbFile
}
```


```go
func New(buf io.Reader) (*Nzb, os.Error)
```

```go
func NewString(data string) (*Nzb, os.Error)
```

```go
type NzbFile struct {
    Groups   []string     `xml:"groups>group"`
    Segments []NzbSegment `xml:"segments>segment"`
    Poster   string       `xml:"attr"`
    Date     int          `xml:"attr"`
    Subject  string       `xml:"attr"`
    Part     int
}
```

```go
type NzbSegment struct {
    XMLName xml.Name `xml:"segment"`
    Bytes   int      `xml:"attr"`
    Number  int      `xml:"attr"`
    Id      string   `xml:"innerxml"`
}
```


