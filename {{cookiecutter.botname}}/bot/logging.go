package bot

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("BotStatus")
var logfile io.Writer

func init() {
	// Initiate the logging folder.
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", 0777)
	}

	logfile, err := os.Create("./logs/" + fmt.Sprint("", int32(time.Now().Unix())) + ".log")
	if err != nil {
		log.Error(err)
	}

	stdBackend := logging.NewLogBackend(os.Stderr, "", 0)
	fileBackend := logging.NewLogBackend(logfile, "", 0)
	logging.SetBackend(stdBackend, fileBackend)
}
