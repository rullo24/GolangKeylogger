# Windows - Simple Go Keylogger

A simple Windows keylogger written in Pure Golang for educational purposes. Captures US keyboard input and logs it to a local file.
NOTE: This currently only works for the Windows OS

## Usage

- Ensure this project is used for educational purposes only.
- Run the executable, and the keylogger will start capturing keyboard input.
- Press `Ctrl+C` to stop the keylogger (if run in a terminal).
    - NOTE: Can be run w/o a terminal (refer to top of main.go file)

## Building

To build this program, certain flags are required to ensure that the terminal window is hidden upon launch. To build the program, head to the "src" folder and input the following command:
    --> "go build -ldflags -H=windowsgui"

## About

After searching extensively online, I couldn't find a comprehensive explanation of how to capture keypresses without importing external modules in Golang. 
All of the code in this project is written by myself and is entirely in Golang.

## Disclaimer

This project is intended for educational purposes only. Do not use it for malicious activities. The author is not responsible for any misuse of the code.

## Contributing

Feel free to contribute by submitting issues or pull requests. Contributions are welcome!
