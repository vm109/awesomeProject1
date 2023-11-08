package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go_app/m/v2/pkg/domain/model"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Api struct {
	http.Handler
	Router *httprouter.Router
	Port   int
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func (a *Api) Initialize() {

	// ERROR HANDLERS
	a.Router.NotFound = http.HandlerFunc(a.NotFound)
	a.Router.MethodNotAllowed = http.HandlerFunc(a.MethodNotAllowed)
}

func (a *Api) NotFound(w http.ResponseWriter, r *http.Request) {
	a.WriteErrorResponse(r, w, model.EndpointNotFoundError())
}

func (a *Api) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	a.WriteJsonBodyResponse(r, w, model.ErrorResponse{Message: "method not allowed"}, http.StatusMethodNotAllowed)
}

func (a *Api) WriteErrorResponse(r *http.Request, w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	fmt.Println(err)
	response := model.ErrorResponse{
		Code: model.ErrInternalServerException,
	}

	if err, ok := err.(*model.Error); ok {
		switch err.Code {
		case model.ErrInvalidData:
			fallthrough
		case model.ErrInvalidParam:
			fallthrough
		case model.ErrAlreadyExists:
			fallthrough
		case model.ErrNotFound:
			status = http.StatusNotFound
			response.Code = err.Code
			response.Message = err.Msg
		case model.ErrNotImplemented:
			status = http.StatusNotImplemented
			response.Code = err.Code
			response.Message = err.Msg
		case model.ErrPreconditionFailed:
			status = http.StatusPreconditionFailed
			response.Code = err.Code
			response.Message = err.Msg
		case model.ErrRestrictedContent:
			status = http.StatusForbidden
			response.Code = err.Code
			response.Message = err.Msg
		case model.ErrInvalidBSON:
			status = http.StatusBadRequest
			response.Code = err.Code
			response.Message = err.Msg
		case model.ErrSizeExceeded:
			status = http.StatusBadRequest
			response.Code = err.Code
			response.Message = err.Msg
		}
	}

	if rawJson, err := json.Marshal(&response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(rawJson)
	}

	w.WriteHeader(status)
	w.Write(nil)
}
func (a *Api) WriteJsonBodyResponse(r *http.Request, w http.ResponseWriter, data interface{}, status int) {
	rawJson, err := json.Marshal(data)

	if err != nil {
		a.WriteErrorResponse(r, w, err)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(rawJson)
}
func (a *Api) RawHandler(method, path string, fn http.HandlerFunc) {
	a.Router.HandlerFunc(method, path, fn)
}

func (a *Api) HandlerWithMiddleware(method, path string, fn http.HandlerFunc) {
	// Assuming 'h' is your middleware instance
	h := a.Handler

	a.Router.Handler(method, path, h)
}

func (a *Api) Start() {
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", a.Port),
		Handler: a,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		fmt.Println("Shutting down server")
		server.Shutdown(context.Background())
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

var (
	oauthConfig *oauth2.Config
)

func NewTaskManagerRestApi(port int) *Api {
	api := &Api{
		Port:   port,
		Router: httprouter.New(),
	}

	api.Handler = middleware(api.Router)
	// HEALTH CHECKS
	api.RawHandler("GET", "/status", api.statushandler)
	api.RawHandler("POST", "/daily/single/create", api.createASingleDailyTask)

	api.RawHandler("GET", "/plan/:id", api.getPlanById)

	return api
}

var templates *template.Template

func uploadHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("/Users/manginav/awesomeProject1/go_app/pkg/templates/upload.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Fprintf(w, "Visit the URL for the auth dialog: %v", url)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	tok, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Unable to retrieve token from web: %v", err)
		return
	}
	fmt.Println(tok)

	// Create a Drive service
	srv, err := drive.NewService(context.Background(), option.WithTokenSource(oauthConfig.TokenSource(context.Background(), tok)))
	if err != nil {
		log.Fatalf("Unable to create Drive service: %v", err)
	}

	files, err := srv.Files.List().Do()
	fmt.Println(files)
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	for _, file := range files.Files {
		fmt.Printf("File Name: %s, File ID: %s\n", file.Name, file.Id)
	}

	fmt.Fprint(w, "Authorization complete! You can close this tab now.")
}
