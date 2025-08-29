package product

import (
	"github.com/Vozhlak/golang-advanced/4-order-api/pkg/req"
	"github.com/Vozhlak/golang-advanced/4-order-api/pkg/res"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type HandlerProductDeps struct {
	ProductRepository *RepositoryProduct
}

type HandlerProduct struct {
	ProductRepository *RepositoryProduct
}

func NewHandlerProduct(router *http.ServeMux, deps HandlerProductDeps) {
	handler := &HandlerProduct{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /api/product", handler.Create)
	router.HandleFunc("GET /api/product/{id}", handler.GetProductById)
	router.HandleFunc("PATCH /api/product/{id}", handler.UpdateProduct)
	router.HandleFunc("DELETE /api/product/{id}", handler.DeleteProduct)
}

func (h *HandlerProduct) Create(w http.ResponseWriter, r *http.Request) {
	body, err := req.HandleBody[RequestProduct](&w, r)
	if err != nil {
		return
	}

	product := &Product{
		Name:        body.Name,
		Description: body.Description,
		Images:      body.Images,
		Price:       body.Price,
	}

	createdProduct, err := h.ProductRepository.Create(product)

	res.Json(w, createdProduct, http.StatusCreated)

}

func (h *HandlerProduct) GetProductById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.ProductRepository.GetProductById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res.Json(w, product, http.StatusOK)
}

func (h *HandlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := req.HandleBody[RequestUpdateProduct](&w, r)
	if err != nil {
		return
	}

	idString := r.PathValue("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.ProductRepository.GetProductById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	updatedProduct, err := h.ProductRepository.Update(&Product{
		Model:       gorm.Model{ID: uint(id)},
		Name:        body.Name,
		Description: body.Description,
		Images:      body.Images,
		Price:       body.Price,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Json(w, updatedProduct, http.StatusOK)
}

func (h *HandlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.ProductRepository.GetProductById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err = h.ProductRepository.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Json(w, nil, http.StatusOK)
}
