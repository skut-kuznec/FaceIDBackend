package staffapp

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	mock_staffapp "github.com/smart48ru/FaceIDApp/internal/app/staffapp/mocks"
	"github.com/smart48ru/FaceIDApp/internal/domain"
	"reflect"
	"testing"
)

func TestAddEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_staffapp.NewMockStaffRepo(ctrl)

	// Положительный тест: ожидаем, что метод Save будет вызван с правильными параметрами и вернет результат
	expectedMockEmployee := domain.Employee{
		Name:    "John Doe",
		PhotoID: 2,
	}
	// Позитивный тест: ожидаем, что метод Save вернет id нового пользователя
	mockRepo.EXPECT().Save(gomock.Any(), expectedMockEmployee).Return(uint64(1), nil)

	app := App{repo: mockRepo}

	id, err := app.AddEmployee(context.Background(), expectedMockEmployee)
	expectedID := uint64(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if id != expectedID {
		t.Errorf("Expected id %v but got %v", 1, expectedID)
	}

	// Негативный тест: ожидаем, что метод Save вернет ошибку
	mockRepo.EXPECT().Save(gomock.Any(), expectedMockEmployee).Return(uint64(0), errors.New("employee not found"))

	id, err = app.AddEmployee(context.Background(), expectedMockEmployee)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	expectedErrMsg := "employee not found"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s' but got '%s'", expectedErrMsg, err.Error())
	}
}

func TestGetEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_staffapp.NewMockStaffRepo(ctrl)

	// Положительный тест: ожидаем, что метод Get будет вызван с правильными параметрами и вернет результат
	mockEmployee := domain.Employee{
		ID:      1,
		Name:    "John Doe",
		PhotoID: 2,
	}
	mockRepo.EXPECT().Get(gomock.Any(), uint64(1)).Return(mockEmployee, nil)

	app := App{repo: mockRepo}

	employee, err := app.GetEmployee(context.Background(), 1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(employee, mockEmployee) {
		t.Errorf("Expected employee %v but got %v", mockEmployee, employee)
	}

	// Негативный тест: ожидаем, что метод Get вернет ошибку
	mockRepo.EXPECT().Get(gomock.Any(), uint64(2)).Return(domain.Employee{}, errors.New("employee not found"))

	employee, err = app.GetEmployee(context.Background(), 2)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	expectedErrMsg := "employee not found"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s' but got '%s'", expectedErrMsg, err.Error())
	}
}

func TestApp_UpdateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_staffapp.NewMockStaffRepo(ctrl)

	// Положительный тест: ожидаем, что метод Get будет вызван с правильными параметрами и вернет результат
	mockEmployee := domain.Employee{
		ID:      1,
		Name:    "John Doe",
		PhotoID: 2,
	}

	mockRepo.EXPECT().Update(gomock.Any(), mockEmployee).Return(mockEmployee, nil)
	app := App{repo: mockRepo}

	employee, err := app.UpdateEmployee(context.Background(), mockEmployee)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(employee, mockEmployee) {
		t.Errorf("Expected employee %v but got %v", mockEmployee, employee)
	}

	// Негативный тест: ожидаем, что метод Get вернет ошибку
	mockRepo.EXPECT().Update(gomock.Any(), mockEmployee).Return(domain.Employee{}, errors.New("employee not found"))

	employee, err = app.UpdateEmployee(context.Background(), mockEmployee)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	expectedErrMsg := "employee not found"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s' but got '%s'", expectedErrMsg, err.Error())
	}

}

func TestApp_DeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_staffapp.NewMockStaffRepo(ctrl)

	// Положительный тест: ожидаем, что метод Delete будет вызван с правильными параметрами и не вернет ошибку
	id := uint64(1)

	mockRepo.EXPECT().Delete(gomock.Any(), id).Return(nil)
	app := App{repo: mockRepo}

	err := app.DeleteEmployee(context.Background(), id)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Негативный тест: ожидаем, что метод Get вернет ошибку
	mockRepo.EXPECT().Delete(gomock.Any(), id).Return(errors.New("employee not found"))

	err = app.DeleteEmployee(context.Background(), id)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	expectedErrMsg := "employee not found"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s' but got '%s'", expectedErrMsg, err.Error())
	}

}

func TestApp_ListEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_staffapp.NewMockStaffRepo(ctrl)

	// Положительный тест: ожидаем, что метод List будет вызван с правильными параметрами и не вернет ошибку
	expectedList := []domain.Employee{{ID: 1, Name: "John", PhotoID: 2}, {ID: 2, Name: "Jane", PhotoID: 12}}

	mockRepo.EXPECT().List(gomock.Any()).Return(expectedList, nil)
	app := App{repo: mockRepo}

	list, err := app.ListEmployee(context.Background())

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(expectedList, list) {
		t.Errorf("Expected employee %v but got %v", expectedList, list)
	}

	// Негативный тест: ожидаем, что метод List вернет ошибку
	mockRepo.EXPECT().List(gomock.Any()).Return([]domain.Employee{}, errors.New("employee not found"))

	list, err = app.ListEmployee(context.Background())

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	expectedErrMsg := "employee not found"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s' but got '%s'", expectedErrMsg, err.Error())
	}
}
