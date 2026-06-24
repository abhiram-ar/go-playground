package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	exisingFilePath := "./testFiles/helloworld.txt"
	newFilePath := "./testFiles/iamtheCreatedOne.txt"

	fileExist, err := fileExists(exisingFilePath)
	fmt.Println("file exists:", fileExist, err)

	fmt.Println("file created:", createFile(newFilePath))

	fmt.Println("write file:", writeFile(newFilePath, []byte("hi there new file")))

	fmt.Println("append file:", appendFile(newFilePath, []byte("::appened to the file file")))

	data, err := readFile(exisingFilePath)
	fmt.Println("read file:", string(data), err)

	size, modTime, _ := fileStatus(exisingFilePath)
	fmt.Println("file size", size)
	fmt.Println("file last modified time", modTime)

	// fmt.Println("delete file:", deleteFile(newFilePath))

}

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// just dont want to return err, since false is discriptive enough in this case
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func createFile(path string) error {
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}

	// newFile is a file descriptor
	// if we dont close it after the job is done, will use memory and cause memory leak
	defer newFile.Close()

	return nil
}

func createFileIfNotExists(path string) error {
	file, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_EXCL|os.O_CREATE, // bitwise OR to combine the permissions
		0666,                            // unmasked chmod for the created file - if not exist
	)

	// os.O_RDWR   - open for read and write
	// os.O_CREATE - create the file if it doesn't exist
	// os.O_EXCL   - "exclusive." This is the key one. When combined with O_CREATE, it means "create the file, but fail if it already exists."

	// Why use O_EXCL at all?
	// It's an atomic check-and-create. Doing "check if file exists, then create it" in two separate steps has a race condition:
	// another process could create the file between your check and your create. O_CREATE|O_EXCL makes the operating system do both in a single atomic operation,
	// which is the correct way to claim a file exclusively (commonly used for lock files).
	// Another reason is that using only O_CREATE will create a file if it does not exist, it wont return an error

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func writeFile(path string, data []byte) error {

	// WriteFile overrides existing file while writing
	return os.WriteFile(
		path,
		data,
		0644, // incase the file does not exist, create a file with this permission
	)
}

func appendFile(path string, data []byte) error {

	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_APPEND, // open in write only mode with append flag
		0644,
	)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	return err
}

func readFile(path string) ([]byte, error) {
	// better use a buffered reader,
	// for simplicy we are reading all the data at once to the ram
	return os.ReadFile(path)

}

func deleteFile(path string) error {
	return os.Remove(path)
}

func fileStatus(path string) (int64, time.Time, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return -1, time.Time{}, err
	}

	size := info.Size()
	modTime := info.ModTime()
	return size, modTime, nil
}
