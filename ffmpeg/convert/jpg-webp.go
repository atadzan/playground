package image

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ConvertJpgToWebp() {
	// "ffmpeg -i in/%8d.png -c:v libwebp out/%8d.webp"
	inputFilePath := "../assets/test-512.jpg"
	outputFilePath := "../assets/test-512.webp"
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		inputFilePath,
		"-c:v",
		"libwebp",
		outputFilePath,
	)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	_ = cmd.Run()
	fmt.Println("Output: ", string(outb.Bytes()), ". Error: ", string(errb.Bytes()))

}
