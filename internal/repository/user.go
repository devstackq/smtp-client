package repository

import (
	"github.com/devstackq/smtp-mailer/config"
	"github.com/devstackq/smtp-mailer/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(models.User) error
	GetListUser() ([]models.User, error)
	GetUserById(string) (*models.User, error)
}
type UserMongo struct {
	user *mongo.Collection
}

func NewUserMongo(client *mongo.Client) *UserMongo {
	userCol := client.Database(config.MONGO_DB_NAME).Collection(config.MONGO_TEMPLATE_COL)
	return &UserMongo{
		user: userCol,
	}
}

func (t UserMongo) Create(user models.User) error {
	return nil
}
func (t UserMongo) GetListUser() ([]models.User, error) {
	return nil, nil
}

func (t UserMongo) GetUserById(id string) (*models.User, error) {
	return nil, nil
}
