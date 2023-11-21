package main

import (
   "encoding/csv"
   "encoding/json"
   "fmt"
   "os"
)

func readCSV(fileName string) ([][]string, []string, error) {
	// open csv file
	csvFile, err := os.Open(fileName)

	// return error if error opening csv
	if err != nil {
		return [][]string{}, []string{}, err
   	}
   	defer csvFile.Close()

	// Read the CSV data
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	
	// read column names
	columnNames, _ := reader.Read()

	// read  remaining data
	data, err := reader.ReadAll()

	// return error if reading csv
	if err != nil {
		return [][]string{}, []string{}, err
	}
	return data, columnNames, err
}

func convertToMap(data [][]string, columnNames []string) []map[string] string {

	// make slice of maps to write csv data
	jsonData := make([]map[string] string, len(data))

	// loop over csv data
	for j, row := range data {

		// initialize maps
		jsonData[j] = make(map[string] string)
		for i, col := range row {

			// save csv data in slice of maps
			jsonData[j][columnNames[i]] = col
		}
	}
	return jsonData
}

func main() {

	// parsing arguments
	args := os.Args
	if len(args) != 2 {
		return
	}

	// call to read column names and data in csv
	data, columnNames, readerr := readCSV(args[1])

	// if error return
	if readerr != nil {
		return
	}

	// convert data to a slice of maps
	jsonDatamap := convertToMap(data, columnNames)
	
	// convert to json format
	jsonData, err := json.MarshalIndent(jsonDatamap, "", "	")

	// return if error converting to json 
	if err != nil {
		return
	}

	// display json data
	fmt.Println(string(jsonData))
}