# charenc

[![GitHub Language](https://img.shields.io/badge/Go-reference-blue)](https://go.dev)
[![GitHub license](https://img.shields.io/github/license/ChenYuTong10/charenc)](https://github.com/ChenYuTong10/charenc/blob/master/LICENSE)

A simple character encoder implemented by Go.

The encoder transforms text encoding from *ANSI, UTF8, BOM UTF8* and *BOM UTF16 BE/LE* to specific encoding which supports *ANSI* and *UTF8* now.

## Install

```bash
go get github.com/ChenYuTong10/charenc
```

## Example

Encode other encodings to `Ansi`.

```Golang
import (
    "log"

    "github.com/ChenYuTong10/charenc"
)

func Foo() {
    stream, err := os.ReadFile("utf8.txt")
        if err != nil {
        log.Printf("read file error: %v", err)
        return
    }

    stream, err = charenc.ToAnsi(stream, "UTF8")
    if err != nil {
        log.Printf("ansi encode error: %v", err)
        return
    }

    // do anything you want
}
```

Encode other encodings to `UTF8`.

```Golang
import (
    "log"

    "github.com/ChenYuTong10/charenc"
)

func Foo() {
    stream, err := os.ReadFile("utf16BE.txt")
        if err != nil {
        log.Printf("read file error: %v", err)
        return
    }

    stream, err = charenc.ToUTF8(stream, "UTF-16 BE")
    if err != nil {
        log.Printf("ansi encode error: %v", err)
        return
    }

    // do anything you want
}
```

Usually, you may detect the encoding of a text and transform it to other encodings. In this case, you can use `github.com/ChenYuTong10/chardet` package to work together.

```Golang
import (
    "log"
    "os"

    "github.com/ChenYuTong10/chardet"
    "github.com/ChenYuTong10/charenc"
)

func Foo() {
    stream, err := os.ReadFile("example.txt")
        if err != nil {
        log.Printf("read file error: %v", err)
        return
    }

    d := new(Detector)
    d.Feed(stream)
    
    encoding := d.Encoding

    // transform encoding to ANSI
    stream, err = charenc.ToAnsi(stream, encoding)
    if err != nil {
        log.Printf("ansi encode error: %v", err)
        return
    }

    // do anything you want
}
```