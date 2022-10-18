package main

import "os"

func main() {
	// os.Mkdir()
	// os.MkdirAll()
	// err := os.Remove()
	f, err := os.Create("file.txt")
	if err != nil {
		return
	}
	defer f.Close()
	for i := 0; i < 10; i++ {
		f.WriteString("Hello")
		f.Write([]byte("Sample write"))
	}
	// Read from files
	f1, err := os.Open("filex.txt")
	if err != nil {
		return
	}
	defer f1.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	// Delete files
	err = os.Remove("fienaem.txt")
	if err != nil {
		return
	}
}
