package staffapp

import (
	"context"
	"fmt"
	"github.com/smart48ru/FaceIDApp/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestApp_AddEmployee(t *testing.T) {
	// Создаем фейковый репозиторий с помощью пакета github.com/stretchr/testify/mock
	mockRepo := new(MockStaffRepo)

	// Создаем экземпляр приложения, передавая фейковый репозиторий
	app := New(mockRepo)

	// Создаем фейковый контекст
	ctx := context.Background()

	// Создаем фейкового сотрудника
	employee := domain.Employee{
		Name:    "John",
		PhotoID: 10,
	}

	// Задаем ожидаемый идентификатор сотрудника, который вернется из фейкового репозитория
	expectedID := uint64(12)

	// Задаем, что при вызове метода Save у фейкового репозитория с переданным контекстом и сотрудником
	// должен возвращаться ожидаемый идентификатор и отсутствие ошибки
	mockRepo.On("Save", ctx, employee).Return(expectedID, nil)

	// Вызываем метод AddEmployee приложения с переданным контекстом и сотрудником
	id, err := app.AddEmployee(ctx, employee)

	// Проверяем, что фейковый репозиторий был вызван с переданными аргументами
	mockRepo.AssertCalled(t, "Save", ctx, employee)

	// Проверяем, что метод AddEmployee вернул ожидаемый идентификатор и отсутствие ошибки
	fmt.Println(expectedID, id)
	assert.Equal(t, expectedID, id)
	assert.Nil(t, err)
}

func TestApp_GetEmployee(t *testing.T) {
	// Создаем мок репозитория
	mockRepo := new(MockStaffRepo)
	// Создаем экземпляр нашего приложения, передавая мок репозитория
	app := New(mockRepo)
	ctx := context.Background()

	// Определяем ожидаемый результат
	expectedEmployee := domain.Employee{
		ID:      2,
		Name:    "John Doe",
		PhotoID: 0,
	}

	// Задаем идентификатор сотрудника для запроса
	ID := uint64(2)

	// Задаем, что при вызове метода Get у фейкового репозитория с переданным контекстом и ID сотрудника
	// должен возвращаться ожидаемая структура и отсутствие ошибки
	mockRepo.On("Get", ctx, expectedEmployee.ID).Return(expectedEmployee, nil)

	// Проверяем, что фейковый репозиторий был вызван с переданными аргументами
	employee, err := app.GetEmployee(ctx, ID)

	// Проверяем, что метод GetEmployee вернул ожидаемую структуру и отсутствие ошибки
	assert.NoError(t, err)
	assert.Equal(t, expectedEmployee, employee)
	mockRepo.AssertCalled(t, "Get", ctx, expectedEmployee.ID)
}

func TestApp_UpdateEmployee(t *testing.T) {
	// Создаем мок репозитория
	mockRepo := new(MockStaffRepo)

	// Создаем экземпляр нашего приложения, передавая мок репозитория
	app := New(mockRepo)
	ctx := context.Background()

	// Создаем тестовый сотрудник
	employee := domain.Employee{
		ID:      1,
		Name:    "John Doe",
		PhotoID: 0,
	}

	// Определяем ожидаемый результат
	expectedEmployee := domain.Employee{
		ID:      1,
		Name:    "John Doe",
		PhotoID: 0,
	}

	// Указываем, что при вызове метода Update репозитория с параметром employee вернуть expectedEmployee и nil
	mockRepo.On("Update", mock.Anything, employee).Return(expectedEmployee, nil)

	// Вызываем метод UpdateEmployee нашего приложения
	actualEmployee, err := app.UpdateEmployee(ctx, employee)

	// Проверяем, что метод Update нашего мок репозитория был вызван с параметром employee
	mockRepo.AssertCalled(t, "Update", mock.Anything, employee)

	// Проверяем, что метод UpdateEmployee вернул ожидаемый результат и ошибку равную nil
	assert.Equal(t, expectedEmployee, actualEmployee)
	assert.Nil(t, err)
}

func TestApp_DeleteEmployee(t *testing.T) {
	// Создаем мок репозитория
	mockRepo := new(MockStaffRepo)

	// Создаем экземпляр приложения, передавая фейковый репозиторий
	app := New(mockRepo)

	// Создаем фейковый контекст
	ctx := context.Background()

	// Задаем идентификатор сотрудника, который будет удален
	employeeID := uint64(123)

	// Задаем, что при вызове метода Delete у фейкового репозитория с переданным контекстом и идентификатором
	// должен отсутствовать возвращаемый результат
	mockRepo.On("Delete", ctx, employeeID).Return(nil)

	// Вызываем метод DeleteEmployee приложения с переданным контекстом и идентификатором сотрудника
	err := app.DeleteEmployee(ctx, employeeID)

	// Проверяем, что фейковый репозиторий был вызван с переданными аргументами
	mockRepo.AssertCalled(t, "Delete", ctx, employeeID)

	// Проверяем, что метод DeleteEmployee вернул отсутствие ошибки
	assert.Nil(t, err)
}

func TestApp_ListEmployee(t *testing.T) {
	// Создаем фейковый список сотрудников
	employees := []domain.Employee{
		{ID: 1, Name: "John Doe", PhotoID: 1},
		{ID: 2, Name: "Jon Brown", PhotoID: 4},
		{ID: 3, Name: "Bob Smith", PhotoID: 12},
	}

	// Создаем мок репозитория
	mockRepo := new(MockStaffRepo)

	// Создаем экземпляр приложения, передавая фейковый репозиторий
	app := New(mockRepo)

	// Создаем фейковый контекст
	ctx := context.Background()

	// Задаем, что при вызове метода List у фейкового репозитория с переданным контекстом
	// должен возвращаться фейковый список сотрудников
	mockRepo.On("List", ctx).Return(employees, nil)

	// Вызываем метод ListEmployee приложения с переданным контекстом
	list, err := app.ListEmployee(ctx)

	// Проверяем, что фейковый репозиторий был вызван с переданным контекстом
	mockRepo.AssertCalled(t, "List", ctx)

	// Проверяем, что метод ListEmployee вернул ожидаемый список сотрудников
	assert.Equal(t, employees, list)

	// Проверяем, что метод ListEmployee вернул отсутствие ошибки
	assert.Nil(t, err)
}
