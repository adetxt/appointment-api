package main

import (
	"appointment-api/config"
	"appointment-api/graph"
	"appointment-api/usecase"
	"appointment-api/utils/db"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-sql-driver/mysql"

	appointmentrepo "appointment-api/repository/appointment_repo"
	userrepo "appointment-api/repository/user_repo"
	workinghourrepo "appointment-api/repository/working_hour_repo"
)

func main() {
	cfg := config.New()

	loc, _ := time.LoadLocation("UTC")
	mainDb := db.Open(mysql.Config{
		Addr:      fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		Net:       "tcp",
		DBName:    cfg.DBName,
		User:      cfg.DBUsername,
		Passwd:    cfg.DBPassword,
		ParseTime: true,
		Loc:       loc,
	})

	mainDb.SetDebug(true)

	userRepo := userrepo.New(mainDb.GetDB())
	workingHourRepo := workinghourrepo.New(mainDb.GetDB())
	appointmentRepo := appointmentrepo.New(mainDb.GetDB())

	userUc := usecase.NewUserUsecase(userRepo)
	workingHourUc := usecase.NewWorkingHourUsecase(workingHourRepo)
	appointmentUc := usecase.NewAppointmentUsecase(appointmentRepo, workingHourRepo, userRepo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserUc:        userUc,
		WorkingHourUc: workingHourUc,
		AppointmentUc: appointmentUc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, nil))
}
