package io

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	errors "github.com/hiden2000/go_ds/errors"
)

const MaxBufSize = 1 << 20

// バッファリング機能付きIO構造体
type CustomIO struct {
	reader bufio.Reader
	writer bufio.Writer
	cache  []string
	sep    byte
}

// New は in を入力先とし out を出力先とする バッファーサイズ= bufSize , 区切り文字= sepの CustomIO を初期化する
// bufSize が MaxBufSize より大きい場合 ErrInvalidValue を返す
func New(in io.Reader, out io.Writer, bufSize uint, sep byte) (*CustomIO, error) {
	if bufSize > MaxBufSize {
		return nil, fmt.Errorf("%w: bufsize is too large", errors.ErrInvalidValue)
	}
	io := &CustomIO{
		reader: *bufio.NewReaderSize(in, int(bufSize)),
		writer: *bufio.NewWriterSize(out, int(bufSize)),
		cache:  []string{},
		sep:    sep,
	}
	return io, nil
}

// NextLine は　次の改行文字までの入力を読み込む
func (io *CustomIO) NextLine() ([]string, error) {
	text, err := io.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	arr := strings.Split(strings.TrimRight(text, "\n"), string(io.sep))
	return arr, nil
}

// NextWord は　次の入力単位を読み込む
func (io *CustomIO) NextWord() (string, error) {
	for len(io.cache) == 0 {
		text, err := io.reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		arr := strings.Split(strings.TrimRight(text, "\n"), string(io.sep))
		for i := len(arr) - 1; i >= 0; i-- {
			io.cache = append(io.cache, arr[i])
		}
	}
	res := io.cache[len(io.cache)-1]
	io.cache = io.cache[:len(io.cache)-1]
	return res, nil
}

// NextInt は　次の入力単位を整数として読み込もうとし，失敗した場合はerrorを返す
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

// NextInt は　次の入力単位を少数として読み込もうとし，失敗した場合はerrorを返す
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

// Println は 出力先に fmt.Println 機能を提供する
func (io *CustomIO) Println(objects ...interface{}) (int, error) {
	n, err := fmt.Fprintln(&io.writer, objects...)
	return n, err
}

// Printf は 出力先に fmt.Printf 機能を提供する
func (io *CustomIO) Printf(format string, objects ...interface{}) (int, error) {
	n, err := fmt.Fprintf(&io.writer, format, objects...)
	return n, err
}

// Flush は 出力先に fmt.Flush 機能を提供する
func (io *CustomIO) Flush() error {
	err := io.writer.Flush()
	return err
}
