#CS4032 Lab 2: Multithreaded TCP Server
#####Student ID: ea5f6b94d6a8a8f1e7890f6a64883cdc2b6125821e20ddd36a33b773bd46b727
Multithreaded TCP Server written in Go.
The server has thread (goroutine) pooling functionality.

##Dependencies
The Go programming language: https://golang.org/  
Bash  

##Usage Instructions
```bash
./compile.sh
./start.sh
```

start.sh can optionally be passed a port number to run the server on e.g.:  
```bash
./start.sh 8080
```

After running the server you can connect to it using telnet or any other raw TCP client.  
If you send "KILL_SERVICE\n" to the server it will shutdown.  
If you send "HELO text\n" (with text being any arbitrary text) the server will respond with the following text (variables will be replaced by their corresponding values):  
"HELO text\nIP:[ip address]\nPort:[port number]\nStudentID:[your student ID]\n"  
If you send any other message to the server it will run the standard function (otherMessage) to process it.  
The default implemetation included in the server simply echos the message back to the client.  

##Configuration
You can configure the server by editing the values in "config" file.  
The default values are:  
CS4032_LAB_2_IP: 127.0.0.1  
CS4032_LAB_2_PORT: 8080  
CS4032_LAB_2_MAX_WORKERS: 15000  
CS4032_LAB_2_MAX_QUEUE: 20000  

If you provide a port number to the start.sh script that value will be used instead of the one in the config file.  
