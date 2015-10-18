#CS4032 Lab 2: Multithreaded TCP Server
Multithreaded TCP Server written in Go.

##Dependencies
The Go programming language
Bash

##Usage Instructions
```bash
./compile.sh
./start.sh
```
After running the server you can connect to it using telnet or any other raw TCP client.
If you send "KILL_SERVICE\n" to the server it will shutdown.
If you send "HELO text\n" (with text being any arbitrary text) the server will respond with the following text (variables will be replaced by their corresponding values):
"HELO text\nIP:[ip address]\nPort:[port number]\nStudentID:[your student ID]\n"
If you send any other message to the server it will run the standard function (otherMessage) to process it.
The default implemetation included in the server simply echos the message back to the client.
