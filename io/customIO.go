package io

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CustomIO struct {
	reader bufio.Reader
	writer bufio.Writer
	cache  []string
}

func MakeIO(in io.Reader, out io.Writer, bufsize int) (*CustomIO, error) {
	if bufsize > 1e6 {
		return nil, fmt.Errorf("bufsize is too large")
	}
	io := &CustomIO{
		reader: *bufio.NewReaderSize(in, bufsize),
		writer: *bufio.NewWriterSize(out, bufsize),
		cache:  []string{},
	}
	return io, nil
}

func (io *CustomIO) NextLine() string {
	for len(io.cache) == 0 {
		text, err := io.reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		arr := strings.Split(strings.TrimRight(text, "\n"), " ")
		for i := len(arr) - 1; i >= 0; i-- {
			io.cache = append(io.cache, arr[i])
		}
	}
	res := io.cache[len(io.cache)-1]
	io.cache = io.cache[:len(io.cache)-1]
	return res
}

func (io *CustomIO) NextInt() int {
	res, err := strconv.Atoi(io.NextLine())
	if err != nil {
		panic(err)
	}
	return res
}

func (io *CustomIO) NextFloat(sep byte) float64 {
	res, err := strconv.ParseFloat(io.NextLine(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (io *CustomIO) Print(objects ...interface{}) {
	_, err := fmt.Fprintln(&io.writer, objects...)
	if err != nil {
		panic(err)
	}
}

func (io *CustomIO) Printf(format string, objects ...interface{}) {
	_, err := fmt.Fprintf(&io.writer, format, objects...)
	if err != nil {
		panic(err)
	}
}

func (io *CustomIO) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}
