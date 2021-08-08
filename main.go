package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CsvLine struct {
	Id     string
	Images []string
}

func main() {
	lines, err := ReadCsv("hydepark.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			Id: line[0],
			Images: []string{
				line[1],
				line[2],
				line[3],
				line[4],
				line[4],
			},
		}

		fmt.Println(data.Id)
		i := 1
		for _, fileUrl := range data.Images {
			file := fmt.Sprintf("hydepark/%s_media%d.jpeg", data.Id, i)
			fmt.Println("Saving as " + file)
			err := DownloadFile(file, fileUrl)
			if err != nil {
				panic(err)
			}
			fmt.Println("Downloaded: " + fileUrl)
			i++
		}
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
