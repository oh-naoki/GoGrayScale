package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	imgUrl := sc.Text()
	convertGrayScale(getImageBytes(imgUrl))
}

func getImageBytes(url string) []byte {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}

	response, err := client.Do(request)
	if err != nil {

	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {

	}

	return body
}

func convertGrayScale(imgBytes []byte) {
	img, err := jpeg.Decode(bytes.NewBuffer(imgBytes))
	if err != nil {
		fmt.Println("image must be jpeg")
	}

	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dest.Set(x, y, gray)
		}
	}

	jpeg.Encode(os.Stdout, dest, nil)
}
