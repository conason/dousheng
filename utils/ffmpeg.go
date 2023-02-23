package utils

import (
	"bytes"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

// ParseCover 视频封面截取
func ParseCover(videoURL string, frameNum int) ([]byte, error) {
	// Returns specified frame as []byte
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoURL).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).Run()
	if err != nil {
		log.Panicln(err)
		//return nil, err
	}
	byte := buf.Bytes()
	return byte, nil
}
