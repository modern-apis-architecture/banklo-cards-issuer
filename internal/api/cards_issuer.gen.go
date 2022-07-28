// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Defines values for NotificationStatus.
const (
	Created NotificationStatus = "created"
	Denied  NotificationStatus = "denied"
)

// AccountResponse defines model for account-response.
type AccountResponse struct {
	Age      *int32  `json:"age,omitempty"`
	Document *string `json:"document,omitempty"`
	Id       *string `json:"id,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// CardResponse defines model for card-response.
type CardResponse struct {
	Id         *string `json:"id,omitempty"`
	LastDigits *string `json:"last_digits,omitempty"`
	Status     *string `json:"status,omitempty"`
}

// CardId defines model for card_id.
type CardId = string

// CreateAccountRequest defines model for create-account-request.
type CreateAccountRequest struct {
	PersonalData *struct {
		Address *struct {
			Number  *string `json:"number,omitempty"`
			ZipCode *string `json:"zip_code,omitempty"`
		} `json:"address,omitempty"`
		BirthDate  *openapi_types.Date `json:"birth_date,omitempty"`
		Document   *string             `json:"document,omitempty"`
		LastName   *string             `json:"last_name,omitempty"`
		MotherName *string             `json:"mother_name,omitempty"`
		Name       *string             `json:"name,omitempty"`
		Phone      *struct {
			Code   *string `json:"code,omitempty"`
			Number *string `json:"number,omitempty"`
		} `json:"phone,omitempty"`
	} `json:"personal_data,omitempty"`
}

// CreateAccountResponse defines model for create-account-response.
type CreateAccountResponse struct {
	Id     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

// CreateCardRequest defines model for create-card-request.
type CreateCardRequest struct {
	ValidUntil *openapi_types.Date `json:"valid_until,omitempty"`
}

// CreateCardResponse defines model for create-card-response.
type CreateCardResponse struct {
	Id *string `json:"id,omitempty"`
}

// Notification defines model for notification.
type Notification struct {
	CardId CardId             `json:"card_id"`
	Status NotificationStatus `json:"status"`
}

// NotificationStatus defines model for Notification.Status.
type NotificationStatus string

// CreateAccountJSONBody defines parameters for CreateAccount.
type CreateAccountJSONBody = CreateAccountRequest

// CreateCardsJSONBody defines parameters for CreateCards.
type CreateCardsJSONBody = CreateCardRequest

// PostCardsIdSubscribeJSONBody defines parameters for PostCardsIdSubscribe.
type PostCardsIdSubscribeJSONBody struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody = CreateAccountJSONBody

// CreateCardsJSONRequestBody defines body for CreateCards for application/json ContentType.
type CreateCardsJSONRequestBody = CreateCardsJSONBody

// PostCardsIdSubscribeJSONRequestBody defines body for PostCardsIdSubscribe for application/json ContentType.
type PostCardsIdSubscribeJSONRequestBody PostCardsIdSubscribeJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// It creates accounts
	// (POST /accounts)
	CreateAccount(ctx echo.Context) error
	// Get account details
	// (GET /accounts/{id})
	GetAccount(ctx echo.Context, id CardId) error
	// It creates cards
	// (POST /accounts/{id}/cards)
	CreateCards(ctx echo.Context, id CardId) error

	// (DELETE /cards/{id})
	DeleteCard(ctx echo.Context, id CardId) error
	// Get card details
	// (GET /cards/{id})
	GetCard(ctx echo.Context, id CardId) error
	// Subscribe to notifications about given card
	// (POST /cards/{id}/subscribe)
	PostCardsIdSubscribe(ctx echo.Context, id CardId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAccount(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateAccount(ctx)
	return err
}

// GetAccount converts echo context to params.
func (w *ServerInterfaceWrapper) GetAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id CardId

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAccount(ctx, id)
	return err
}

// CreateCards converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCards(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id CardId

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateCards(ctx, id)
	return err
}

// DeleteCard converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id CardId

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteCard(ctx, id)
	return err
}

// GetCard converts echo context to params.
func (w *ServerInterfaceWrapper) GetCard(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id CardId

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCard(ctx, id)
	return err
}

// PostCardsIdSubscribe converts echo context to params.
func (w *ServerInterfaceWrapper) PostCardsIdSubscribe(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id CardId

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostCardsIdSubscribe(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/accounts", wrapper.CreateAccount)
	router.GET(baseURL+"/accounts/:id", wrapper.GetAccount)
	router.POST(baseURL+"/accounts/:id/cards", wrapper.CreateCards)
	router.DELETE(baseURL+"/cards/:id", wrapper.DeleteCard)
	router.GET(baseURL+"/cards/:id", wrapper.GetCard)
	router.POST(baseURL+"/cards/:id/subscribe", wrapper.PostCardsIdSubscribe)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RXa2/bOBb9KwLbD7uAWMuS/JCBRdEHUKTALoq239psSolXDhuJ1JBUWo+h/z4g9bJs",
	"KfZ0msE0n2I+Lu89555Dao8SkReCA9cKbfZIJbeQE/svSRJRco0lqEJwBWYMvpO8yOy/ZAtoE4QuoiIp",
	"c+AabZA/9wM/8OeBH8yRixhFG0TXHlmtwhhDFPg4XCVLHC2iAPtxAmRJqD9fxshFGVH6hpMc0Aa9FoBc",
	"1Px4KwBVLiqkKEBqBqo7fI9SIXNiDmZcBz5ykd4VUP+ELUizr89u384qLRnfmkmT4MjwQS4jsxMTlYs0",
	"0wYb9F4I7XzcFeCkQjonOHZpivgrJNrETIikU0DXKAbLte/FKfZXfoTDNEpw5Ech9lZ07S1oFNO0Q5Gy",
	"LTNsIkMEcpHSRJfmNykKKe6BnuL5EBBtuJH5NvLlYAwLnUDipk6nIFqD5GiDvvz/06fNd5vI5vp6v67+",
	"9Xzz+TMeDIbVv/dBdTw696unX7Z5f1KffCKBaMA9Pb+VoPQR+AVIJTjJbijRxHYepRKUrZmXeQwSbVCw",
	"QC76nRU3iaDmDKOCEEeeZ46JmdS3ZrudidYLPPexF6GhcLx15MM6WkejUsiFvgXZjv2XyJ3z9lgjLipu",
	"BbdJt2mYYF2SvvnDQRAEqKpOGuCkziO99VUPJ9rwI83RAzLaHsfEH+J0oGw7MMLeg7p+WMADNC8X+AHA",
	"QwwminSnwTmtf3xkXELjfTumpeOV0/7iJ4sw8ZYx9sJ4iUPfI3jteylOfRJRSpYk9dNDK6lDX+4kP+AU",
	"E8lP19l4y5iM70nG6E3JNcus2nwPeyH21qf5D1ae7cNz2Q9SOpv5ND3z1TxeLuYBJos54DBdERyT9Qr7",
	"QbyGaBGv0zC+kIsLU55GmwvNUpYQzQQfEUPv308lpGiDnsz6J8aseV/M2mWD1gBe5mjzqWsuF1HgDCi6",
	"HgPewMokULujCddFuz7Vk7nveSrMQRRUIllRV4BeEn6XCQc7r4ikymFKlSCdF++uFOqxenDRPUhVx5o/",
	"855Z2xcFcFIwczvYIddcZ7e2zFnT07WdirpZDYYW0ivD9yuLwIt6HaprBaVfCrqrHYfrxvtIUWQNGbOv",
	"qmakBvksBeM+Ug2x1bIEO1A3hM3Z9+aPl0XTeDaNIU8vEmHWOJ33mOYp85zIHdqgq3ZCOR2+LtJkq0yD",
	"dEPXZlfHwGzPaGVS3MIIC29A9xQURJIcNEgT77iFPt5Ce6rDKHCjEJDmCWxmDfP9ZW3bdAiweylYrWqq",
	"6xNKvJ9GyWVc1MVS0IRl6oiLN6A7ONoVl3FhSzwrDKvBfyYnj6bTwVVykUj9R0phuisML7UKmeDON5Zl",
	"TgxOIUUCSj2k2KQhtG2R+nfdH/b/TqgUMqifiMPGeG3HTQK/gFbD00vIQlcXV+N0DIU76VG/SNE/z6Au",
	"68Npa0oOp8813UyVsYkd1y//cyDb2H8Xwq1HJiTLYpLc1Z9kR++z/dPGNZ7Fgu6ezEqZVc+1uAP+n+Mp",
	"O1oduu9jGNogwYuc7Lh1NHzXsyIjjB88lu1acdc84kv79rsz8U/a438HCTiqTIw7pWWW7Sat6j0kwO7B",
	"UYxvM3AOK3BILErtEKcgMoGsOXAo0ndCWZWqK/qh66Yfvy2GL25L2uhHVymH3zClZOjcS9rscZuY44/o",
	"P3vt/EWyPhzy02mROloMaFAND1t2D7y5T4Ycdsif33psCvgbxLdC3Blz6K15zMKVA/em2YcH9Mofhqvc",
	"8Rg54WQLOXDdfmEcbh/Z1jzHDnf2m7r3VnVd/REAAP//LLLvu/AVAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
