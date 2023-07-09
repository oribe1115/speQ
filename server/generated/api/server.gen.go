// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
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
	"github.com/labstack/echo/v4"
)

// ContestInfo defines model for ContestInfo.
type ContestInfo struct {
	Description        *string    `json:"description,omitempty"`
	EndTime            *time.Time `json:"end_time,omitempty"`
	ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty"`
	StartTime          *time.Time `json:"start_time,omitempty"`
	Title              *string    `json:"title,omitempty"`
	VotingFreezeTime   *time.Time `json:"voting_freeze_time,omitempty"`
}

// ProblemInfo defines model for ProblemInfo.
type ProblemInfo struct {
	Description *string `json:"description,omitempty"`

	// 表示順の指定にも使用
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// ProblemSolvedInfo defines model for ProblemSolvedInfo.
type ProblemSolvedInfo struct {
	ContestantId TraPId `json:"contestantId"`
	ProblemId    int    `json:"problemId"`
}

// ScoreInfo defines model for ScoreInfo.
type ScoreInfo struct {
	ContestantId TraPId  `json:"contestantId"`
	Score        float32 `json:"score"`
}

// UserInfo defines model for UserInfo.
type UserInfo struct {
	IsAdmin bool   `json:"isAdmin"`
	IsRoot  bool   `json:"isRoot"`
	TraPId  TraPId `json:"traPId"`
}

// TraPId defines model for traPId.
type TraPId = string

// PutContestInfoJSONBody defines parameters for PutContestInfo.
type PutContestInfoJSONBody = ContestInfo

// PutProblemsJSONBody defines parameters for PutProblems.
type PutProblemsJSONBody = []ProblemInfo

// UnmarkProblemAsSolvedJSONBody defines parameters for UnmarkProblemAsSolved.
type UnmarkProblemAsSolvedJSONBody = ProblemSolvedInfo

// MarkProblemAsSolvedJSONBody defines parameters for MarkProblemAsSolved.
type MarkProblemAsSolvedJSONBody = ProblemSolvedInfo

// PostScoreJSONBody defines parameters for PostScore.
type PostScoreJSONBody = ScoreInfo

// PutAdminUsersJSONBody defines parameters for PutAdminUsers.
type PutAdminUsersJSONBody = []TraPId

// PutContestantsJSONBody defines parameters for PutContestants.
type PutContestantsJSONBody = []TraPId

// GetContestantsJSONBody defines parameters for GetContestants.
type GetContestantsJSONBody = []TraPId

// PostVoteJSONBody defines parameters for PostVote.
type PostVoteJSONBody = TraPId

// PutContestInfoJSONRequestBody defines body for PutContestInfo for application/json ContentType.
type PutContestInfoJSONRequestBody = PutContestInfoJSONBody

// PutProblemsJSONRequestBody defines body for PutProblems for application/json ContentType.
type PutProblemsJSONRequestBody = PutProblemsJSONBody

// UnmarkProblemAsSolvedJSONRequestBody defines body for UnmarkProblemAsSolved for application/json ContentType.
type UnmarkProblemAsSolvedJSONRequestBody = UnmarkProblemAsSolvedJSONBody

// MarkProblemAsSolvedJSONRequestBody defines body for MarkProblemAsSolved for application/json ContentType.
type MarkProblemAsSolvedJSONRequestBody = MarkProblemAsSolvedJSONBody

// PostScoreJSONRequestBody defines body for PostScore for application/json ContentType.
type PostScoreJSONRequestBody = PostScoreJSONBody

// PutAdminUsersJSONRequestBody defines body for PutAdminUsers for application/json ContentType.
type PutAdminUsersJSONRequestBody = PutAdminUsersJSONBody

// PutContestantsJSONRequestBody defines body for PutContestants for application/json ContentType.
type PutContestantsJSONRequestBody = PutContestantsJSONBody

// GetContestantsJSONRequestBody defines body for GetContestants for application/json ContentType.
type GetContestantsJSONRequestBody = GetContestantsJSONBody

// PostVoteJSONRequestBody defines body for PostVote for application/json ContentType.
type PostVoteJSONRequestBody = PostVoteJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Put Contest Info
	// (PUT /admin/info)
	PutContestInfo(ctx echo.Context) error
	// Put Problems
	// (PUT /admin/problems)
	PutProblems(ctx echo.Context) error
	// Unmark Problem as Solved
	// (DELETE /admin/problems/solved)
	UnmarkProblemAsSolved(ctx echo.Context) error
	// Mark Problems as Solved
	// (POST /admin/problems/solved)
	MarkProblemAsSolved(ctx echo.Context) error
	// Post New Score
	// (POST /admin/scores)
	PostScore(ctx echo.Context) error
	// Put Admin Users
	// (PUT /admin/users/admin)
	PutAdminUsers(ctx echo.Context) error
	// Put Contestant Users
	// (PUT /admin/users/contestant)
	PutContestants(ctx echo.Context) error
	// Get Contest Info
	// (GET /info)
	GetContestInfo(ctx echo.Context) error
	// Get Admin Users
	// (GET /users/admin)
	GetAdmins(ctx echo.Context) error
	// Get Contestant Users
	// (GET /users/contestant)
	GetContestants(ctx echo.Context) error
	// Get My Information
	// (GET /users/me)
	GetMe(ctx echo.Context) error
	// Get Root Users
	// (GET /users/root)
	GetRootUsers(ctx echo.Context) error
	// Get Current Vote
	// (GET /vote)
	GetVote(ctx echo.Context) error
	// Post New Vote
	// (POST /vote)
	PostVote(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PutContestInfo converts echo context to params.
func (w *ServerInterfaceWrapper) PutContestInfo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutContestInfo(ctx)
	return err
}

// PutProblems converts echo context to params.
func (w *ServerInterfaceWrapper) PutProblems(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutProblems(ctx)
	return err
}

// UnmarkProblemAsSolved converts echo context to params.
func (w *ServerInterfaceWrapper) UnmarkProblemAsSolved(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UnmarkProblemAsSolved(ctx)
	return err
}

// MarkProblemAsSolved converts echo context to params.
func (w *ServerInterfaceWrapper) MarkProblemAsSolved(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.MarkProblemAsSolved(ctx)
	return err
}

// PostScore converts echo context to params.
func (w *ServerInterfaceWrapper) PostScore(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostScore(ctx)
	return err
}

// PutAdminUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PutAdminUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutAdminUsers(ctx)
	return err
}

// PutContestants converts echo context to params.
func (w *ServerInterfaceWrapper) PutContestants(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutContestants(ctx)
	return err
}

// GetContestInfo converts echo context to params.
func (w *ServerInterfaceWrapper) GetContestInfo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetContestInfo(ctx)
	return err
}

// GetAdmins converts echo context to params.
func (w *ServerInterfaceWrapper) GetAdmins(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAdmins(ctx)
	return err
}

// GetContestants converts echo context to params.
func (w *ServerInterfaceWrapper) GetContestants(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetContestants(ctx)
	return err
}

// GetMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetMe(ctx)
	return err
}

// GetRootUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetRootUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRootUsers(ctx)
	return err
}

// GetVote converts echo context to params.
func (w *ServerInterfaceWrapper) GetVote(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVote(ctx)
	return err
}

// PostVote converts echo context to params.
func (w *ServerInterfaceWrapper) PostVote(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostVote(ctx)
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

	router.PUT(baseURL+"/admin/info", wrapper.PutContestInfo)
	router.PUT(baseURL+"/admin/problems", wrapper.PutProblems)
	router.DELETE(baseURL+"/admin/problems/solved", wrapper.UnmarkProblemAsSolved)
	router.POST(baseURL+"/admin/problems/solved", wrapper.MarkProblemAsSolved)
	router.POST(baseURL+"/admin/scores", wrapper.PostScore)
	router.PUT(baseURL+"/admin/users/admin", wrapper.PutAdminUsers)
	router.PUT(baseURL+"/admin/users/contestant", wrapper.PutContestants)
	router.GET(baseURL+"/info", wrapper.GetContestInfo)
	router.GET(baseURL+"/users/admin", wrapper.GetAdmins)
	router.GET(baseURL+"/users/contestant", wrapper.GetContestants)
	router.GET(baseURL+"/users/me", wrapper.GetMe)
	router.GET(baseURL+"/users/root", wrapper.GetRootUsers)
	router.GET(baseURL+"/vote", wrapper.GetVote)
	router.POST(baseURL+"/vote", wrapper.PostVote)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
