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

func EncodeToHEVCBest() error {
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		"/home/belet/test/new/Oppenheimer-Trailer.mp4",
		"-c:v",
		"libvpx-vp9",
		"-deadline",
		"best",
		"vp9-best.webm",
	)
	// Cmd 2
	//cmd := exec.Command(
	//	"ffmpeg",
	//	"-i",
	//	"/home/belet/test/new/Oppenheimer-Trailer.mp4",
	//	"-c:v",
	//	"libvpx-vp9",
	//	"-deadline",
	//	"good",
	//	"vp9-good.webm",
	//)

	// Cmd 2
	//cmd := exec.Command(
	//	"ffmpeg",
	//	"-i",
	//	"/home/belet/test/new/Oppenheimer-Trailer.mp4",
	//	"-c:v",
	//	"libvpx-vp9",
	//	"vp9-default.webm",
	//)

	// Adding response body to buffer of executing command

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func EncodeToHEVCGood() error {
	fmt.Println("Deadline Good started")
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		"/home/belet/test/new/Oppenheimer-Trailer.mp4",
		"-c:v",
		"libvpx-vp9",
		"-deadline",
		"good",
		"vp9-good.webm",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
func EncodeToHEVCDefault() error {
	fmt.Println("Started default")

	cmd := exec.Command(
		"ffmpeg",
		"-i",
		"/home/belet/test/new/Oppenheimer-Trailer.mp4",
		"-c:v",
		"libvpx-vp9",
		"vp9-default.webm",
	)

	// Adding response body to buffer of executing command

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
