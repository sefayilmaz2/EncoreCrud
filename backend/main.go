package main

import (
	"crud/crud/config"
	"crud/crud/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Veritabanını başlatma
	config.InitDB()

	// Router oluşturma
	r := mux.NewRouter()

	// CORS ayarları
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                       // Tüm kaynaklara izin veriyor
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // İzin verilen HTTP yöntemleri
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // İzin verilen başlıklar
		handlers.AllowCredentials(),                                                  // Kimlik bilgilerini izin ver
		handlers.MaxAge(3600),                                                        // CORS önbellek süresi
	)

	// Test için bir endpoint
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	// Ürün yönetimi endpointleri
	r.HandleFunc("/api/products", controllers.AddProduct).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", controllers.GetProductByID).Methods("GET")
	r.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	// CORS middleware'i route'lara uygula
	handler := corsOptions(r)

	// Sunucuyu başlat
	log.Println("Sunucu 5000 portunda çalışıyor...")
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
