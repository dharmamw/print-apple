package apple

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	appleEntity "print-apple/internal/entity/apple"
	"print-apple/pkg/response"
)

// IAppleSvc is an interface to User Service
type IAppleSvc interface {
	GetAppleFromFireBase(ctx context.Context) ([]appleEntity.Apple, error)
}

type (
	// Handler ...
	Handler struct {
		appleSvc IAppleSvc
	}
)

// New for user domain handler initialization
func New(is IAppleSvc) *Handler {
	return &Handler{
		appleSvc: is,
	}
}

// AppleHandler will return user data
func (h *Handler) AppleHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp     *response.Response
		metadata interface{}
		result   interface{}
		err      error
		errRes   response.Error
	)
	// Make new response object
	resp = &response.Response{}
	// Defer will be run at the end after method finishes
	defer resp.RenderJSON(w, r)

	switch r.Method {
	// Check if request method is GET
	case http.MethodGet:
		// Ambil semua data user
		var _type string
		if _, getOK := r.URL.Query()["get"]; getOK {
			_type = r.FormValue("get")
		}
		switch _type {
		case "printappleall":
			result, err = h.appleSvc.GetAppleFromFireBase(context.Background())
		}
	default:
		err = errors.New("400")
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   201,
				Msg:    "Failed to process request due to server error",
				Status: true,
			}
		}

		// Logging
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.Error = errRes
		return
	}

	// Inserting data to response
	resp.Data = result
	resp.Metadata = metadata
	// Logging
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
