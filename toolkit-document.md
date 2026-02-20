1. Title and Objective
Project Title

Prompt-Powered Kickstart: Building a Beginner’s Toolkit for Go and Gin

Objective

The objective of this project is to learn the Go programming language and the Gin web framework using Generative AI prompts and to create a beginner-friendly toolkit that helps new learners build and run a simple REST API.

Technology Chosen

Go Programming Language

Gin Web Framework

Why I Chose This Technology

Go is widely used for building backend services due to its simplicity performance and fast compilation. Gin provides a lightweight and beginner-friendly way to build REST APIs in Go.

End Goal

To build a minimal REST API with two endpoints that returns JSON responses and can be tested using a browser or curl.

2. Quick Summary of the Technology

Go is a statically typed compiled programming language designed for simplicity and efficiency. It is commonly used in backend development cloud services and microservices.

Gin is a fast and minimal web framework for Go that simplifies HTTP routing and JSON responses.

Real-world use case:
Many production APIs and backend services use Go and Gin to serve web and mobile applications.

3. System Requirements

Operating System: Windows Linux or macOS

Tools Required:

Go version 1.20 or higher

Terminal or Command Prompt

Code editor such as VS Code

Packages:

Gin web framework

4. Installation and Setup Instructions
Step 1: Install Go

Download and install Go from the official website. After installation restart the terminal.

Verify installation:

go version

Step 2: Create the project folder
mkdir go-gin-toolkit
cd go-gin-toolkit

Step 3: Initialize Go module
go mod init go-gin-toolkit


This creates a go.mod file which manages dependencies.

Step 4: Install Gin framework
go get -u github.com/gin-gonic/gin

5. Minimal Working Example
Description

This example creates a REST API with two endpoints:

/ returns a welcome message

/status returns the API status

The server runs on port 8080 by default and supports environment-based configuration.

Code: main.go
package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go and Gin Beginner Toolkit",
		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API is running",
		})
	})

	router.Run(":" + port)
}

Running the application
go run main.go

Testing the API
curl http://localhost:8080


Expected output:

{"message":"Welcome to the Go and Gin Beginner Toolkit"}

curl http://localhost:8080/status


Expected output:

{"status":"API is running"}

6. AI Prompt Journal
Prompt 1

Prompt Used:
Explain how to install Go and verify it on my system.

Outcome:
The AI provided installation steps and explained how to verify Go using the terminal.

Evaluation:
Helpful for setting up the environment correctly.

Prompt 2

Prompt Used:
How do I initialize a Go module and why is go.mod required.

Outcome:
The AI explained dependency management and project structure.

Evaluation:
Very helpful for understanding Go project setup.

Prompt 3

Prompt Used:
I am getting a syntax error unexpected name gin in c.JSON how do I fix it.

Outcome:
The AI identified a missing comma in the function arguments.

Evaluation:
Highly helpful for debugging a real error.

Prompt 4

Prompt Used:
curl cannot connect to localhost what could be wrong with my Go API.

Outcome:
The AI helped diagnose server startup and port configuration issues.

Evaluation:
Helpful in understanding how to test APIs correctly.

7. Common Issues and Fixes

go command not found
Solution: Install Go and ensure it is added to PATH.

Syntax error in c.JSON
Solution: Ensure commas are correctly placed between arguments.

curl failed to connect
Solution: Confirm the server is running and listening on the correct port.

8. References

Official Go Documentation

Gin Web Framework Documentation

MDN Web Docs for HTTP status codes