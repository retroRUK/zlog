package zlog

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"runtime"
)

func HttpError(w http.ResponseWriter, msg string, err error, code int) {
	file, line, fn := callerInfo(2)

    slog.Error(msg,
        "error", err,
        "file", file,
        "line", line,
        "func", fn,
    )

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func Error(msg string, err error) {
	file, line, fn := callerInfo(2)

    slog.Error(msg,
        "error", err,
        "file", file,
        "line", line,
        "func", fn,
    )
}

func callerInfo(skip int) (file string, line int, function string) {
    pc, file, line, ok := runtime.Caller(skip)
    if !ok {
        return "???", 0, "unknown"
    }
    fn := runtime.FuncForPC(pc)
    function = fn.Name()
    return file, line, function
}
