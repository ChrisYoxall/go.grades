package log

// Since we are creating a log package must alias the stdlib log package below
import (
	"io"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {

	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	return f.Write(data)
}

func Run(dest string) {
	log = stlog.New(fileLog(dest), "", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		msg, err := io.ReadAll(r.Body)
		if err != nil || len(msg) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		write(string(msg))
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
