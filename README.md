# TCP-to-HTTP Server – Go

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![TCP](https://img.shields.io/badge/TCP-6DB33F?logo=network&logoColor=white)
![HTTP](https://img.shields.io/badge/HTTP-F0DB4F?logo=http&logoColor=white)

---

## Table of Contents

- [Overview](#overview)  
- [Learning Objectives](#learning-objectives)  
- [Project Description](#project-description)  
- [Features](#features)  
- [Application Structure](#application-structure)  
- [Tech Stack](#tech-stack)  
- [Setup Instructions](#setup-instructions)  
- [How It Works](#how-it-works)  
- [Why I Built This](#why-i-built-this)  
- [Future Improvements](#future-improvements)  
- [License](#license)

---

## Overview

This project is a **TCP-to-HTTP server** built from scratch in **Go**.  
It demonstrates how to build a **minimal HTTP server** using TCP sockets, without relying on any high-level web frameworks.

---

## Learning Objectives

Through this project, I learned to:

- Work with **raw TCP connections** in Go  
- Parse **basic HTTP requests manually**  
- Send **HTTP responses** over TCP  
- Use **Go’s concurrency primitives** for handling multiple connections  
- Structure a Go project with proper modules and packages  

---

## Project Description

The server listens on a TCP port and converts incoming requests into HTTP responses.  

It demonstrates:

- Handling multiple clients concurrently using **goroutines**  
- Parsing HTTP request lines  
- Sending valid HTTP responses  
- A foundation for building more advanced HTTP servers  

---

## Features

- Minimal TCP-to-HTTP server  
- Handles multiple connections concurrently  
- Parses HTTP request lines  
- Sends HTTP 200 responses with plain text content  
- Simple, lightweight, and written entirely in Go  

---

## Application Structure

```

tcptohttp/
│
├── main.go          # Entry point for the TCP-to-HTTP server
└── go.mod           # Go module definition

````

---

## Tech Stack

- **Go**  
- **TCP sockets**  
- **HTTP protocol basics**  
- **Goroutines**  

---

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/SkanderGasmi/tcptohttp.git
cd tcptohttp
````

### 2. Initialize Go modules (if not already done)

```bash
go mod tidy
```

### 3. Run the server

```bash
go run .
```

### 4. Access the server

Open your browser or terminal:

```text
http://localhost:8080
```

Or use `curl`:

```bash
curl http://localhost:8080
```

You should see:

```
Hello World!
```

---

## How It Works

1. The server **listens on TCP port 8080**.
2. Incoming connections are accepted and handled in **separate goroutines**.
3. The server reads the first line of the HTTP request.
4. A basic **HTTP response** is sent back to the client.

This minimal implementation gives full control over **TCP and HTTP** handling in Go.

---

## Why I Built This

I built this project to **understand the fundamentals of networking and HTTP servers**.

It demonstrates:

* TCP socket programming in Go
* HTTP protocol basics
* Concurrent connection handling using goroutines

---

## Future Improvements

* Fully parse HTTP headers and methods
* Implement support for multiple routes
* Serve static files (HTML, CSS, JS)
* Add logging and metrics for requests
* Convert it into a lightweight production-ready HTTP server
