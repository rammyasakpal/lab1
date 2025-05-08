Binary Record Sorter in Go

This project implements a high-performance file sorter for fixed-format binary records, built in Go. It demonstrates system-level programming skills, binary parsing, and efficient memory handling—key areas of interest for infrastructure and network.

Overview
The program reads a file containing variable-length binary records with:
- A 4-byte big-endian length prefix
- A 10-byte fixed key
- A variable-length value

It sorts the records lexicographically by key and writes the sorted output to a new binary file.

Features

- Efficient memory-safe parsing of custom binary record structures
- Sorting logic built with sort.Slice and bytes.Compare for byte-level precision
- Designed for system-level file handling and I/O performance
- Validates record length and handles malformed or truncated input robustly

How to Build
bash : go build -o bin/sort src/sort.go
