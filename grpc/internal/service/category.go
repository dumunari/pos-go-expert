package service

import (
	"context"
	"grpc/internal/database"
	"grpc/internal/pb"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CategoryService implements the gRPC service for managing categories.
type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

// NewCategoryService initializes a new CategoryService with the provided database.
// It returns a pointer to the CategoryService instance.
func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

// CreateCategory creates a new category in the database.
// It returns the created category or an error if the creation fails.
func (c *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create category: %v", err)
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, _ *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list categories: %v", err)
	}

	var categoryList pb.CategoryList
	for _, category := range categories {
		categoryList.Categories = append(categoryList.Categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &categoryList, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(req.Id)
	if err != nil {
		if err == database.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "category not found: %v", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "failed to get category: %v", err)
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	// Read categories from the stream until EOF is reached.
	for {
		// Receive a category from the stream.
		category, err := stream.Recv()
		if err == io.EOF {
			// If EOF is reached, send the accumulated categories and close the stream.
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		// Validate the received category.
		// If the category name is empty, return an error.
		if category.Name == "" {
			return status.Errorf(codes.InvalidArgument, "category name cannot be empty")
		}
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			// If an error occurs during creation, return the error.
			return err
		}

		// Append the created category to the list.
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

// This method allows bidirectional streaming for category creation.
// It receives categories from the client, creates them in the database,
// and sends the created categories back to the client.
// It continues to receive and send categories until the stream is closed.
func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	// Loop to receive categories from the stream.
	for {
		// Receive a category from the stream.
		category, err := stream.Recv()
		// If EOF is reached, return nil to indicate the stream is closed.
		// This allows the server to gracefully handle the end of the stream.
		if err == io.EOF {
			return nil
		}
		// If an error occurs while receiving, return the error.
		if err != nil {
			return err
		}

		// Validate the received category.
		// If the category name is empty, return an error.
		if category.Name == "" {
			return status.Errorf(codes.InvalidArgument, "category name cannot be empty")
		}
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		// Send the created category back to the client.
		// If an error occurs while sending, return the error.
		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return err
		}
	}
}
