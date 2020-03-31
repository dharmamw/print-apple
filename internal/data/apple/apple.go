package apple

import (
	"context"
	"log"
	"math"
	"sort"
	"strings"
	"time"

	// "strconv"

	appleEntity "print-apple/internal/entity/apple"
	"print-apple/pkg/errors"
	firebaseclient "print-apple/pkg/firebaseClient"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

//Sort by date(TglTransf)
type timeSlice []appleEntity.Apple

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].TglTransf.Before(p[j].TglTransf)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

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

// New ...
func New(fb *firebaseclient.Client) Data {
	d := Data{
		fb: fb.Client,
	}
	return d
}

// GetPrintApple ...
func (d Data) GetPrintApple(ctx context.Context) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
	)
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

// GetPrintAppleStorage ...
func (d Data) GetPrintAppleStorage(ctx context.Context) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
	)
	iter := d.fb.Collection("PrintAppleStorage").Documents(ctx)
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
		appleFirebase = append(appleFirebase, apple)
	}
	return appleFirebase, err
}

//UpdateStorage ...
func (d Data) UpdateStorage(ctx context.Context, TransFH string) error {
	_, err := d.fb.Collection("PrintApple").Doc(TransFH).Update(ctx, []firestore.Update{{
		Path: "printCount", Value: 1}, {Path: "printed", Value: "Y"}})
	doc, err := d.fb.Collection("PrintApple").Doc(TransFH).Get(ctx)
	appleValidate := doc.Data()
	if appleValidate != nil {

	} else if appleValidate == nil {
		return errors.Wrap(err, "Data Not Exist")
	}
	log.Println(appleValidate)
	_, err = d.fb.Collection("PrintAppleStorage").Doc(TransFH).Set(ctx, appleValidate)
	return err
}

// DeleteAndUpdateStorage ...
func (d Data) DeleteAndUpdateStorage(ctx context.Context, TransFH string) error {
	doc, err := d.fb.Collection("PrintApple").Doc(TransFH).Get(ctx)
	appleValidate := doc.Data()
	if appleValidate == nil {
		return errors.Wrap(err, "Data Not Exist")
	}
	_, err = d.fb.Collection("PrintApple").Doc(TransFH).Delete(ctx)
	return err
}

//Insert ...
func (d Data) Insert(ctx context.Context, apple appleEntity.Apple) error {
	_, err := d.fb.Collection("PrintApple").Doc(apple.TransFH).Set(ctx, apple)

	return err
}

// GetPrintPageTemp ...
func (d Data) GetPrintPageTemp(ctx context.Context, page int, length int) ([]appleEntity.Apple, error) {
	var (
		apple    appleEntity.Apple
		apples   []appleEntity.Apple
		iter     *firestore.DocumentIterator
		lastDoc  *firestore.DocumentSnapshot
		err      error
		totalDoc int
	)

	iterPage := d.fb.Collection("PrintApple").Documents(ctx)
	for {
		_, err := iterPage.Next()
		if err == iterator.Done {
			break
		}
		totalDoc++
	}

	if page == 1 {
		// Kalau page 1 ambil data langsung dari query
		iter = d.fb.Collection("PrintApple").Limit(length).Documents(ctx)
	} else {
		// Kalau page > 1 ambil data sampai page sebelumnya
		previous := d.fb.Collection("PrintApple").Limit((page - 1) * length).Documents(ctx)
		docs, err := previous.GetAll()
		if err != nil {
			return nil, err
		}
		// Ambil doc terakhir
		lastDoc = docs[len(docs)-1]
		// Query mulai setelah doc terakhir
		iter = d.fb.Collection("PrintApple").StartAfter(lastDoc).Limit(length).Documents(ctx)
	}

	// Looping documents
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return apples, errors.Wrap(err, "[DATA][GetPrintTempStorage] Failed to iterate Document!")
		}
		err = doc.DataTo(&apple)
		if err != nil {
			return apples, errors.Wrap(err, "[DATA][GetPrintTempStorage] Failed to Populate Struct!")
		}
		test := float64(totalDoc) / float64(length)
		apple.TotalPage = int(math.Ceil(test))

		log.Println(test)
		apples = append(apples, apple)
	}
	return apples, err
}

// GetPrintPageFinal ...
func (d Data) GetPrintPageFinal(ctx context.Context, page int, length int) ([]appleEntity.Apple, error) {
	var (
		apple    appleEntity.Apple
		apples   []appleEntity.Apple
		iter     *firestore.DocumentIterator
		lastDoc  *firestore.DocumentSnapshot
		err      error
		totalDoc int
	)

	iterPage := d.fb.Collection("PrintAppleStorage").Documents(ctx)
	for {
		_, err := iterPage.Next()
		if err == iterator.Done {
			break
		}
		totalDoc++
	}

	if page == 1 {
		// Kalau page 1 ambil data langsung dari query
		iter = d.fb.Collection("PrintAppleStorage").Limit(length).Documents(ctx)
	} else {
		// Kalau page > 1 ambil data sampai page sebelumnya
		previous := d.fb.Collection("PrintAppleStorage").Limit((page - 1) * length).Documents(ctx)
		docs, err := previous.GetAll()
		if err != nil {
			return nil, err
		}
		// Ambil doc terakhir
		lastDoc = docs[len(docs)-1]
		// Query mulai setelah doc terakhir
		iter = d.fb.Collection("PrintAppleStorage").StartAfter(lastDoc).Limit(length).Documents(ctx)
	}

	// Looping documents
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return apples, errors.Wrap(err, "[DATA][GetPrintFinalStorage] Failed to iterate Document!")
		}
		err = doc.DataTo(&apple)
		if err != nil {
			return apples, errors.Wrap(err, "[DATA][GetPrintFinalStorage] Failed to Populate Struct!")
		}
		test := float64(totalDoc) / float64(length)
		apple.TotalPage = int(math.Ceil(test))

		log.Println(test)
		apples = append(apples, apple)
		log.Println(totalDoc)
		log.Println(apple.TotalPage)
	}
	return apples, err
}

// GetByTransFHTemp ...
func (d Data) GetByTransFHTemp(ctx context.Context, TransFH string) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
	)

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
		if apple.TransFH[:3] == TransFH {
			appleFirebase = append(appleFirebase, apple)
		}
	}
	return appleFirebase, err
}

// GetByTransFHFinal ...
func (d Data) GetByTransFHFinal(ctx context.Context, TransFH string) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
	)

	iter := d.fb.Collection("PrintAppleStorage").Documents(ctx)
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
		if apple.TransFH[:3] == TransFH {
			appleFirebase = append(appleFirebase, apple)
		}
	}
	return appleFirebase, err
}

// GetByTglTransfTemp ...
func (d Data) GetByTglTransfTemp(ctx context.Context, TglTransf0 string, TglTransf1 string) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
		t1            time.Time
		t2            time.Time
	)

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
		t1, _ = time.Parse("2006-01-02", TglTransf0)
		t2, _ = time.Parse("2006-01-02", TglTransf1)
		if apple.TglTransf.After(t1) && apple.TglTransf.Before(t2) {
			appleFirebase = append(appleFirebase, apple)
			sort.Sort(timeSlice(appleFirebase))

		}
	}
	return appleFirebase, err
}

// GetByTglTransfFinal ...
func (d Data) GetByTglTransfFinal(ctx context.Context, TglTransf0 string, TglTransf1 string) ([]appleEntity.Apple, error) {
	var (
		appleFirebase []appleEntity.Apple
		err           error
		t1            time.Time
		t2            time.Time
	)

	iter := d.fb.Collection("PrintAppleStorage").Documents(ctx)
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
		t1, _ = time.Parse("2006-01-02", TglTransf0)
		t2, _ = time.Parse("2006-01-02", TglTransf1)
		if apple.TglTransf.After(t1) && apple.TglTransf.Before(t2) {
			appleFirebase = append(appleFirebase, apple)
			sort.Sort(timeSlice(appleFirebase))
		}

	}
	return appleFirebase, err
}

// GetComplexPageFinal ...
func (d Data) GetComplexPageFinal(ctx context.Context, page int, length int, sortBy string) ([]appleEntity.Apple, error) {
	var (
		apple    appleEntity.Apple
		apples   []appleEntity.Apple
		iter     *firestore.DocumentIterator
		lastDoc  *firestore.DocumentSnapshot
		err      error
		totalDoc int
	)

	iterPage := d.fb.Collection("PrintAppleStorage").Documents(ctx)
	for {
		_, err := iterPage.Next()
		if err == iterator.Done {
			break
		}
		totalDoc++
	}

	if page == 1 {
		log.Println(sortBy)
		switch sortBy {
		case "newest":
			iter = d.fb.Collection("PrintAppleStorage").Limit(length).Documents(ctx)
			sort.Sort(sort.Reverse(timeSlice(apples)))
		case "ams":
			iter := d.fb.Collection("PrintAppleStorage").Documents(ctx)
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
				if apple.PaymentMethod == strings.ToUpper(sortBy) {
					apples = append(apples, apple)
				}
				iter = d.fb.Collection("PrintAppleStorage").Limit(length).Documents(ctx)
			}
		}
	} else {
		// Kalau page > 1 ambil data sampai page sebelumnya
		previous := d.fb.Collection("PrintAppleStorage").Limit((page - 1) * length).Documents(ctx)
		docs, err := previous.GetAll()
		if err != nil {
			return nil, err
		}
		// Ambil doc terakhir
		lastDoc = docs[len(docs)-1]
		// Query mulai setelah doc terakhir
		iter = d.fb.Collection("PrintAppleStorage").StartAfter(lastDoc).Limit(length).Documents(ctx)
	}

	// Looping documents
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return apples, errors.Wrap(err, "[DATA][GetPrintFinalStorage] Failed to iterate Document!")
		}
		err = doc.DataTo(&apple)
		if err != nil {
			return apples, errors.Wrap(err, "[DATA][GetPrintFinalStorage] Failed to Populate Struct!")
		}
		test := float64(totalDoc) / float64(length)
		apple.TotalPage = int(math.Ceil(test))

		log.Println(test)
		apples = append(apples, apple)
		log.Println(totalDoc)
		log.Println(apple.TotalPage)
	}
	return apples, err
}
