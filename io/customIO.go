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

func MakeIO(in io.Reader, out io.Writer, bufsize uint) (*CustomIO, error) {
	if bufsize > 1e6 {
		return nil, fmt.Errorf("bufsize is too large")
	}
	io := &CustomIO{
		reader: *bufio.NewReaderSize(in, int(bufsize)),
		writer: *bufio.NewWriterSize(out, int(bufsize)),
		cache:  []string{},
	}
	return io, nil
}

func (io *CustomIO) NextLine() ([]string, error) {
	text, err := io.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	arr := strings.Split(strings.TrimRight(text, "\n"), " ")
	return arr, nil
}

func (io *CustomIO) NextWord() (string, error) {
	for len(io.cache) == 0 {
		text, err := io.reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		arr := strings.Split(strings.TrimRight(text, "\n"), " ")
		for i := len(arr) - 1; i >= 0; i-- {
			io.cache = append(io.cache, arr[i])
		}
	}
	res := io.cache[len(io.cache)-1]
	io.cache = io.cache[:len(io.cache)-1]
	return res, nil
}

func (io *CustomIO) NextInt() (int, error) {
	s, err := io.NextWord()
	if err != nil {
		return 0, err
	}
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (io *CustomIO) NextFloat(sep byte) (float64, error) {
	s, err := io.NextWord()
	if err != nil {
		return 0.0, err
	}
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return res, nil
}

func (io *CustomIO) Println(objects ...interface{}) (int, error) {
	n, err := fmt.Fprintln(&io.writer, objects...)
	return n, err
}

func (io *CustomIO) Printf(format string, objects ...interface{}) (int, error) {
	n, err := fmt.Fprintf(&io.writer, format, objects...)
	return n, err
}

func (io *CustomIO) Flush() error {
	err := io.writer.Flush()
	return err
}
