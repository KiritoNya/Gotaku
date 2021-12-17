package image

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Image is a struct that describes an image
type Image struct {
	//Id of image
	Id string

	//Url of image
	Url string

	//Name of image
	Name string

	//Data of image
	Data []byte

	// Format of image (.jpg,.png ecc...)
	Format string
}

//GetData is a function that get image data.
func (i *Image) GetData(client *http.Client) error {
	//Check url
	if i.Url == "" {
		return errors.New("Url of image not found")
	}

	req, err := client.Get(i.Url)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	i.Data, err = io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return nil
}

//GetName is a function that get the image name
func (i *Image) GetName() error {
	//Check url
	if i.Url == "" {
		return errors.New("Url of image not found")
	}

	i.Name = filepath.Base(i.Url)

	return nil
}

//GetFormat is a function that get the image format
func (i *Image) GetFormat() error {
	if i.Name == "" {
		err := i.GetName()
		if err != nil {
			return err
		}
	}
	i.Format = filepath.Ext(i.Url)

	if i.Format == "" {
		return errors.New("Image format not found")
	}

	return nil
}

//Save is a function that save the image in a file
func (i *Image) Save(dirPath string) error {
	//Check image data
	if i.Data == nil {
		err := i.GetData(http.DefaultClient)
		if err != nil {
			return err
		}
	}

	//Check image name
	if i.Name == "" {
		err := i.GetName()
		if err != nil {
			return err
		}
	}

	//Create file
	f, err := os.Create(filepath.Join(dirPath, i.Name))
	if err != nil {
		return err
	}
	defer f.Close()

	//Write file
	_, err = f.Write(i.Data)
	if err != nil {
		return err
	}

	return nil
}
