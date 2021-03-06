package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/asaskevich/govalidator.v8"

	"github.com/supergiant/supergiant/pkg/message"
	"github.com/supergiant/supergiant/pkg/sgerrors"
)

// Handler is a http controller for account entity
type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(r *mux.Router) {
	r.HandleFunc("/accounts", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/accounts", h.ListAll).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{accountName}", h.Get).Methods(http.MethodGet)
	r.HandleFunc("/accounts/{accountName}", h.Update).Methods(http.MethodPut)
	r.HandleFunc("/accounts/{accountName}", h.Delete).Methods(http.MethodDelete)
}

// Create register new cloud account
func (h *Handler) Create(rw http.ResponseWriter, r *http.Request) {
	account := new(CloudAccount)
	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		message.SendInvalidJSON(rw, err)
		return
	}

	ok, err := govalidator.ValidateStruct(account)
	if !ok {
		message.SendValidationFailed(rw, err)
		return
	}

	if err = h.service.Create(r.Context(), account); err != nil {
		logrus.Errorf("create account: %v", err)
		message.SendUnknownError(rw, err)
		return
	}
}

// ListAll retrieves all cloud accounts
func (h *Handler) ListAll(rw http.ResponseWriter, r *http.Request) {
	accounts, err := h.service.GetAll(r.Context())
	if err != nil {
		logrus.Errorf("accounts list all %v", err)
		message.SendUnknownError(rw, err)
		return
	}
	if err := json.NewEncoder(rw).Encode(accounts); err != nil {
		logrus.Errorf("accounts list all %v", err)
		message.SendUnknownError(rw, err)
		return
	}
}

// Get retrieves individual account by name
func (h *Handler) Get(rw http.ResponseWriter, r *http.Request) {
	accountName := mux.Vars(r)["accountName"]
	if accountName == "" {
		msg := message.New("account name can't be blank", "", sgerrors.CantChangeID, "")
		message.SendMessage(rw, msg, http.StatusBadRequest)
		return
	}
	account, err := h.service.Get(r.Context(), accountName)
	if err != nil {
		if sgerrors.IsNotFound(err) {
			message.SendNotFound(rw, "account", err)
			return
		}
		logrus.Errorf("account get %v", err)
		message.SendUnknownError(rw, err)
		return
	}

	if err := json.NewEncoder(rw).Encode(account); err != nil {
		logrus.Errorf("account get %v", err)
		message.SendUnknownError(rw, err)
		return
	}
}

// Update saves updated state of an cloud account, account name can't be changed
func (h *Handler) Update(rw http.ResponseWriter, r *http.Request) {
	account := new(CloudAccount)
	if err := json.NewDecoder(r.Body).Decode(account); err != nil {
		message.SendInvalidJSON(rw, err)
		return
	}

	ok, err := govalidator.ValidateStruct(account)
	if !ok {
		message.SendValidationFailed(rw, err)
		return
	}
	if err := h.service.Update(r.Context(), account); err != nil {
		logrus.Errorf("account update: %v", err)
		message.SendUnknownError(rw, err)
		return
	}
}

// Delete cloud account
func (h *Handler) Delete(rw http.ResponseWriter, r *http.Request) {
	accountName := mux.Vars(r)["accountName"]
	if accountName == "" {
		msg := message.New("account name can't be blank", "", sgerrors.CantChangeID, "")
		message.SendMessage(rw, msg, http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), accountName); err != nil {
		logrus.Errorf("account delete %v", err)
		message.SendUnknownError(rw, err)
		return
	}
}
