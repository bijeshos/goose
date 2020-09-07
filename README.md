# Goose
A utility program written in Go
 
# Build Instructions
- Clone the repository
- Grant execute permission for the build script
    - On Linux
        - `chmod 775 build.sh`
- Execute build script
    - On Linux
        - `./build.sh`
             
- For creating the binary without using build script, choose any of the following options:
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

- `$ goose sync dir`
    - This will sync source directory with target directory
    - Files from source directory would be copied to target directory




**_[work in progress]_**