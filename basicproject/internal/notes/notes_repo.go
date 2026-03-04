package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//repo layer => access db calls

type Repo struct {
	coll *mongo.Collection
}

// constructor for repo
func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

// parentcontext from router
func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {
	//member function of repo to create note in db
	//child context
	// if r == nil {
	//     fmt.Println("REPO IS NIL")
	// }

	// if r.coll == nil {
	//     fmt.Println("COLLECTION IS NIL")
	// }
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		fmt.Println("Mongo Insert Error:", err)
		return Note{}, fmt.Errorf("insert failed: %w", err)
	}
	return note, nil
}

func (r *Repo) List(ctx context.Context) ([]Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.coll.Find(opCtx, bson.M{})
	if err != nil {
		fmt.Println("Mongo Insert Error:", err)
		return nil, fmt.Errorf("insert failed: %w", err)
	}
	//to avoid memory leak we need to close the cursor after using it
	defer cursor.Close(opCtx)
	var notes []Note
	if err := cursor.All(opCtx, &notes); err != nil {
		fmt.Println("Mongo Cursor Error:", err)
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return notes, nil
}

func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var note Note
	err := r.coll.FindOne(opCtx, bson.M{"_id": id}, options.FindOne()).Decode(&note)
	fmt.Printf("Repo GetByID: Found Note: %+v\n", note)
	if err != nil {
		fmt.Println("Mongo Find Error:", err)
		return Note{}, fmt.Errorf("find failed: %w", err)
	}
	return note, nil
}

func (r *Repo) UpdateByID(ctx context.Context, id primitive.ObjectID, req UpdateNoteRequest) (Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"title":     req.Title,
			"content":   req.Content,
			"pinned":    req.Pinned,
			"updatedAt": time.Now().UTC(),
		},
	}
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	var updated Note
	//find one and update => ctx , filter,update
	err := r.coll.FindOneAndUpdate(opCtx, filter, update, &opts).Decode(&updated)
	if err != nil {
		fmt.Println("Mongo Update Error:", err)
		return Note{}, fmt.Errorf("update failed: %w", err)
	}
	return updated, nil
}

func (r *Repo) DeleteByID(ctx context.Context, id primitive.ObjectID) (bool,error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	 
	res, err := r.coll.DeleteOne(opCtx, bson.M{"_id": id}) 
	if err != nil {
		fmt.Println("Mongo Delete Error:", err)
		return 	false, fmt.Errorf("delete failed: %w", err)
	}
	if res.DeletedCount == 0 {
		return false, nil
	}
	return true,nil
}
