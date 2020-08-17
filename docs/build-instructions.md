# Build instructions

- clone the repository
- to build the project
    - `go build -o goose github.com/bijeshos/goose`
    - `env GOOS=windows GOARCH=amd64 go build -o goose.exe` (for Windows 64 bit)
    - `env GOOS=linux GOARCH=amd64 go build -o goose.exe` (for Linux 64 bit)
- to install the binaries
    - `go install -o goose github.com/bijeshos/goose`
- to run the binary    
    - `./goose`
- to run main program    
    - `go run main.go`
    
    
    