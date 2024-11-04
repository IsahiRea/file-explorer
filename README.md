# File Explorer

A simple, cross-platform file explorer built with Go and Fyne. This project will provide a graphical interface for navigating directories, creating, renaming, and deleting files and folders.

## Features

- **Basic Interface**: Simple and intuitive GUI with buttons for common file operations.
- **File and Directory Navigation**: Navigate through files and directories.
- **File Operations**: Create, rename, and delete files and folders.

## Requirements

- **Go**: Version 1.18 or higher.
- **Fyne**: Install Fyne: https://docs.fyne.io/started/

## Installation

1. **Install Fyne**:
   ```bash
   go get fyne.io/fyne/v2@latest
   go install fyne.io/fyne/v2/cmd/fyne@latest
   ```

2. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/file-explorer.git
   cd file-explorer
   ```

3. **Run the project**:
   ```bash
   go run main.go
   ```

## Usage

- **Run the Application**:
  After starting, you will see the main file explorer window with a list of files in the current directory.

- **Navigate Folders**:
  - Click on any folder to navigate into it.
  - Use the "Up" button to go up one directory level.

- **File Operations**:
  - **Create File**: Click "Create File" and enter the file name.
  - **Rename File**: Select a file and click "Rename File" to enter a new name.
  - **Delete File**: Select a file and click "Delete File" to remove it.

## Project Structure

```plaintext
file-explorer/
│
├── main.go         # Main application code
├── README.md       # Project documentation
└── go.mod          # Module dependencies
```

## Desired Future Enhancements

- **File Preview**: View file contents or preview images within the app.
- **Sorting and Filtering**: Allow sorting files by name, size, date, etc.
- **Search Functionality**: Implement search to find files quickly.
