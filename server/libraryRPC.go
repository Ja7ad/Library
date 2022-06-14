package main

import (
	"context"
	"github.com/Ja7ad/transaction/proto/protoModel/library"
	bookMgr "github.com/Ja7ad/transaction/server/book"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LibraryServer struct {
	library.UnimplementedLibraryServiceServer
}

func (*LibraryServer) GetBooks(ctx context.Context, _ *emptypb.Empty) (*library.GetBooksResponse, error) {
	books, err := bookMgr.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
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
	bookID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	publisherID, err := primitive.ObjectIDFromHex(request.PublisherId)
	if err != nil {
		return nil, err
	}
	book, err := bookMgr.FindBook(ctx, request.Name, request.PublisherName, bookID, publisherID)
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
	book, err := bookMgr.AddBook(ctx, request.Name, request.Publisher)
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
	book, err := bookMgr.UpdateBook(ctx, bookID, request.Name, request.Publisher)
	if err != nil {
		return nil, err
	}

	return &library.Book{
		Id:          book.Id.Hex(),
		Name:        book.Name,
		PublisherId: book.PublisherId.Hex(),
	}, nil
}

func (*LibraryServer) DeleteBook(ctx context.Context, request *library.DeleteBookRequest) (*emptypb.Empty, error) {
	bookID, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	if err := bookMgr.DeleteBook(ctx, bookID); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
