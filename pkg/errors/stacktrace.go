package errors

import (
	"runtime"
)

const (
	maxStackDepth = 50
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

func newStacktrace(skipFrames int) stacktrace {
	var trace = make([]uintptr, maxStackDepth)

	depth := runtime.Callers(skipFrames, trace[:])
	trace = trace[:depth]

	return stacktrace{
		depth:  depth,
		frames: runtime.CallersFrames(trace),
	}
}

func (s stacktrace) Caller() Stackframe {
	frame, _ := s.frames.Next()
	return Stackframe{
		File:     frame.File,
		Line:     frame.Line,
		Function: frame.Function,
	}
}

func (s stacktrace) Frames() []Stackframe {
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
