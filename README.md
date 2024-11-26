
# Ürün Senkronizasyon ve Yönetim Sistemi

Bu proje, farklı sistemler arasında ürün verilerinin senkronizasyonunu yönetmek için tasarlanmıştır. Yeni ürünlerin eklenmesi, mevcut ürünlerin güncellenmesi ve ürün detaylarının verimli bir şekilde alınmasını sağlar. Proje **Microsoft SQL Server** üzerinde klasik SELECT, INSERT, UPDATE, DELETE şeklinde çalışır.

---

## Özellikler
- Mevcut ürünlerin verilerinin güncellenmesi.
- Tüm ürünlerin veya belirli bir ürünün ID'ye göre detaylarının getirilmesi.
- Ürünlerin ID'ye göre silinmesi.
- SQL tabanlı yapı ile yüksek performans ve ölçeklenebilirlik.

---

## Gereksinimler
1. **Microsoft SQL Server**: Çalışan bir SQL Server kurulumu gereklidir.
2. **SQL Server Management Studio (SSMS)**: SQL sorgularını çalıştırmak ve yönetmek için gereklidir.
3. **Veritabanı Erişimi**: `INSERT`, `UPDATE`, `SELECT` ve `DELETE` işlemleri için gerekli izinlere sahip olmalısınız.

---

## Kurulum Adımları

### 1. Proje Deposu Klonlama
- Bu projeyi bilgisayarınıza klonlayın:
- ```bash git clone https://github.com/sefayilmaz2/EncoreCrud.git ```

---

### 2. Veritabanını Hazırlama
1. Yeni bir veritabanı oluşturun veya mevcut bir veritabanını kullanın.
2. `/backend/crud/models/productModels.go` dosyasına göre tablo oluşturun.

---

### 3. Veritabanı Bağlantısını Yapılandırma
- `/backend/crud/config/dbConfig.go` dosyasından kendi SQL bağlantı yolunuzu giriniz.

## Kullanım
### Yeni Ürün Ekleme
- `/backend/main.go` dosyasında belirlenen port ve API Url bilgisine göre istek atılır. 
- Default bilgilere göre http://localhost:5000/api/products URL adresine POST isteği ile birlikte Örnek Kod;
```
{
    "urunAdi":"TEST1",
    "fiyat":"1000",
    "miktar":"10"
}
```
data ile istek gönderilir.

### Ürünleri Getirme
- Default bilgilere göre http://localhost:5000/api/products URL adresine GET isteği gönderilir. Return olarak Models yapısı geri gönderilir.

### ID Değerine Göre Ürün Getirme
- Default bilgilere göre http://localhost:5000/api/products/{id} URL adresine GET isteği ile birlikte id değeri gönderilir. Return olarak Models yapısı geri gönderilir.

### Ürün Bilgilerini Güncelleme
- Default bilgilere göre http://localhost:5000/api/products/{id} URL adresine PUT isteği ile birlikte id değeri gönderilir. Ek ilave olarak örnek;
```
{
    "urunAdi":"TEST1",
    "fiyat":"1000",
    "miktar":"10"
}
```
json formatında data gönderilir.

### Ürün Silme
- Default bilgilere göre http://localhost:5000/api/products/{id} URL adresine DELETE isteği ile birlikte id değeri gönderilir.

---

# İletişim
- Sorularınız veya destek talepleriniz için sefa.yilmaz@c1soft.com adresine ulaşabilirsiniz.
