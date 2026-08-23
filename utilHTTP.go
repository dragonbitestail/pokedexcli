package main

import (
  "fmt"
	"io"
  "net/http"
	"strings"
)

type requestType string

const (
    reqGetC requestType  = "GET"
    reqPostC requestType = "POST"
)
   
func getResponse(uri string, rType requestType) (*http.Response, error) {
    // Create new request:
    req, err := http.NewRequest(string(rType), uri, nil)
		// OR, to include data in the request body as a chunk of bytes. In this case jsonData
		// would usually be a some json.Marshel'd struct.
    //req, err := http.NewRequest(string(rType), uri, bytes.NewBuffer(jsonData))
    if err != nil {
      fmt.Println("getLastModified(): Returning error for creating request.")
      return nil, err
    }
    
    // Set header on the request:
    // req.Header.Set("x-api-key", "123456789")
    
    // Make the request using '&'.
		// Ensures you won't accidentally duplicate the client state when passing it around your code.
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      fmt.Println("getLastModified(): Returning error for making request.")
      return nil, err
    }
    // CALLER IS EXPECTED TO CLOSE THE BODY!!!! defer resp.Body.Close()
		// OR, optionally, change func behavior to return contents of the body and not the response object.

		logr.Info(getAllRespHeaders(resp.Header))

    return resp, nil
}

//func getResponseWithBodyAsBytes(uri string, rType requestType) (*http.Response, []byte, error) {
func getBodyBytesWithResponseHTTP(uri string, rType requestType) (*http.Response, []byte, error) {
    // Create new request:
    req, err := http.NewRequest(string(rType), uri, nil)
		// OR, to include data in the request body as a chunk of bytes. In this case jsonData
		// would usually be a some json.Marshel'd struct.
    //req, err := http.NewRequest(string(rType), uri, bytes.NewBuffer(jsonData))
    if err != nil {
      return nil, nil, err
    }

    // Set header on the request:
    // req.Header.Set("x-api-key", "123456789")

    // Make the request using '&'.
		// Ensures you won't accidentally duplicate the client state when passing it around your code.
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      fmt.Println("getLastModified(): Returning error for making request.")
      return nil, nil, err
    }
    defer resp.Body.Close()

		bodyBytes, err := getResponseBody(resp)
		if err != nil {
			return nil, nil, err
		}

		// Body becomes unavailble once we return due to the defer'd Body.Close()
		// So, we pull Body out of response and return them separately.
    return resp, bodyBytes, nil
}


func getResponseBody(resp *http.Response) ([]byte, error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
} 

func printAllRespHeaders(respHeader http.Header){
    fmt.Println("RESPONSE Headers::")
    for field, value := range respHeader {
        fmt.Printf("%s => %s\n", field, value)
    }
	return 
}

func getAllRespHeaders(respHeader http.Header) string {
		var lines []string
		lines = append(lines, "HTTP Reponse HEADERS")
    for field, value := range respHeader {
			line := fmt.Sprintf("%s: %s", field, value)
			lines = append(lines, line)
    }
	return strings.Join(lines, "|")
}
