package service

import (
	"github.com/paulodutra/golang-grpc/internal/database"
	"github.com/paulodutra/golang-grpc/internal/pb"
)

type CategoryService struct{
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.name, in.description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:			 category.ID,
		Name: 	     category.Name,
		Description: category.Description
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}