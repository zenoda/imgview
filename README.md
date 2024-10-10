# imgview

## 1. Installation Instructions
### Go
Go 1.8 or newer is required.

### GTK, GLib, Cairo
Due to imgview's dependency on gotk3, according to the installation instructions for gotk3, it currently requires GTK 3.6-3.22, GLib 2.36-2.40, and Cairo 1.10 or 1.12:

- [Linux](https://github.com/gotk3/gotk3/wiki/Installing-on-Linux)
- [macOS](https://github.com/gotk3/gotk3/wiki/Installing-on-macOS)
- [Windows](https://github.com/gotk3/gotk3/wiki/Installing-on-Windows)

### imgview

```
go get -u github.com/zenoda/imgview
```

## 2. Example Usage
```go
package main
import (
    "github.com/zenoda/imgview"
    "image/jpeg"
    "os"
)

func main() {
    f, err := os.Open("pic.jpg")
    if err != nil {
    panic(err)
    }
    defer f.Close()
    img, err := jpeg.Decode(f)
    if err != nil {
    panic(err)
    }
    imgview.Show(&img)
}
```
