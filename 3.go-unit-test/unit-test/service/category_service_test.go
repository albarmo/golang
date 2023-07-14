package service

import (
	"fmt"
	"testing"
	"unit-test/repository"
	"unit-test/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T){
	category := entity.Category{
		Id: "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)
	categoryResult, err := categoryService.Get("2")
	
	fmt.Println(categoryResult)

	assert.Nil(t,err)
	assert.NotNil(t, categoryResult)
	assert.Equal(t, category.Id, categoryResult.Id) 
	assert.Equal(t, category.Name, categoryResult.Name) 
}