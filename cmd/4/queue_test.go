package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}
}

func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("","tmp")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	return file
}

func BenchmarkBufferdWrite(b *testing.B) {
	// buf := bytes.Buffer{}
	bufferedFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferedFile))
}

func BenchmarkUnBufferdWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}