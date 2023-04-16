package wttr

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func ShowWTTR() {
    // Create an HTTP client
    client := &http.Client{}

    // Create a new GET request
	req, err := http.NewRequest("GET", "http://wttr.in/?A", nil)
	if err != nil {
	    fmt.Println("Error creating request:", err)
	    return
    }

    // Add headers if needed
    req.Header.Add("HeaderName", "HeaderValue")

    // Send the request and get the response
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }

    // Close the response body when we're done
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    // Print the response body
	fmt.Println(string(body))
}

