# File Reader

A command-line utility written in Go that reads and extracts text content from various file formats, including PDF, XLSX, and CSV files.

## Features

- Supports multiple file formats:
  - PDF files (.pdf)
  - Excel spreadsheets (.xlsx)
  - CSV files (.csv)
- Preserves document structure in output
- Consistent UTF-8 text output
- Easy-to-use command-line interface

## Installation

1. Make sure you have Go installed on your system
2. Clone the repository:
   ```bash
   git clone https://github.com/RootControl/file-reader.git
   cd file-reader
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Build the executable:
   ```bash
   go build
   ```
5. Move the executable to a directory in your PATH (optional):
   ```bash
   sudo mv file-reader /usr/local/bin/readfile
   ```

## Usage

Basic usage:
```bash
readfile path/to/your/file
```

Examples:
```bash
# Read a PDF file
readfile document.pdf

# Read an Excel file
readfile spreadsheet.xlsx

# Read a CSV file
readfile data.csv
```

## Output Format

- PDF files: Extracted text with preserved formatting
- XLSX files: Tab-separated values with sheet names and structure
- CSV files: Tab-separated values with preserved row structure

## Dependencies

- github.com/ledongthuc/pdf - PDF parsing
- github.com/xuri/excelize/v2 - Excel file parsing
- Standard Go libraries for CSV parsing

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
