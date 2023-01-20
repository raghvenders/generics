
package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"log"
)

type CompressorReader interface {
	ReaderCloser
}

type CompressorWriter interface {
	WriterCloser
}

func WriteCompress[C CompressorWriter](comp C, op string, data string) ([]byte, error) {

	switch op {

	case "gw":
		fmt.Println("Writing Gzip")
	case "zw":
		fmt.Println("Writing Zlib")

	}

	fmt.Println("Data into Compression :: ", data)
	_, err := comp.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}

	if err := comp.Close(); err != nil {
		log.Fatal(err)
	}

	return nil, nil

}

func ReadCompress[C CompressorReader](comp C, op string) ([]byte, error) {

	switch op {
	case "gr":
		fmt.Println("Reading Gzip")
	case "zr":
		fmt.Println("Reading Zlib")

	}

	/*
		n, err := io.Copy(os.Stdout, comp)
		if err != nil {
			fmt.Errorf("error2: [%w]", err)
		}


		fmt.Printf("%d \n", n)
	*/

	rdbuf := make([]byte, 30)

	n, err := comp.Read(rdbuf)
	if err != nil {
		fmt.Errorf("error2: [%w]", err)
	}

	fmt.Printf("Data Length : %d, Content : %v \n", n, string(rdbuf))

	if err := comp.Close(); err != nil {
		log.Fatal(err)
	}

	//v, e := comp.Read(buf.Bytes())
	//if e != nil {
	//fmt.Printf("Value : %d %w", v, e)
	//}

	return nil, nil

}

type ReaderCloser interface {
	io.Reader
	io.Closer
}

type WriterCloser interface {
	io.Writer
	io.Closer
}

func main() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	WriteCompress(zw, "gw", "Generics Gzip Compression\n")

	zr, _ := gzip.NewReader(&buf)

	ReadCompress(zr, "gr")

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	//w.Write([]byte("hello, world\n"))
	//w.Close()

	WriteCompress(w, "zw", "Generics Zlib Compression\n")

	r, _ := zlib.NewReader(&b)
	ReadCompress(r, "zr")

	//-----------

	var rbuf bytes.Buffer
	rzw := gzip.NewWriter(&rbuf)

	WriteCompress(rzw, "gw", "Wrong Gzip Compression\n")

	wr, _ := zlib.NewReader(&rbuf)
	ReadCompress(wr, "zr")

	//CompressAll(zlib)

	/*

		rdbuf := make([]byte, 2)

		n, err := zr.Read(rdbuf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(n, string(rdbuf))
	*/

	//gzip := &gzip1{
	//	data:   "hello world",
	//	format: "ONE",
	//}

	//gzip.read()

	//zlib := &zlib{
	//	data:     "hello world",
	//	compress: 10,
	//}

	//zlib.read()

	//v, e := zr.Read([]byte("Hello World\n"))
	//fmt.Printf("Value : %d %w", v, e)

	//buf.WriteTo(os.Stdout)
	//os.Stdout.Write(buf.Bytes())

	//x := buf.Bytes()

	//c := make([]byte, len(x))
	//copy(c, x)

	//os.Stdout.Write(c)

	//newbuf := bytes.NewBuffer(c)

	//zt, _ := gzip.NewReader(newbuf)
	//fmt.Printf("%v \n", zt == nil)
	//io.Copy(os.Stdout, zt)

}
