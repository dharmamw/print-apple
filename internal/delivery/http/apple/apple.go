package apple

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	appleEntity "print-apple/internal/entity/apple"
	"print-apple/pkg/response"
)

// IAppleSvc is an interface to User Service
type IAppleSvc interface {
	GetPrintApple(ctx context.Context) ([]appleEntity.Apple, error)
	GetPrintAppleStorage(ctx context.Context) ([]appleEntity.Apple, error)
	DeleteAndUpdateStorage(ctx context.Context, TransFH string) error
	GetPrintPageTemp(ctx context.Context, page int, length int) (map[string]interface{}, error)
	GetPrintPageFinal(ctx context.Context, page int, length int) (map[string]interface{}, error)
	GetByTransFHTemp(ctx context.Context, TransFH string) ([]appleEntity.Apple, error)
	GetByTransFHFinal(ctx context.Context, TransFH string) ([]appleEntity.Apple, error)
	GetByTglTransfTemp(ctx context.Context, TglTransf0 string, TglTransf1 string) ([]appleEntity.Apple, error)
	GetByTglTransfFinal(ctx context.Context, TglTransf0 string, TglTransf1 string) ([]appleEntity.Apple, error)
	//	GetComplexPageFinal(ctx context.Context, page int, length int, sortBy string) ([]appleEntity.Apple, error)
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
		apple    appleEntity.Apple
		page     int
		length   int
		// sortBy   string
	)
	// Make new response object
	resp = &response.Response{}
	body, _ := ioutil.ReadAll(r.Body)
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
		case "printAll":
			result, err = h.appleSvc.GetPrintApple(context.Background())
		case "reprintAll":
			result, err = h.appleSvc.GetPrintAppleStorage(context.Background())
		case "printByID":
			result, err = h.appleSvc.GetByTransFHTemp(context.Background(), r.FormValue("codeID")) // 3 digit pertama
		case "reprintByID":
			result, err = h.appleSvc.GetByTransFHFinal(context.Background(), r.FormValue("codeID")) // 3 digit pertama
		case "printByTgl":
			result, err = h.appleSvc.GetByTglTransfTemp(context.Background(), r.FormValue("start"), r.FormValue("end"))
		case "reprintByTgl":
			result, err = h.appleSvc.GetByTglTransfFinal(context.Background(), r.FormValue("start"), r.FormValue("end"))
		case "printPage":
			page, err = strconv.Atoi(r.FormValue("page"))
			length, err = strconv.Atoi(r.FormValue("length"))
			result, err = h.appleSvc.GetPrintPageTemp(context.Background(), page, length)
		case "reprintPage":
			page, err = strconv.Atoi(r.FormValue("page"))
			length, err = strconv.Atoi(r.FormValue("length"))
			result, err = h.appleSvc.GetPrintPageFinal(context.Background(), page, length)
			// case "getComplexPageFinal":
			// 	page, err = strconv.Atoi(r.FormValue("page"))
			// 	length, err = strconv.Atoi(r.FormValue("length"))
			// 	sortBy, err = r.FormValue("sort")
			// 	result, err = h.appleSvc.GetComplexPageFinal(context.Background(), page, length, sortBy)
			// 	log.Println(sortBy)
		}

	case http.MethodPut:
		json.Unmarshal(body, &apple)
		var _type string
		if _, putOK := r.URL.Query()["put"]; putOK {
			_type = r.FormValue("put")
		}
		switch _type {
		case "printSuccess":
			err = h.appleSvc.DeleteAndUpdateStorage(context.Background(), r.FormValue("ID"))
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
