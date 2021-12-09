package main

import (
	"io"
	"os"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

func main() {
	var files []io.Writer
	fileName := "myLog.log"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)

	if err != nil {
		hclog.Default().Error("File create failed", "error", err.Error())
	}
	files = append(files, f, os.Stdout)
	mw := io.MultiWriter(files...)
	hclog.Default().Info("hello world")
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:            "my-app",
		Level:           hclog.LevelFromString("DEBUG"),
		DisableTime:     false,
		JSONFormat:      true,
		IncludeLocation: true,
		Output:          mw,
		TimeFormat:      "02 Jan 06 15:04 MST",
	})

	input := "5.5"
	_, errr := strconv.ParseInt(input, 10, 32)
	if errr != nil {
		appLogger.Error("Invalid input for ParseInt", "input", input, "error", errr.Error(), "msg", "hello")
		appLogger.Info("This is a binarry input", "bits", hclog.Binary(17))
		// appLogger.Trace("End of your file")
	}

}
