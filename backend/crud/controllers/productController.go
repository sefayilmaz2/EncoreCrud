package controllers

import (
	"crud/crud/config"
	"crud/crud/models"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db := config.DB
	if db == nil {
		http.Error(w, "Veritabanı bağlantısı kurulamadı", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT Id, UrunAdi, Fiyat, Miktar FROM Urunler")
	if err != nil {
		http.Error(w, "Veritabanı sorgusu başarısız: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.UrunAdi, &product.Fiyat, &product.Miktar); err != nil {
			http.Error(w, "Sonuçlar alınırken hata oluştu: "+err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := mux.Vars(r)["id"]

	db := config.DB
	if db == nil {
		http.Error(w, "Veritabanı bağlantısı kurulamadı", http.StatusInternalServerError)
		return
	}

	var product models.Product
	query := "SELECT Id, UrunAdi, Fiyat, Miktar FROM Urunler WHERE Id = @Id"
	err := db.QueryRow(query, sql.Named("Id", id)).Scan(&product.ID, &product.UrunAdi, &product.Fiyat, &product.Miktar)
	if err != nil {
		http.Error(w, "Ürün bulunamadı: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.Unmarshal(body, &product); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	db := config.DB
	if db == nil {
		http.Error(w, "Veritabanı bağlantısı kurulamadı", http.StatusInternalServerError)
		return
	}

	var newID int
	query := `
		INSERT INTO Urunler (UrunAdi, Fiyat, Miktar) 
		OUTPUT INSERTED.Id 
		VALUES (@UrunAdi, @Fiyat, @Miktar)
	`

	err = db.QueryRow(
		query,
		sql.Named("UrunAdi", product.UrunAdi),
		sql.Named("Fiyat", product.Fiyat),
		sql.Named("Miktar", product.Miktar),
	).Scan(&newID)

	if err != nil {
		http.Error(w, "Veritabanına veri eklenirken hata oluştu: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message": "Eklendi",
		"id":      newID,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := mux.Vars(r)["id"]
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var product models.Product
	if err := json.Unmarshal(body, &product); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	db := config.DB
	if db == nil {
		http.Error(w, "Veritabanı bağlantısı kurulamadı", http.StatusInternalServerError)
		return
	}

	query := `
		UPDATE Urunler 
		SET UrunAdi = @UrunAdi, Fiyat = @Fiyat, Miktar = @Miktar 
		WHERE Id = @Id
	`
	result, err := db.Exec(query,
		sql.Named("UrunAdi", product.UrunAdi),
		sql.Named("Fiyat", product.Fiyat),
		sql.Named("Miktar", product.Miktar),
		sql.Named("Id", id),
	)
	if err != nil {
		http.Error(w, "Güncelleme sırasında hata oluştu: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":      "Güncellendi",
		"rowsAffected": rowsAffected,
	})
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := mux.Vars(r)["id"]

	db := config.DB
	if db == nil {
		http.Error(w, "Veritabanı bağlantısı kurulamadı", http.StatusInternalServerError)
		return
	}

	query := "DELETE FROM Urunler WHERE Id = @Id"
	result, err := db.Exec(query, sql.Named("Id", id))
	if err != nil {
		http.Error(w, "Silme sırasında hata oluştu: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":      "Silindi",
		"rowsAffected": rowsAffected,
	})
}
