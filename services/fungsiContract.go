package services

import (
	"Toko/modules"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerToko struct {
	dbcollection *mongo.Collection
	ctx          context.Context
}

func CallHandlerToko(dbcollection *mongo.Collection, ctx context.Context) Kontrak {
	return &HandlerToko{
		dbcollection: dbcollection,
		ctx:          ctx,
	}
}

func (h *HandlerToko) CreateBook(kodeToko string, buku *modules.Buku) error {
	fmt.Println(buku)
	id := bson.M{"kode_toko": kodeToko}
	data := bson.M{"$push": bson.M{"buku": bson.M{
		"kode_buku":   uuid.New().String(),
		"title":       buku.Title,
		"price":       buku.Price,
		"description": buku.Description,
	}}}

	_, err := h.dbcollection.UpdateOne(h.ctx, id, data)
	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}

func (h *HandlerToko) GetBook(kodeToko string, kodeBuku string) (*modules.Toko, error) {
	var tokos *modules.Toko
	query := bson.M{"kode_toko": kodeToko, "buku": kodeBuku}
	err := h.dbcollection.FindOne(h.ctx, query).Decode(&tokos)
	return tokos, err
}

func (h *HandlerToko) CreateToko(toko *modules.Toko) error {
	_, err := h.dbcollection.InsertOne(h.ctx, toko)
	return err
}

func (h *HandlerToko) GetToko(KodeToko string) (*modules.Toko, error) {
	var toko *modules.Toko
	query := bson.D{bson.E{Key: "kode_toko", Value: KodeToko}}
	err := h.dbcollection.FindOne(h.ctx, query).Decode(&toko)
	return toko, err
}

func (h *HandlerToko) GetAll() ([]*modules.Toko, error) {
	var tokos []*modules.Toko
	cursor, err := h.dbcollection.Find(h.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(h.ctx) {
		var toko modules.Toko
		err := cursor.Decode(&toko)
		if err != nil {
			return nil, err
		}
		tokos = append(tokos, &toko)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(h.ctx)

	if len(tokos) == 0 {
		return nil, errors.New("documents not found")
	}
	return tokos, nil
}

func (h *HandlerToko) UpdateToko(m *modules.Toko, kodeToko string) error {
	filter := bson.D{bson.E{Key: "kode_toko", Value: kodeToko}}
	fmt.Println(m.KodeToko)
	update := bson.D{
		bson.E{
			Key: "$set",
			Value: bson.D{
				bson.E{
					Key:   "nama_toko",
					Value: m.NamaToko,
				},
				bson.E{
					Key:   "domisili",
					Value: m.Domisili,
				},
			},
		},
	}
	fmt.Println("ini updateToko", m)
	result, _ := h.dbcollection.UpdateOne(h.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (h *HandlerToko) UpdateBook(kodeToko string, kodeBuku string, book *modules.Buku) error {
	query := bson.M{"$and": []interface{}{bson.M{"kode_toko": kodeToko, "buku.kode_buku": kodeBuku}}}
	update := bson.M{
		"$set": bson.M{
			"buku.$.title":       book.Title,
			"buku.$.price":       book.Price,
			"buku.$.description": book.Description,
		},
	}
	result, err := h.dbcollection.UpdateOne(h.ctx, query, update)
	fmt.Println("ini error", err)
	fmt.Println("ini result", result)
	fmt.Println("ini update", update)
	fmt.Println("ini query", query)
	if result.MatchedCount != 1 {
		errors.New("no matched document found for update")
	}
	if err != nil {
		err.Error()
	}
	return nil
}

func (h *HandlerToko) DeleteToko(KodeToko string) error {
	filter := bson.D{bson.E{Key: "kode_toko", Value: KodeToko}}
	result, _ := h.dbcollection.DeleteOne(h.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (h *HandlerToko) DeleteBook(kodeToko string, kodeBuku string) error {
	query := bson.M{"kode_toko": kodeToko}
	change := bson.M{"$pull": bson.M{"buku": bson.M{"kode_buku": kodeBuku}}}
	toko, err := h.dbcollection.UpdateOne(h.ctx, query, change)
	fmt.Println("ini query", query)
	fmt.Println("ini change", change)
	fmt.Println("ini fungsi ", toko)
	if err != nil {
		fmt.Println("ini errornya", err)
		return err
	}
	return nil
}
