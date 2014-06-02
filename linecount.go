// Ben Eggers
// GNU GPL'd

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Demonstrates Go's file/directory IO. Takes a command-line argument representing
// a file or directory, and recursively scans all subdirectories, printing out
// the line count of all files. Also prints out a total. This program is extremely
// naive--it will hang up on directory loops, and will count the lines of binaries
// (which is pretty meaningless, unless you have a weird concept of "meaning").

func main() {
	// Command-line argument stuff -- see flags.go
	verbose := false // whether or not to print all files
	ext := ""        // filetypes to count
	flag.StringVar(&ext, "ext", "", "Types of files to count.")
	flag.BoolVar(&verbose, "v", false, "Whether to print every file, "+
		"instead of just the total.")
	flag.Parse()

	counts := make(map[string]uint64) // will store the line counts
	files := flag.Args()              // files we still need to deal with

	// keep going until there are no files left to process
	for len(files) > 0 {
		// get a file from the list
		curFile := files[0]
		files = files[1:] // and slice it off the list

		// Make sure the file is valid
		fi, err := os.Stat(curFile)
		if err != nil {
			fmt.Println("error on stat", err)
			os.Exit(1)
		}

		if !fi.IsDir() {
			// regular file, open and process

			// Get all the bytes, and check the validity in the process
			file, err := ioutil.ReadFile(curFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// count the newlines in the file
			counts[curFile] = 1
			for i := 0; i < len(file); i++ {
				if file[i] == '\n' {
					counts[curFile]++
				}
			}

		} else {
			// Directory, add all the files from it to our list

			// Get all the subFiles, and check the validity
			newFiles, err := ioutil.ReadDir(curFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Add them all to our files list
			for i := 0; i < len(newFiles); i++ {
				f := newFiles[i]

				if !strings.HasPrefix(f.Name(), ".") &&
					(f.IsDir() || strings.HasSuffix(f.Name(), ext)) {
					// only add files if they're a directory, or have our extension
					// (and aren't dotfiles)
					files = append(files, curFile+"/"+f.Name())
				}
			}
		}
	}

	// Now print out the newlines per file, and the total
	total := uint64(0)
	for file, count := range counts {
		total += count
		if verbose {
			fmt.Println(file, ":", count)
		}
	}

	fmt.Println("Total:", total)
}
