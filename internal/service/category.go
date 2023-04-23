package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/georgejr3211/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categories []*pb.Category
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) CreateCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.Category, error) {
	categoryResponse := pb.Category{
		Id:          in.Name,
		Name:        in.Name,
		Description: in.Description,
	}

	s.categories = append(s.categories, &categoryResponse)

	return &categoryResponse, nil
}

func (s *CategoryService) ListCategories(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	return &pb.CategoryList{
		Categories: s.categories,
	}, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	var category *pb.Category
	for _, c := range s.categories {
		if c.Id == in.Id {
			category = c
			break
		}
	}

	if category == nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		fmt.Println("added a new category", category.Name)
		c := &pb.Category{
			Id:          category.Name,
			Name:        category.Name,
			Description: category.Description,
		}

		s.categories = append(s.categories, c)

		categories.Categories = append(categories.Categories, c)
	}
}

func (s *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		c := &pb.Category{
			Id:          category.Name,
			Name:        category.Name,
			Description: category.Description,
		}

		s.categories = append(s.categories, c)

		if err = stream.Send(c); err != nil {
			return err
		}
	}
}
