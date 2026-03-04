package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//repo layer => access db calls

type Repo struct {
	coll *mongo.Collection
}

//constructor for repo
func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll : db.Collection("notes"),
	}
}
					//parentcontext from router
func(r *Repo) Create(ctx context.Context,note Note) (Note,error){
//member function of repo to create note in db
	//child context
	// if r == nil {
    //     fmt.Println("REPO IS NIL")
    // }

    // if r.coll == nil {
    //     fmt.Println("COLLECTION IS NIL")
    // }
	opCtx , cancel := context.WithTimeout(ctx,5*time.Second)
	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		fmt.Println("Mongo Insert Error:", err)
		return Note{}, fmt.Errorf("insert failed: %w", err)
	}
	return note,nil ;
}


func (r *Repo) List(ctx context.Context) ([]Note,error){
	opCtx , cancel := context.WithTimeout(ctx,5*time.Second)
	defer cancel()


	cursor , err := r.coll.Find(opCtx,bson.M{})
	if err != nil {
		fmt.Println("Mongo Insert Error:", err)
		return nil, fmt.Errorf("insert failed: %w", err)
	}
	//to avoid memory leak we need to close the cursor after using it
	defer cursor.Close(opCtx)
	var notes []Note
	if err := cursor.All(opCtx,&notes); err != nil {
		fmt.Println("Mongo Cursor Error:", err)
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return notes,nil ;
}