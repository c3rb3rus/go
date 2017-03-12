package main

import "fmt"

import "io/ioutil"
import "os"
import "log"
import "bufio"

func main() {
	var err error
	err = readFile()
	if err != nil {
		log.Fatal(err)
	}
	err = writeFile()
	if err != nil {
		log.Fatal(err)
	}
	err = writeFile2()
	if err != nil {
		log.Fatal(err)
	}
	err = readFile2()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile() error {
	f, err := os.Open("readme.txt")
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Cannot open file: %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Cannot read file: %v", err)
	}

	fmt.Printf("%s", string(b))
	return nil
}

func readFile2() error {
	f, err := os.Open("readme2.txt")
	if err != nil {
		return fmt.Errorf("Cannot open file: %v", err)
	}
	r := bufio.NewReader(f)
	line, _, err := r.ReadLine()
	if err != nil {
		return fmt.Errorf("Cannot read file: %v", err)
	}
	fmt.Println(string(line))
	return nil
}

func writeFile() error {
	s := "my text"
	err := ioutil.WriteFile("written.txt", []byte(s), 0644)
	if err != nil {
		return fmt.Errorf("Cannot write to file: %v", err)
	}
	return nil
}

func writeFile2() error {
	var err error
	f, err := os.Create("written2.txt")
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Cannot create file: %v", err)
	}
	_, err = f.Write([]byte("string1"))
	if err != nil {
		return fmt.Errorf("Cannot write to file: %v", err)
	}
	_, err = f.WriteAt([]byte("o"), 3)
	if err != nil {
		return fmt.Errorf("Cannot write to file: %v", err)
	}
	_, err = f.WriteString("string2")
	if err != nil {
		return fmt.Errorf("Cannot write to file: %v", err)
	}
	return nil
}
