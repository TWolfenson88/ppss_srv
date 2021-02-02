package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"test_server/db"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// Создаем запрос с указанием нашего хендлера. Нам не нужно
	// указывать параметры, поэтому вторым аргументом передаем nil
	req, err := http.NewRequest("GET", "/app/get/GetAllStans", nil)
	if err != nil {
		t.Fatal(err)
	}

	db, err := db.InitDB(db.DBcfg)
	if err != nil {
		fmt.Println("DB ERROR")
		log.Fatal(db)
	}

	hnd := &Hndle{db}

	// Мы создаем ResponseRecorder(реализует интерфейс http.ResponseWriter)
	// и используем его для получения ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hnd.GetAllStans)

	// Наш хендлер соответствует интерфейсу http.Handler, а значит
	// мы можем использовать ServeHTTP и напрямую указать
	// Request и ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Проверяем код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем тело ответа
	expected := "About Page "
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
