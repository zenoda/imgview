package imgview

import (
	"bytes"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"image"
	"image/png"
	"log"
	"os"
)

func Show(img *image.Image) {
	gtk.Init(&os.Args)

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	window.SetTitle("Image View")
	window.SetResizable(false)
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})
	pbl, err := gdk.PixbufLoaderNew()
	if err != nil {
		panic(err)
	}
	var buff bytes.Buffer
	err = png.Encode(&buff, *img)
	if err != nil {
		panic(err)
	}
	pixBuff, err := pbl.WriteAndReturnPixbuf(buff.Bytes())
	if err != nil {
		panic(err)
	}
	image, err := gtk.ImageNewFromPixbuf(pixBuff)
	if err != nil {
		panic(err)
	}

	window.Add(image)

	window.ShowAll()

	gtk.Main()

}
