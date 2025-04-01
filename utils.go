package main

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
	"path/filepath"
	"github.com/ledongthuc/pdf"
	"github.com/xuri/excelize/v2"
)

func readFile(filePath string) ([]byte, error) {
	// Get the file's absolute path
	filePath, err := filepath.Abs(filePath)

	if err != nil {
		return nil, err
	}

	// Check if the file exists and is accessible
	if _, err := os.Stat(filePath); os.IsNotExist(err) || os.IsPermission(err) {
		return nil, err
	}

	// Read the file contents
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	// Check for different file types and handle them accordingly
	switch strings.ToLower(filepath.Ext(filePath)) {
		case ".pdf":
			f, r, err := pdf.Open(filePath)
			// close file
			defer f.Close()
			
			if err != nil {
				return nil, err
			}
			
			var buf bytes.Buffer
			b, err := r.GetPlainText()
			
			if err != nil {
				return nil, err
			}
			
			buf.ReadFrom(b)
			return buf.Bytes(), nil
		case ".csv":
			// Open the CSV file
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			// Create CSV reader
			reader := csv.NewReader(file)

			var buf bytes.Buffer
			// Read all records
			records, err := reader.ReadAll()
			if err != nil {
				return nil, err
			}

			// Write each record as a tab-separated line
			for _, record := range records {
				buf.WriteString(strings.Join(record, "\t") + "\n")
			}

			return buf.Bytes(), nil
		case ".xlsx":
			// Open the xlsx file
			f, err := excelize.OpenFile(filePath)
			if err != nil {
				return nil, err
			}
			defer f.Close()

			var buf bytes.Buffer
			// Get all sheet names
			sheets := f.GetSheetList()
			
			// Iterate through each sheet
			for _, sheet := range sheets {
				// Write sheet name
				buf.WriteString("Sheet: " + sheet + "\n")
				
				// Get all rows in the sheet
				rows, err := f.GetRows(sheet)
				if err != nil {
					return nil, err
				}
				
				// Write each row
				for _, row := range rows {
					buf.WriteString(strings.Join(row, "\t") + "\n")
				}
				buf.WriteString("\n")
			}
			
			return buf.Bytes(), nil
		default:
			return data, nil
	}
}