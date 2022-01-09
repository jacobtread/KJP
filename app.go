package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/aerogo/aero"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  uint8  `json:"error_code"`
}

const UserAgent = "KAMAR/ CFNetwork/ Darwin/"
const DefaultKey = "vtku"

func main() {
	// Create a new webserver
	app := aero.New()
	// Maps the pass-through routes and params to KAMAR

	// Middleware for adding CORS Headers
	app.Use(func(handler aero.Handler) aero.Handler {
		return func(context aero.Context) error {
			response := context.Response()
			response.SetHeader("Access-Control-Allow-Origin", "*")
			response.SetHeader("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
			response.SetHeader("Access-Control-Allow-Headers", "*")
			response.SetHeader("Access-Control-Allow-Credentials", "true")
			err := handler(context)
			return err
		}
	})
	// Bind all methods to HandleRequest
	app.Any("/api/:command", HandleRequest)
	// Start the webserver
	app.Run()
}

// HandleRequest handles incoming requests and makes KAMAR requests from them
func HandleRequest(context aero.Context) error {
	// Get the incoming request
	request := context.Request()
	// Search for a mapping for the requested command
	mapping, exists := Mappings[context.Get("command")]
	if !exists { // If the command doesn't exist reply with 404 Not Found
		context.SetStatus(404)
		return context.JSON(ErrorResponse{Error: "No mapping found", Code: 0})
	}
	method := request.Method()
	if mapping.Method != method { // If the request method doesn't match the mapping method
		context.SetStatus(405) // Set the status to 405 Method Not Allowed
		// Let them know the method they should be using instead
		return context.JSON(ErrorResponse{Error: fmt.Sprintf("Unsupported method expected '%s'", mapping.Method)})
	}
	// Get the Authorization header (used for KAMAR keys)
	key := request.Header("Authorization")
	if len(key) < 1 { // If nothing was provided
		if mapping.RequireKey == true { // If a key is required for this mapping
			context.SetStatus(403) // Set the status to 403 Forbidden
			// Let them know the header is required
			return context.JSON(ErrorResponse{Error: "You must provided a Authorization header to access this route"})
		}
		// If it's not required just use the default key "vtku"
		key = DefaultKey
	}
	// Get the portal domain from the X-Portal header
	portalUrl := request.Header("X-Portal")
	if len(portalUrl) < 1 { // If the header is missing
		context.SetStatus(400) // Set the status to 400 Bad Request
		// Let them know that the X-Portal header is required
		return context.JSON(ErrorResponse{Error: "Missing X-Portal header. Don't know where to send the request"})
	}
	// Prepend the protocol and append the api route
	portalUrl = fmt.Sprintf("https://%s/api/api.php", portalUrl)
	// The request values to be provided to the KAMAR request
	values := url.Values{}
	// Add the Key to the values
	values.Set("Key", key)
	// Add the Command to the values
	values.Set("Command", mapping.Command)
	if method == "POST" { // If the method is POST we will have post body
		// This map is used to store the unprocessed JSON
		rawData := map[string]interface{}{}
		// Get the bytes of the request body
		bytes, err := context.Request().Body().Bytes()
		if err != nil { // If we failed to get the body bytes
			context.SetStatus(400) // Set the status to 400 Bad Request
			// Let them know they are missing a body when its required
			return context.JSON(ErrorResponse{Error: "Post request missing request body.", Code: 0})
		}
		// Deserialize the JSON body
		err = json.Unmarshal(bytes, &rawData)
		if err != nil { // If we failed to deserialize the body
			context.SetStatus(400) // Set the status to 400 Bad Request
			// Let them know the JSON they provided is malformed
			return context.JSON(ErrorResponse{Error: "Malformed request body. Expected JSON", Code: 0})
		}
		// Loop over all the mapping parameters
		for key, value := range mapping.Parameters {
			// Check if the raw data contains the parameter key
			data, exists := rawData[key]
			if value.Required && !exists { // If the value is required but is not present
				context.SetStatus(422) // Set the status to 422 Unprocessable Entity
				// Let them know that they didn't provide a required field
				return context.JSON(ErrorResponse{Error: fmt.Sprintf("Missing required field '%s'", key)})
			} else {
				if reflect.ValueOf(data).Kind() == reflect.String { // If the type is a string
					if len(data.(string)) < 1 && value.Required { // Check that its not an empty string and required
						context.SetStatus(422) // Set the status to 422 Unprocessable Entity
						// Let them know that they didn't provide a required field
						return context.JSON(ErrorResponse{Error: fmt.Sprintf("Missing required field '%s'", key)})
					}
				}
				// Set the parameter value with the KAMAR naming
				values.Add(value.Name, data.(string))
			}
		}

	} else if method == "GET" {
		// Loop over all the mapping parameters
		for key, value := range mapping.Parameters {
			// Get the value from the query string
			param := context.Query(key)
			if value.Required && len(param) < 1 { // If the value is required but is not present
				context.SetStatus(422) // Set the status to 422 Unprocessable Entity
				// Let them know that they didn't provide a required field
				return context.JSON(ErrorResponse{Error: fmt.Sprintf("Missing required parameter '%s'", key)})
			}
			// Set the parameter value with the KAMAR naming
			values.Add(value.Name, param)
		}
	}
	// Make the request to KAMAR and get back the XML bytes
	res, err := MakeRequest(portalUrl, context.IP(), &values)
	if err != nil { // If we got an error while connecting to KAMAR
		context.SetStatus(500) // Set the status to 500 Server Error
		// Let them know we failed to connect
		return context.JSON(ErrorResponse{Error: "Failed to get response from KAMAR", Code: 0})
	}
	// Create a new response pointer for the mapping
	out := mapping.Response()
	// Deserialize the XML into the mapping response
	err = xml.Unmarshal(res, &out)
	if err != nil { // If we failed to deserialize the XML
		context.SetStatus(500) // Set the status to 500 Server Error
		// Let them know we parse the response
		return context.JSON(ErrorResponse{Error: "Failed to map response from KAMAR", Code: 0})
	}
	// Return the new JSON mapped response
	return context.JSON(out)
}

// MakeRequest Makes an HTTP request to the provided url with the provided data and returns its bytes
func MakeRequest(endpoint string, ip string, data *url.Values) ([]byte, error) {
	// Create a new HTTP client
	client := &http.Client{}
	// Encode the body data
	encoded := data.Encode()
	// Create a new POST request with the encoded body
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(encoded))
	if err != nil { // If we encountered an error
		return nil, err // Return the error
	}
	// Set the User-Agent header this is required otherwise
	// KAMAR will prevent us from making calls to the API
	req.Header.Add("User-Agent", UserAgent)
	// Set the Content-Type to application/x-www-form-urlencoded
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Set the content length to the body content length
	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
	// Add a forwarded for header so the server knows we are a proxy for the provided IP
	req.Header.Add("X-Forwarded-For", ip)
	// Pretend that the request came from the KAMAR app
	req.Header.Add("X-Requested-With", "nz.co.KAMAR")
	// Send the HTTP request to KAMAR
	res, err := client.Do(req)
	if err != nil { // If we encountered an error
		return nil, err // Return the error
	}
	// Deffer and close the body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)
	// Read all the body data
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { // If we encountered an error
		return nil, err // Return the error
	}
	// Return the response bytes we got back
	return body, nil
}
