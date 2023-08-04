package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Replace YOUR_API_KEY with your actual Tinify API key
	apiKey := "1PW34PQl8SQMzMFhwCF4fllYtWmDnPxF"
	imagePath := "download-bg.png"

	// Read the image from the local file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error reading the image file:", err)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.tinify.com/shrink", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set the basic authentication header with your Tinify API key
	req.SetBasicAuth("api", apiKey)

	// Set the request body with the image data
	req.Body = ioutil.NopCloser(bytes.NewReader(imageData))

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	imgURL := resp.Header.Get("Location")
	resp, err = http.Get(imgURL)
	//print resp
	fmt.Println(resp)
	fmt.Println(resp.Body)
	fmt.Println(resp.Header)
	if err != nil {
		fmt.Println("Error sending HTTP GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body, which contains the compressed image
	compressedData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Write the compressed image to a new file
	compressedImagePath := "compressed.png"
	err = ioutil.WriteFile(compressedImagePath, compressedData, 0644)
	if err != nil {
		fmt.Println("Error writing the compressed image:", err)
		return
	}

	fmt.Println("Image compression successful. Compressed image saved to", compressedImagePath)
}
