# add-latency

Windows build command: `env GOOS=windows go build -ldflags "-X main.os=windows -H=windowsgui -s -w" -o ./builds/windows.exe`
MacOS build command: `env GOOS=darwin go build -ldflags "-X main.os=darwin -s -w" -o ./builds/macos`
Linux build command: `env GOOS=linux go build -ldflags "-X main.os=linux -s -w" -o ./builds/linux`

