package main;

import (
	"fmt"
	"flag"
	"./huffmantree"
	"io/ioutil"
	"runtime"
)

func main() {
	src, dst, ext, err := parseArgs();
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Printf("\n%v\n", err)
			default:
				fmt.Println("Usage:")
				flag.PrintDefaults()
				fmt.Println(err)
			}
		}
	}()
	if err != nil {
		panic(err)
	}
	var fileCont []byte
	if fileCont, err = ioutil.ReadFile(*src); err != nil {
		panic(err)
	}
	var processed []byte
	if *ext {
		processed = huffmantree.Decompress(fileCont)
	} else {
		processed = huffmantree.Compress(fileCont)
	}
	if err = ioutil.WriteFile(*dst, processed, 0777); err != nil {
		panic(err)
	}
}

func parseArgs() (*string, *string, *bool, error) {
	srcFile := flag.String("src", "", "specify the `source file` to be compressed or to be decompressed")
	dstFile := flag.String("dst", "", "specify the `destination file` to be compressed or to be decompressed")
	isExt := flag.Bool("ext", false, "extract file")
	flag.Parse()
	return srcFile, dstFile, isExt, nil
}
