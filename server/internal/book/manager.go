package book

import (
	"context"
	"github.com/Ja7ad/library/server/global"
	models2 "github.com/Ja7ad/library/server/internal/book/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(ctx context.Context) ([]*models2.Book, error) {
	books, err := models2.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func FindBook(ctx context.Context, name, publisherName string, bookID, publisherID primitive.ObjectID) (*models2.Book, error) {
	book, err := models2.FindBook(ctx, name, publisherName, bookID, publisherID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func AddBook(ctx context.Context, name, publisherName string) (*models2.Book, error) {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	publisher, err := models2.GetPublisherByName(sessCtx, publisherName)
	if err != nil {
		if publisher, err = addPublisher(sessCtx, publisherName); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	book := &models2.Book{
		Id:          primitive.NewObjectID(),
		Name:        name,
		PublisherId: publisher.Id,
	}

	if err := book.Insert(sessCtx); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return nil, err
	}

	return book, nil
}

func UpdateBook(ctx context.Context, bookID primitive.ObjectID, name, publisherName string) (*models2.Book, error) {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	publisher, err := models2.GetPublisherByName(sessCtx, publisherName)
	if err != nil {
		if publisher, err = addPublisher(sessCtx, publisherName); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	book, err := models2.GetBookByID(sessCtx, bookID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	book.Name = name
	book.PublisherId = publisher.Id

	if err := book.Update(sessCtx); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(ctx context.Context, bookID primitive.ObjectID) error {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(ctx)

	book, err := models2.GetBookByID(sessCtx, bookID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := book.Delete(sessCtx); err != nil {
		return err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return err
	}

	return nil
}

func ReserveBook(ctx context.Context, bookID, userID primitive.ObjectID) (*models2.Book, error) {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

}

func addPublisher(ctx context.Context, publisherName string) (*models2.Publisher, error) {
	publisher := &models2.Publisher{
		Id:   primitive.NewObjectID(),
		Name: publisherName,
	}
	if err := publisher.Insert(ctx); err != nil {
		return nil, err
	}
	return publisher, nil
}
