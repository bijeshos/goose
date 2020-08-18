# Goose
A utility program written in Go

# Build Instructions
- Clone the repository
- Build the project as follows (choose any):
    - `go build -o goose`
    - `env GOOS=windows GOARCH=amd64 go build -o goose.exe` (for Windows 64 bit)
    - `env GOOS=linux GOARCH=amd64 go build -o goose` (for Linux 64 bit)
- Install the binaries as follows
    - `go install -o goose github.com/bijeshos/goose`
- to run the binary    
    - `./goose`
- to run main program    
    - `go run main.go`

# Configuration
- Defaults
    - Config file location: `$HOME/.goose/goose-config.yaml`
    - Ignore file location: `$HOME/.goose/goose.ignore`
 

# Commands

- `goose sync dir`
    - This will sync source and target directories




**_[work in progress]_**