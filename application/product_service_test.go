package application_test

import (
	"errors"
	"github.com/souluanf/hexagonal-arch-golang/application"
	mockapplication "github.com/souluanf/hexagonal-arch-golang/application/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)
	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_GetWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	expectedError := errors.New("erro ao buscar produto")
	persistence.EXPECT().Get(gomock.Any()).Return(nil, expectedError).Times(1)
	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Get("id_inexistente")
	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)
	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("abc", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create_InvalidProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{Persistence: persistence}
	_, err := service.Create("product 1", -10)

	require.Error(t, err)
}

func TestProductService_Create_PersistenceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	expectedError := errors.New("error on save product")

	persistence.EXPECT().Save(gomock.Any()).Return(nil, expectedError).Times(1)
	service := application.ProductService{Persistence: persistence}
	_, err := service.Create("product 1", 10)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)

	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().Disable().Return(nil).AnyTimes()

	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Enable_ProductEnableError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)
	expectedError := errors.New("error on enable product")

	product.EXPECT().Enable().Return(expectedError).Times(1)

	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.NotNil(t, result)

	_, ok := result.(*application.Product)
	require.True(t, ok, "the result should be an instance of *Product")
}

func TestProductService_Enable_PersistenceSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).Times(1)

	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	expectedError := errors.New("erro ao salvar produto")

	persistence.EXPECT().Save(gomock.Any()).Return(nil, expectedError).Times(1)

	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.NotNil(t, result)
}

func TestProductService_Disable_ProductDisableError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)
	expectedError := errors.New("erro ao desabilitar produto")

	product.EXPECT().Disable().Return(expectedError).Times(1)

	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	service := application.ProductService{Persistence: persistence}

	result, err := service.Disable(product)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.NotNil(t, result)
}

func TestProductService_Disable_PersistenceSaveError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapplication.NewMockProductInterface(ctrl)

	product.EXPECT().Disable().Return(nil).Times(1)

	persistence := mockapplication.NewMockProductPersistenceInterface(ctrl)
	expectedError := errors.New("erro ao salvar produto desabilitado")

	persistence.EXPECT().Save(gomock.Any()).Return(nil, expectedError).Times(1)

	service := application.ProductService{Persistence: persistence}

	result, err := service.Disable(product)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.NotNil(t, result)
}
