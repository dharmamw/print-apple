package apple

import (
	"context"
	"log"

	appleEntity "print-apple/internal/entity/apple"
	// "print-apple/pkg/errors"
	firebaseclient "print-apple/pkg/firebaseClient"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	// "github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		fb *firestore.Client
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	getAllUsers  = "GetAllUsers"
	qGetAllUsers = "SELECT * FROM user_test"
)

var (
	readStmt = []statement{
		{getAllUsers, qGetAllUsers},
	}
)

// New ...
// db *sqlx.DB
func New(fb *firebaseclient.Client) Data {
	d := Data{
		// db: db,
		fb: fb.Client,
	}

	// d.initStmt()
	return d
}

// GetAppleFromFireBase ...
func (d Data) GetAppleFromFireBase(ctx context.Context) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
	)
	// test := d.fb.Collection("user_test")
	iter := d.fb.Collection("PrintApple").Documents(ctx)
	for {
		var apple appleEntity.Apple
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		log.Println(doc)
		err = doc.DataTo(&apple)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(apple)
		appleFirebase = append(appleFirebase, apple)
	}
	return appleFirebase, err
}
