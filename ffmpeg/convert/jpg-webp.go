package image

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func ConvertJpgToWebp() {
	// "ffmpeg -i in/%8d.png -c:v libwebp out/%8d.webp"
	inputFilePath := "../assets/test-48.jpg"
	outputFilePath := "../assets/test-49.webp"
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

func ConvertJpgToWebpWithOutput(input string) (io.Reader, error) {
	//  os.Pipe() function to create an in-memory pipe and directly pass the pipe reader to the HTTP request body.
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		input,
		"-c:v",
		"libwebp",
		"-f",
		"webp",
		"-",
	)
	pipeReader, pipeWriter := io.Pipe()
	cmd.Stdout = pipeWriter
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg command: %v", err)
	}

	go func() {
		defer pipeWriter.Close()
		err = cmd.Wait()
		if err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("ffmpeg command failed: %v", err))
		}
	}()
	return pipeReader, nil
}
