package service

import (
	"context"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/Ja7ad/library/server/internal/managers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LibraryServer struct {
	library.UnimplementedBookServiceServer
}

func (*LibraryServer) GetBooks(ctx context.Context, _ *emptypb.Empty) (*library.GetBooksResponse, error) {
	books, err := managers.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	panic("test")
	pbBooks := &library.GetBooksResponse{}
	for _, book := range books {
		pbBook := &library.Book{
			Id:            book.Id.Hex(),
			Name:          book.Name,
			PublisherName: book.PublisherName,
		}
		if !book.UserId.IsZero() {
			pbBook.UserId = book.UserId.Hex()
		}
		pbBooks.Books = append(pbBooks.Books, pbBook)
	}

	return pbBooks, nil
}

func (*LibraryServer) FindBook(ctx context.Context, request *library.FindBookRequest) (*library.Book, error) {
	var (
		bookID primitive.ObjectID
		err    error
	)
	if len(request.Id) != 0 {
		bookID, err = primitive.ObjectIDFromHex(request.Id)
		if err != nil {
			return nil, err
		}
	}
	book, err := managers.FindBook(ctx, bookID)
	if err != nil {
		return nil, err
	}
	pbBook := &library.Book{
		Id:            book.Id.Hex(),
		Name:          book.Name,
		PublisherName: book.PublisherName,
	}
	if !book.UserId.IsZero() {
		pbBook.UserId = book.UserId.Hex()
	}
	return pbBook, nil
}

func (*LibraryServer) AddBook(ctx context.Context, request *library.AddBookRequest) (*library.Book, error) {
	book, err := managers.AddBook(ctx, request.Name, request.Publisher)
	if err != nil {
		return nil, err
	}
	return &library.Book{
		Id:          book.Id.Hex(),
		Name:        book.Name,
		PublisherId: book.PublisherId.Hex(),
	}, nil
}

func (*LibraryServer) UpdateBook(ctx context.Context, request *library.UpdateBookRequest) (*library.Book, error) {
	bookID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	book, err := managers.UpdateBook(ctx, bookID, request.Name, request.Publisher)
	if err != nil {
		return nil, err
	}

	return &library.Book{
		Id:          book.Id.Hex(),
		Name:        book.Name,
		PublisherId: book.PublisherId.Hex(),
	}, nil
}

func (*LibraryServer) ReserveBook(ctx context.Context, request *library.ReserveBookRequest) (*emptypb.Empty, error) {
	bookIDs := []primitive.ObjectID{}
	for _, b := range request.BookId {
		bookID, err := primitive.ObjectIDFromHex(b)
		if err != nil {
			return nil, err
		}
		bookIDs = append(bookIDs, bookID)
	}
	userID, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}

	if err := managers.ReserveBook(ctx, userID, bookIDs...); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (*LibraryServer) DeleteBook(ctx context.Context, request *library.DeleteBookRequest) (*emptypb.Empty, error) {
	bookID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	if err := managers.DeleteBook(ctx, bookID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
