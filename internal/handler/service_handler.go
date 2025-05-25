package handler

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sonishivam10/service-catalog/internal/repository"
	"github.com/sonishivam10/service-catalog/internal/service"
	"github.com/sonishivam10/service-catalog/pkg/response"
)

type ServiceHandler struct {
	ServiceUsecase service.ServiceUsecase
}

func NewServiceHandler(r *mux.Router, serviceUsecase service.ServiceUsecase) {
	handler := &ServiceHandler{
		ServiceUsecase: serviceUsecase,
	}

	r.HandleFunc("/services", handler.ListServices).Methods("GET")
	r.HandleFunc("/services/{id}", handler.GetService).Methods("GET")
	r.HandleFunc("/services/{id}/versions", handler.GetVersions).Methods("GET")
}

// @Summary List services
// @Description Get all services with filtering, sorting, and pagination
// @Tags services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param search query string false "Search keyword"
// @Param sort_by query string false "Sort by field (name, created_at)"
// @Param sort_order query string false "asc or desc"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /services [get]
func (h *ServiceHandler) ListServices(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	params := repository.FilterParams{
		Search:    r.URL.Query().Get("search"),
		SortBy:    r.URL.Query().Get("sort_by"),
		SortOrder: r.URL.Query().Get("sort_order"),
		Page:      1,
		PageSize:  10,
	}

	if page, err := strconv.Atoi(r.URL.Query().Get("page")); err == nil && page > 0 {
		params.Page = page
	}
	if size, err := strconv.Atoi(r.URL.Query().Get("page_size")); err == nil && size > 0 {
		params.PageSize = size
	}

	services, err := h.ServiceUsecase.ListServices(r.Context(), params)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch services")
		return
	}

	pagination := &response.Pagination{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	response.JSON(w, http.StatusOK, services, pagination)
}

// @Summary Get a Service by ID
// @Description Returns a single service by its UUID.
// @Tags services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Service UUID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse "Invalid UUID format"
// @Failure 404 {object} response.APIResponse "Service not found"
// @Router /services/{id} [get]
func (h *ServiceHandler) GetService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid service ID")
		return
	}

	service, err := h.ServiceUsecase.GetService(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Service not found")
		return
	}

	response.JSON(w, http.StatusOK, service, nil)
}

// @Summary List Versions of a Service
// @Description Returns a list of all versions for a specific service.
// @Tags services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Service UUID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse "Invalid UUID format"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /services/{id}/versions [get]
func (h *ServiceHandler) GetVersions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid service ID")
		return
	}

	versions, err := h.ServiceUsecase.GetVersionsByServiceID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch versions")
		return
	}

	response.JSON(w, http.StatusOK, versions, nil)
}
