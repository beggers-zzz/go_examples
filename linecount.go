// Ben Eggers
// GNU GPL'd

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Demonstrates Go's file/directory IO. Takes a command-line argument representing
// a file or directory, and recursively scans all subdirectories, printing out
// the line count of all files. Also prints out a total. This program is extremely
// naive--it will hang up on directory loops, and will count the lines of binaries
// (which is pretty meaningless, unless you have a weird concept of "meaning").

func main() {
	if len(os.Args) != 2 {
		// Silly user
		fmt.Println("Usage:", os.Args[0], "FILE_OR_DIR_NAME")
		os.Exit(1)
	}

	// Make sure the passed file is valid
	fi, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	counts := make(map[string]uint64) // will store the line counts
	files := make([]os.FileInfo, 0)   // files we still need to deal with

	files = append(files, fi) // set up our file slice with the first file

	// keep going until there are no files left to process
	for len(files) > 0 {
		// get a file from the list
		curFile := files[0]
		files = files[1:] // and slice it off the list

		if !curFile.IsDir() {
			// regular file, open and process

			file, err := ioutil.ReadFile(curFile.Name())
			if err != nil {
				fmt.Println("Got error:", err)
				os.Exit(1)
			}

			counts[curFile.Name()] = 1
			for i := 0; i < len(file); i++ {
				if file[i] == '\n' {
					counts[curFile.Name()]++
				}
			}

		} else {
			// Directory, add all the files from it to our list
			newFiles, err := ioutil.ReadDir(curFile.Name())
			if err != nil {
				fmt.Println("Got error:", err)
				os.Exit(1)
			}

			files = append(files, newFiles...)
		}
	}

	// Now print out the newlines per file, and the total
	total := uint64(0)
	for file, count := range counts {
		fmt.Println(file, ":", count)
		total += count
	}
	fmt.Println("Total:", total)
}
