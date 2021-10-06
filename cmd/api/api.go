package api

import (
	"patients-api/internal/storage"
	"patients-api/patient"
	"patients-api/patient/store"
)

func Start(port string) {
	db := storage.ConnectToDB()
	defer db.Close()


	r := routes(patient.NewService(store.New(db)))
	server := newServer(port, r)

	server.Start()
}