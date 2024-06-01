package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/engageapp/pkg/models"
	"github.com/engageapp/pkg/utils"
	"github.com/go-chi/chi/v5"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Will contain the code that initializes app dependancies
type Base struct {
	Router    *chi.Mux
	DB        *sql.DB
	UserModel *models.UserModel
	RabbitMQ  *amqp.Connection
	Chan      *amqp.Channel
	PostModel *models.PostModel
}

func (b *Base) Init() {

	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")
	dbPassword := os.Getenv("DBPASSWORD")
	dbUser := os.Getenv("DBUSER")
	dbPort := os.Getenv("DBPORT")

	startTime := time.Now()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	b.DB = utils.ConnectDB(dsn)
	utils.Log("INFO", "app", "connection done in %v", time.Since(startTime))

	b.RabbitMQ = utils.ConnectQueue()
	//defer b.RabbitMQ.Close()

	b.Chan = utils.CreateChannel(b.RabbitMQ)
	//defer b.Chan.Close()

}
