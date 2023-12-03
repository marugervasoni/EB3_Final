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
	handlerturno "github.com/jum8/EBE3_Final.git/cmd/server/handler/turno"
	"github.com/jum8/EBE3_Final.git/internal/odontologo"
	"github.com/jum8/EBE3_Final.git/internal/paciente"
	"github.com/jum8/EBE3_Final.git/internal/turno"
	"github.com/jum8/EBE3_Final.git/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/jum8/EBE3_Final.git/docs"
)

const (
	puerto = "8080"
)

// @title           EBE3- FINAL: API CLINICA
// @version         1.0
// @description     This is an API for a clinic.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	repositoryTurno := turno.NewRepositoryTurno(db)
	serviceTurno := turno.NewServiceTurno(repositoryTurno, repositoryOdontologo, repositoryPaciente)
	controllerTurno := handlerturno.NewTurnoHandler(serviceTurno)	

	engine := gin.Default()

	baseGroup := engine.Group("/api/v1")

	baseGroup.GET("/ping", controllerPing.HandlerPing())

	baseGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	odontologoGroup := baseGroup.Group("/odontologos")
	{
		odontologoGroup.GET("", controllerOdontologo.HandlerGetAll())
		odontologoGroup.GET(":id", controllerOdontologo.HandlerGetById())
		odontologoGroup.POST("", middleware.Authenticate(), controllerOdontologo.HandlerCreate())
		odontologoGroup.PUT(":id", middleware.Authenticate(), controllerOdontologo.HandlerUpdate())
		odontologoGroup.DELETE(":id", middleware.Authenticate(), controllerOdontologo.HandlerDelete())
		odontologoGroup.PATCH(":id", middleware.Authenticate(), controllerOdontologo.HandlerPatch())
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

	turnoGroup := baseGroup.Group("/turnos")
    {
			turnoGroup.POST("", controllerTurno.HandlerCreate())
			turnoGroup.GET(":id", controllerTurno.HandlerGetById())
			turnoGroup.PUT(":id", controllerTurno.HandlerUpdate())
			turnoGroup.PATCH(":id", controllerTurno.HandlerPatch())
			turnoGroup.DELETE(":id", controllerTurno.HandleDelete())
			turnoGroup.GET("", controllerTurno.HandlerGetByDNI())
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