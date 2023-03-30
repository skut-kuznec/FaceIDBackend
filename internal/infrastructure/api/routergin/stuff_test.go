package routergin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	// Создаем экземпляр приложения
	app := gin.Default()
	// Добавляем маршрут для метода GetAll
	router := RouterGin{}
	app.GET("/all", router.GetAllStaffEndPoint)

	// Создаем тестовый запрос
	req, err := http.NewRequest(http.MethodGet, "/all", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Создаем контекст выполнения теста
	w := httptest.NewRecorder()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Выполняем запрос и проверяем результат
	app.ServeHTTP(w, req.WithContext(ctx))
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}
}
