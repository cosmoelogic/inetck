# inetck
A simple program to display your systems' internal network addresses. Because 'ip a' just displays too much text

## Usage
### Use A Release
1. Head over to the latest release and download the executable
2. If the file is not executable for you, you can use `chmod +x inetck`
3. Then, simply execute the file

### Build Manually
#### Build Requirements
* Go
#### One Liner
`git clone https://github.com/cosmoelogic/inetck/ && cd inetck && go build -o inetck main.go`
#### Step-By-Step
1. `git clone https://github.com/cosmoelogic/inetck/`
2. `cd inetck`
3. `go build -o inetck main.go`

### And if you want it available everywhere
`mv inetck /usr/local/bin/inetck`
