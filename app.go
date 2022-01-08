package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/aerogo/aero"
	"io"
	"io/ioutil"
	"log"
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
	app.Any("/api/:command", HandleRequest)
	app.Run()
}

func HandleRequest(context aero.Context) error {
	request := context.Request()

	mapping, exists := Mappings[context.Get("command")]
	if !exists {
		return context.JSON(ErrorResponse{Error: "Command mapping unknown?!", Code: 0})
	}

	method := request.Method()
	if mapping.Method != method {
		context.SetStatus(405)
		return context.JSON(ErrorResponse{Error: fmt.Sprintf("Unsupported method expected '%s'", mapping.Method)})
	}

	key := request.Header("Authorization")
	if len(key) < 1 {
		key = DefaultKey
	}

	portalUrl := request.Header("X-Portal")
	if len(portalUrl) < 1 {
		context.SetStatus(400)
		return context.JSON(ErrorResponse{Error: "Missing X-Portal header. Don't know where to send the request"})
	}

	portalUrl = fmt.Sprintf("https://%s/api/api.php", portalUrl)

	params := map[string]interface{}{}
	if method == "POST" {
		rawData := map[string]interface{}{}
		err := parsePostData(&rawData, context)
		if err != nil {
			return err
		}
		for key, value := range mapping.Parameters {
			data, exists := rawData[key]
			if value.Required && !exists {
				context.SetStatus(422)
				return context.JSON(ErrorResponse{Error: fmt.Sprintf("Missing required field '%s'", key)})
			} else {
				if reflect.ValueOf(data).Kind() == reflect.String {
					if len(data.(string)) < 1 && value.Required {
						context.SetStatus(422)
						return context.JSON(ErrorResponse{Error: fmt.Sprintf("Missing required field '%s'", key)})
					}
				}
				params[value.Name] = data
			}
		}

	} else if method == "GET" {
		for key, value := range mapping.Parameters {
			param := context.Query(key)
			if value.Required && len(param) < 1 {
				context.SetStatus(422)
				return context.JSON(ErrorResponse{Error: fmt.Sprintf("Missing required parameter '%s'", key)})
			}
			params[value.Name] = param
		}
	} else {
		context.SetStatus(405)
		return context.JSON(ErrorResponse{Error: "Unsupported method only GET and POST are supported."})
	}

	data := url.Values{}
	data.Set("Command", mapping.Command)
	for key, value := range params {
		data.Set(key, value.(string))
	}

	res, err := GetXML(portalUrl, context.IP(), &data)
	if err != nil {
		context.SetStatus(500)
		return context.JSON(ErrorResponse{Error: "Failed to get response from KAMAR", Code: 0})
	}

	fmt.Println(string(res))
	var out = mapping.Response()
	err = xml.Unmarshal(res, &out)
	if err != nil {
		fmt.Println(err)
		context.SetStatus(500)
		return context.JSON(ErrorResponse{Error: "Failed to map response from KAMAR", Code: 0})
	}

	fmt.Println(out)

	return context.JSON(out)
}

func parsePostData(rawData *map[string]interface{}, context aero.Context) error {
	bytes, err := context.Request().Body().Bytes()
	if err != nil {
		context.SetStatus(400)
		return context.JSON(ErrorResponse{Error: "Post request missing request body.", Code: 0})
	}
	err = json.Unmarshal(bytes, &rawData)
	if err != nil {
		context.SetStatus(400)
		return context.JSON(ErrorResponse{Error: "Malformed request body. Expected JSON", Code: 0})
	}
	return nil
}

func GetXML(endpoint string, ip string, data *url.Values) ([]byte, error) {
	if !data.Has("Key") {
		data.Add("Key", DefaultKey)
	}
	client := &http.Client{}

	encoded := data.Encode()

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(encoded))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "KAMAR/ CFNetwork/ Darwin/")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
	req.Header.Add("X-Forwarded-For", ip)
	req.Header.Add("X-Requested-With", "nz.co.KAMAR")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println(res.Status)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
