package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	handlerodontologos "github.com/jum8/EBE3_Final.git/cmd/server/handler/odontologo"
	handlerpaciente "github.com/jum8/EBE3_Final.git/cmd/server/handler/paciente"
	"github.com/jum8/EBE3_Final.git/cmd/server/handler/ping"
	"github.com/jum8/EBE3_Final.git/internal/odontologo"
	"github.com/jum8/EBE3_Final.git/internal/paciente"
)

const (
	puerto = "8080"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := connectDB()

	controllerPing := ping.NewControllerPing()

	repositoryOdontologo := odontologo.NewRepositoryOdontologo(db)
	serviceOdontologo := odontologo.NewServiceOdontologo(repositoryOdontologo)
	controllerOdontologo := handlerodontologos.NewControllerOdontologo(serviceOdontologo)

	repositoryPaciente := paciente.NewRepositoryPaciente(db)
	servicePaciente := paciente.NewServicePaciente(repositoryPaciente)
	controllerPaciente := handlerpaciente.NewPacienteHandler(servicePaciente)	

	engine := gin.Default()

	baseGroup := engine.Group("/api/v1")

	baseGroup.GET("/ping", controllerPing.HandlerPing())

	odontologoGroup := baseGroup.Group("/odontologos")
	{
		odontologoGroup.GET("", controllerOdontologo.HandlerGetAll())
		odontologoGroup.GET(":id", controllerOdontologo.HandlerGetById())
		odontologoGroup.POST("", controllerOdontologo.HandlerCreate())
		odontologoGroup.PUT(":id", controllerOdontologo.HandlerUpdate())
		odontologoGroup.DELETE(":id", controllerOdontologo.HandlerDelete())
		odontologoGroup.PATCH(":id", controllerOdontologo.HandlerPatch())
	}

	pacienteGroup := baseGroup.Group("/pacientes")
    {
        pacienteGroup.GET("", controllerPaciente.HandlerGetAll())
        pacienteGroup.GET(":id", controllerPaciente.HandlerGetById())
        pacienteGroup.POST("", controllerPaciente.HandlerCreate())
        pacienteGroup.PUT(":id", controllerPaciente.HandlerUpdate())
        pacienteGroup.DELETE(":id", controllerPaciente.HandlerDelete())
        pacienteGroup.PATCH(":id", controllerPaciente.HandlerPatch())
    }

	if err := engine.Run(fmt.Sprintf(":%s", puerto)); err != nil {
		panic(err)
	}

	defer db.Close()
}

func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")

	// Create the data source.
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open the connection.
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
