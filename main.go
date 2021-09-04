package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kupec/lzw-example/lzw"
)

func main() {
	var uncompressAction bool

	flag.BoolVar(&uncompressAction, "u", false, "uncompress mode")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Printf("USAGE: %s INPUT_FILE OUTPUT_FILE\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	input_file_path := args[0]
	output_file_path := args[1]

	input_file, err := os.Open(input_file_path)
	if err != nil {
		log.Fatal("Cannot open input file: ", err)
	}
	defer input_file.Close()

	output_file, err := os.Create(output_file_path)
	if err != nil {
		log.Fatal("Cannot create output file: ", err)
	}
	defer output_file.Close()

	err = lzw.Compress(input_file, output_file)
	if err != nil {
		log.Fatal("Cannot compress: ", err)
	}
}
