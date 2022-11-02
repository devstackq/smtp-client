package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/devstackq/smtp-mailer/config"
	"github.com/devstackq/smtp-mailer/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TemplateRepository interface {
	GetListTemplates() ([]models.Template, error)
	GetTemplateById(string) (*models.Template, error)
	Create(models.Template) error
}

type TemplateMongo struct {
	template *mongo.Collection
}

func NewTemplateMongo(client *mongo.Client) *TemplateMongo {
	tmplCol := client.Database(config.MONGO_DB_NAME).Collection(config.MONGO_TEMPLATE_COL)
	return &TemplateMongo{
		template: tmplCol,
	}
}

func (t TemplateMongo) Create(tmpl models.Template) error {
	tmpl.ID = primitive.NewObjectID()

	id, err := t.template.InsertOne(context.TODO(), tmpl)
	if err != nil {
		return err
	}
	fmt.Println(id.InsertedID.(primitive.ObjectID), "inserted id ")
	return nil
}

func (t TemplateMongo) GetTemplateById(id string) (*models.Template, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{
		{Key: "_id", Value: objID},
	}
	var tmpl models.Template

	if err := t.template.FindOne(
		context.TODO(),
		filter,
	).Decode(&tmpl); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("no found template by id")
		}
		return nil, err
	}

	return &tmpl, nil
}

func (t TemplateMongo) GetListTemplates() ([]models.Template, error) {
	resp, err := t.template.Find(context.TODO(), nil)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	var tmpls []models.Template
	if err := resp.All(context.TODO(), &tmpls); err != nil {
		return nil, err
	}
	return tmpls, nil
}
