package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"inventory-api/internal/db"
	"inventory-api/internal/utils"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		utils.RespondWithError(ctx, w, http.StatusRequestTimeout, "Request cancelled or timed out")
		return
	default:
	}

	var p db.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.service.AddProduct(ctx, &p); err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", p.ID),
	)
	logger.Info("Product created successfully")

	utils.RespondWithJSON(ctx, w, http.StatusCreated, p)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		utils.RespondWithError(ctx, w, http.StatusRequestTimeout, "Request cancelled or timed out")
		return
	default:
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p db.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.ID = id

	if err := h.service.UpdateProduct(ctx, &p); err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", p.ID),
	)
	logger.Info("Product updated successfully")

	utils.RespondWithJSON(ctx, w, http.StatusOK, p)
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		utils.RespondWithError(ctx, w, http.StatusRequestTimeout, "Request cancelled or timed out")
		return
	default:
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.service.GetProduct(ctx, id)
	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusNotFound, "Product not found")
		return
	}

	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", id),
	)
	logger.Info("Product retrieved successfully")

	utils.RespondWithJSON(ctx, w, http.StatusOK, product)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		utils.RespondWithError(ctx, w, http.StatusRequestTimeout, "Request cancelled or timed out")
		return
	default:
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(ctx, w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.service.DeleteProduct(ctx, id); err != nil {
		utils.RespondWithError(ctx, w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	logger := zap.L().With(
		zap.String("request_id", utils.GetRequestID(ctx)),
		zap.String("user_id", utils.GetUserID(ctx)),
		zap.Int("product_id", id),
	)
	logger.Info("Product deleted successfully")

	utils.RespondWithJSON(ctx, w, http.StatusOK, map[string]string{"message": "Product deleted"})
}
