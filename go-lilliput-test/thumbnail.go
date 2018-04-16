package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/discordapp/lilliput"
)

func main() {
	var EncodeOptions = map[string]map[int]int{
		"jpeg": map[int]int{lilliput.JpegQuality: 100},
	}

	inputFilename := "../images/earth.jpg"
	outputFilename := "thumbs/earth.jpg"
	outputImage := make([]byte, 2*1024*1024)

	inputBuf, _ := ioutil.ReadFile(inputFilename)

	decoder, _ := lilliput.NewDecoder(inputBuf)
	defer decoder.Close()

	outputType := "." + strings.ToLower(decoder.Description())

	ops := lilliput.NewImageOps(8192)
	defer ops.Close()

	opts := &lilliput.ImageOptions{
		FileType:      outputType,
		Width:         128,
		Height:        128,
		ResizeMethod:  lilliput.ImageOpsFit,
		EncodeOptions: EncodeOptions[outputType],
	}

	outputImage, _ = ops.Transform(decoder, opts, outputImage)
	os.Mkdir("thumbs", 0755)
	os.Remove(outputFilename)
	ioutil.WriteFile(outputFilename, outputImage, 0400)
}
