package log

import (
	"context"
	"io"
	"path/filepath"
	"time"

	"github.com/zongshuai/kratos/pkg/log/internal/filewriter"
)

// level idx
const (
	_infoIdx = iota
	_warnIdx
	_errorIdx
	_totalIdx
)

var _fileNames = map[int]string{
	_infoIdx:  "info.log",
	_warnIdx:  "warning.log",
	_errorIdx: "error.log",
}

// FileHandler .
type FileHandler struct {
	render Render
	fws    [_totalIdx]*filewriter.FileWriter
	fw     *filewriter.FileWriter
}

// NewFile crete a file logger.
func NewFile(dir, project string, bufferSize, rotateSize int64, maxLogFile int) *FileHandler {
	// new info writer
	newWriter := func(name string) *filewriter.FileWriter {
		var options []filewriter.Option
		if rotateSize > 0 {
			options = append(options, filewriter.MaxSize(rotateSize))
		}
		if maxLogFile > 0 {
			options = append(options, filewriter.MaxFile(maxLogFile))
		}
		w, err := filewriter.New(filepath.Join(dir, name), options...)
		if err != nil {
			panic(err)
		}
		return w
	}

	format := "%J"
	if project == "" {
		format = "[%D %T] [%L] [%i] [%S] %M"
	}

	handler := &FileHandler{
		render: newPatternRender(format),
	}

	if project != "" {
		handler.fw = newWriter(project + ".log")
	} else {
		for idx, name := range _fileNames {
			handler.fws[idx] = newWriter(name)
		}
	}
	return handler
}

// Log loggint to file .
func (h *FileHandler) Log(ctx context.Context, lv Level, args ...D) {
	d := toMap(args...)
	// add extra fields
	addExtraField(ctx, d)
	d[_time] = time.Now().Format(_timeFormat)
	var w io.Writer

	if c.Project == "" {
		switch lv {
		case _warnLevel:
			w = h.fws[_warnIdx]
		case _errorLevel:
			w = h.fws[_errorIdx]
		default:
			w = h.fws[_infoIdx]
		}
	} else {
		w = h.fw
		d = format(d)
	}

	h.render.Render(w, d)
	w.Write([]byte("\n"))
}

// Close log handler
func (h *FileHandler) Close() error {
	if c.Project == "" {
		for _, fw := range h.fws {
			// ignored error
			_ = fw.Close()
		}
	} else {
		_ = h.fw.Close()
	}

	return nil
}

// SetFormat set log format
func (h *FileHandler) SetFormat(format string) {
	h.render = newPatternRender(format)
}

func format(oldFormat map[string]interface{}) map[string]interface{} {
	newFormat := map[string]interface{}{}
	newFormat["trace_id"] = ""

	content := map[string]interface{}{}
	content["message"] = ""
	for k, v := range oldFormat {
		switch k {
		case _level:
			newFormat["level"] = v
		case _time:
			newFormat["time"] = v
		case _tid:
			newFormat["trace_id"] = v
		case _log:
			content["message"] = v
		case _appID:
		case _levelValue:
		case _zone:
		case _deplyEnv:
		case _instanceID:
			//discard
			continue
		default:
			content[k] = v
		}
	}

	newFormat["type"] = "golang"
	newFormat["project"] = c.Project
	newFormat["content"] = content

	return newFormat
}
