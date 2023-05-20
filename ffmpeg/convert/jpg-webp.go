package image

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func ConvertJpgToWebpWithOutput(input string) (io.Reader, error) {
	//  os.Pipe() function to create an in-memory pipe and directly pass the pipe reader to the HTTP request body.
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		input,
		"-c:v",
		"libwebp",
		"-pix_fmt",
		"yuv420p",
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

	//var waitToConvert chan int

	go func() {
		defer pipeWriter.Close()
		err = cmd.Wait()
		if err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("ffmpeg command failed: %v", err))
		}
		//waitToConvert <- 1
	}()
	//<-waitToConvert
	return pipeReader, nil
}

func ConvertJpgToWebpFromResponseBody(body []byte, outputPath string) error {
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		"-",
		"-c:v",
		"libwebp",
		"-pix_fmt",
		"yuv420p",
		"-f",
		"webp",
		outputPath,
	)
	// Adding response body to buffer of executing command
	cmd.Stdin = bytes.NewReader(body)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
