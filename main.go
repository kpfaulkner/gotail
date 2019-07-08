package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// tail functionality....   simple :)
func main() {

	filename := flag.String("f", "", "the filename to tail")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats,err := file.Stat()
	if err != nil {
		log.Fatalf("Unable to process file: %s\n", err.Error())
	}

	fileSize := stats.Size()

	// will just read last 10000 bytes of file
	if fileSize > 10000 {
		file.Seek( int64(fileSize - int64(10000)), 0)
	}
	scanner := bufio.NewReader(file)

	for {


		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// EOF of file... just sleep a bit and continue again
				time.Sleep(time.Duration(1) * time.Second)
			} else {
				fmt.Printf("break %s\n", err.Error())
				break
			}
		}
    //l := strings.TrimRight( line,"\n")
    fmt.Printf("%s", string(line))
	}

}
