package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	traceDepth = 128
	skipFrames = 3
)

type stacktrace struct {
	depth  int
	frames *runtime.Frames
}

type Stackframe struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`
}

func newStacktrace() *stacktrace {
	var trace = make([]uintptr, traceDepth)

	depth := runtime.Callers(skipFrames, trace[:])
	trace = trace[:depth]

	return &stacktrace{
		frames: runtime.CallersFrames(trace),
		depth:  depth,
	}
}

func (s *stacktrace) ToString() string {
	result := strings.Builder{}

	for {
		frame, more := s.frames.Next()
		line := fmt.Sprintf("%s\n\t%s (%d)", frame.Function, frame.File, frame.Line)
		result.WriteString(line)

		if !more {
			break
		}

		result.WriteString("\n")
	}

	return result.String()
}

func (s *stacktrace) ToJson() []Stackframe {
	result := make([]Stackframe, s.depth)

	i := 0
	for {
		frame, more := s.frames.Next()
		result[i] = Stackframe{
			File:     frame.File,
			Line:     frame.Line,
			Function: frame.Function,
		}
		i++

		if !more {
			break
		}
	}

	return result
}
