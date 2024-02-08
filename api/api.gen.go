// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
)

// SensorData defines model for SensorData.
type SensorData struct {
	SensorID  *int       `json:"SensorID,omitempty"`
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	Value     *float32   `json:"Value,omitempty"`
}

// AddSensorDataJSONRequestBody defines body for AddSensorData for application/json ContentType.
type AddSensorDataJSONRequestBody = SensorData

// GlobalResponses defines the response model for the global responses.
type GlobalResponses struct {
	ResponseCode    string `json:"responseCode,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add Sensor Data
	// (POST /api/v1.0/sensor-data)
	AddSensorData(c *fiber.Ctx) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// AddSensorData operation middleware
func (siw *ServerInterfaceWrapper) AddSensorData(c *fiber.Ctx) error {

	return siw.Handler.AddSensorData(c)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Post(options.BaseURL+"/api/v1.0/sensor-data", wrapper.AddSensorData)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/3RSO2/cMAz+KwLb0T5f+li8JUgHDwEKJOhSZOBJPFupXpXoaw4H//eCti/1kE58fhT1",
	"fbyAjj7FQIELtBcoeiCPs/tIocR8j4wSpRwTZba0qXX34vM5EbRgA1NPGaYKnqynwuiTlI8xe2RowSBT",
	"zdYTVFdM4WxDL5Af6EbaTAujP8iw6a03Hl5IM0ySsuEYpdlZTaHMuIBeuh66JxnHlp2Ey55KPqFuv3dQ",
	"wYlysTFACze7/W4PFbzWfawT6l/YC6S3PIyHnY6+iZhsraOhnkKTx7Au/1pvC7W3xjj6g1mo+QkPbyE8",
	"TxXERAGThRY+r+8l5GEmscFkm9PNbt+Uec3aXLmOhcUK48g2hs5AC7fGbCSpINPvkQrfRXOWZh0DU5hx",
	"mJKzekY2L0U+exVWvI+ZjtDCh+af8s0qe7N5YCbaUNHZJl4YW9mUNRVHhcaI4YHm1AELrWvZTAZaziNN",
	"kigphrIczqf9Xsz/56IxZFQZtaZSjqNzZ9Hzy3uwOzRqJUHVqgsndNYoG9LI8ywBfn0P2AWmHNCpR8on",
	"yupbznG5tTJ6j/m8sK0217PQUeZ+kfkCY3bQwsCc2qZxUaMbRLbpefobAAD//8PdFJJUAwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}