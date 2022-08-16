package managers

import (
	"context"
	"github.com/Ja7ad/library/server/db/collection"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/internal/models"
	"github.com/Ja7ad/library/server/internal/models/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(ctx context.Context) ([]*schema.Book, error) {
	books, err := models.GetAll[schema.Book](ctx, collection.BookCollection())
	if err != nil {
		return nil, err
	}
	return books, nil
}

func FindBook(ctx context.Context, bookID primitive.ObjectID) (*schema.Book, error) {
	book, err := models.GetWithId[schema.Book](ctx, collection.BookCollection(), bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func AddBook(ctx context.Context, name, publisherName string) (*schema.Book, error) {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	publisher, err := models.GetWithName[schema.Publisher](sessCtx, collection.PublisherCollection(), publisherName)
	if err != nil {
		if publisher, err = addPublisher(sessCtx, publisherName); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	book := &schema.Book{
		Id:          primitive.NewObjectID(),
		Name:        name,
		PublisherId: publisher.Id,
	}

	if err := models.Insert[schema.Book](sessCtx, collection.BookCollection(), book); err != nil {
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

func UpdateBook(ctx context.Context, bookID primitive.ObjectID, name, publisherName string) (*schema.Book, error) {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	publisher, err := models.GetWithName[schema.Publisher](ctx, collection.PublisherCollection(), publisherName)
	if err != nil {
		if publisher, err = addPublisher(sessCtx, publisherName); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	book, err := models.GetWithId[schema.Book](sessCtx, collection.BookCollection(), bookID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	book.Name = name
	book.PublisherId = publisher.Id

	if err := models.Update[schema.Book](sessCtx, collection.BookCollection(), bookID, book); err != nil {
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
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return err
	}

	book, err := models.GetWithId[schema.Book](sessCtx, collection.BookCollection(), bookID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := models.Delete[schema.Book](sessCtx, collection.BookCollection(), book.Id); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return err
	}

	return nil
}

func Reserve(ctx context.Context, userID primitive.ObjectID, bookIDs ...primitive.ObjectID) error {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(sessCtx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return err
	}

	for _, bookID := range bookIDs {
		book, err := models.GetWithId[schema.Book](sessCtx, collection.BookCollection(), bookID)
		if err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return err
			}
			return err
		}
		book.UserId = userID
		if err := models.Update[schema.Book](sessCtx, collection.BookCollection(), book.Id, book); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return err
			}
			return err
		}
	}

	if err := ReserveBook(sessCtx, userID, bookIDs...); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}

func addPublisher(ctx context.Context, publisherName string) (*schema.Publisher, error) {
	publisher := &schema.Publisher{
		Id:   primitive.NewObjectID(),
		Name: publisherName,
	}
	if err := models.Insert[schema.Publisher](ctx, collection.PublisherCollection(), publisher); err != nil {
		return nil, err
	}
	return publisher, nil
}
