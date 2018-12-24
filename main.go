package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/valyala/gozstd"
)

func main() {
	imgFile := "./images/beautiful-ultra-hd-wallpapers-for-desktop-4k.jpg"

	imgByte, imgFileExtension, err := readImageFile(imgFile)

	if err != nil {
		panic(err)
	}

	compressedData := gozstd.Compress(nil, imgByte)

	decompressedData, err := gozstd.Decompress(nil, compressedData)
	if err != nil {
		log.Fatalf("cannot decompress data: %s", err)
	}

	decompressedImg, _, _ := image.Decode(bytes.NewReader(decompressedData))

	newImgFile := fmt.Sprintf("compressed-%d.%s", time.Now().Unix(), imgFileExtension)

	//save the imgByte to file
	out, err := os.Create("./output/" + newImgFile)

	if err != nil {
		panic(err)
	}

	err = jpeg.Encode(out, decompressedImg, nil)

	if err != nil {
		panic(err)
	}
	fmt.Println("File successfully written")
}

func readImageFile(imgFile string) ([]byte, string, error) {

	existingImageFile, err := os.Open(imgFile)

	if err != nil {
		// just panic
		return nil, "", err
	}

	defer existingImageFile.Close()

	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		// Handle error
		return nil, "", err
	}

	log.Print(imageType)

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, imageData, nil)

	if err != nil {
		return nil, "", err
	}

	return buf.Bytes(), imageType, err
}
