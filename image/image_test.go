package image_test

import (
	"github.com/KiritoNya/gotaku/image"
	"net/http"
	"testing"
)

var urlImage string = "https://cdn.mangaworld.in/mangas/609e903ee781cc10f5cee3fa.png?1639758876984"

func TestImage_GetName(t *testing.T) {
	var img image.Image
	img.Url = urlImage

	err := img.GetName()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Name:", img.Name)
}

func TestImage_GetFormat(t *testing.T) {
	var img image.Image
	img.Url = urlImage

	err := img.GetFormat()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Format:", img.Format)
}

func TestImage_GetData(t *testing.T) {
	var img image.Image
	img.Url = urlImage

	err := img.GetData(http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Data:", img.Data)
}

func TestImage_Save(t *testing.T) {
	var img image.Image
	img.Url = urlImage

	err := img.GetData(http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}

	err = img.Save("./")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("[OK]")
}
