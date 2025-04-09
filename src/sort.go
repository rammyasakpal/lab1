package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

//defining a record structure -1

type Record struct {
	Key   [10]byte
	Value []byte
	Raw   []byte
}

//Reading the incomping records-step 2

func readRecords(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var records []Record
	for {
		var lengthBuf [4]byte
		_, err := io.ReadFull(file, lengthBuf[:])
		if err == io.EOF {
			break //reading completed
		} else if err != nil {
			return nil, err
		}

		length := binary.BigEndian.Uint32(lengthBuf[:])
		if length < 10 {
			return nil, fmt.Errorf("invalid length: %d", length)
		}

		data := make([]byte, length)
		_, err = io.ReadFull(file, data)
		if err != nil {
			return nil, err
		}

		var key [10]byte
		copy(key[:], data[:10])
		value := data[10:]

		full := append(lengthBuf[:], data...)
		records = append(records, Record{Key: key, Value: value, Raw: full})

	}

	return records, nil
}

//output file

func writeRecords(filename string, records []Record) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, record := range records {
		_, err := file.Write(record.Raw)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("usage: %v inputfile outputfile\n", os.Args[0])
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])

	records, err := readRecords(inputFile)
	if err != nil {
		log.Fatalf("failed to read records: %v", err)
	}

	sort.Slice(records, func(i, j int) bool {
		return bytes.Compare(records[i].Key[:], records[j].Key[:]) < 0
	})

	err = writeRecords(outputFile, records)
	if err != nil {
		log.Fatalf("failed to write records: %v", err)
	}

}
