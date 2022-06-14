package book

import (
	"context"
	"github.com/Ja7ad/library/proto/protoModel/user"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/internal/book/models"
	userGrpc "github.com/Ja7ad/library/server/internal/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(ctx context.Context) ([]*models.Book, error) {
	books, err := models.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func FindBook(ctx context.Context, name, publisherName string, bookID, publisherID primitive.ObjectID) (*models.Book, error) {
	book, err := models.FindBook(ctx, name, publisherName, bookID, publisherID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func AddBook(ctx context.Context, name, publisherName string) (*models.Book, error) {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.BookClient.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	publisher, err := models.GetPublisherByName(sessCtx, publisherName)
	if err != nil {
		if publisher, err = addPublisher(sessCtx, publisherName); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	book := &models.Book{
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

func UpdateBook(ctx context.Context, bookID primitive.ObjectID, name, publisherName string) (*models.Book, error) {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.BookClient.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	publisher, err := models.GetPublisherByName(sessCtx, publisherName)
	if err != nil {
		if publisher, err = addPublisher(sessCtx, publisherName); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	book, err := models.GetBookByID(sessCtx, bookID)
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

	if err := global.BookClient.StartTransaction(sessCtx); err != nil {
		return err
	}

	book, err := models.GetBookByID(sessCtx, bookID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := book.Delete(sessCtx); err != nil {
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

func ReserveBook(ctx context.Context, userID primitive.ObjectID, bookIDs ...primitive.ObjectID) error {
	sessCtx, err := global.BookClient.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(sessCtx)

	if err := global.BookClient.StartTransaction(sessCtx); err != nil {
		return err
	}

	userGRPC := new(userGrpc.UserServer)
	reserveReq := &user.ReserveUserBookRequest{UserId: userID.Hex()}

	for _, bookID := range bookIDs {
		book, err := models.GetBookByID(sessCtx, bookID)
		if err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return err
			}
			return err
		}
		book.UserId = userID
		if err := book.Update(sessCtx); err != nil {
			if err := sessCtx.AbortTransaction(ctx); err != nil {
				return err
			}
			return err
		}
		reserveReq.BookId = append(reserveReq.BookId, bookID.Hex())
	}

	if _, err := userGRPC.ReserveBook(sessCtx, reserveReq); err != nil {
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

func addPublisher(ctx context.Context, publisherName string) (*models.Publisher, error) {
	publisher := &models.Publisher{
		Id:   primitive.NewObjectID(),
		Name: publisherName,
	}
	if err := publisher.Insert(ctx); err != nil {
		return nil, err
	}
	return publisher, nil
}
