package store

import (
	"database/sql"
	"log"
	"patients-api/patient/model"
)

type Store interface {
	insert(p model.PatientInsert) error
	loadAll() []model.Patient
	loadOne(id int64) (model.Patient, error)
}

type Service struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &Service{db: db}
}

func (s *Service) insert(p model.PatientInsert) error {
	_, err := s.db.Exec("insert into patient (first_name, last_name, address, phone, email) values (?,?,?,?,?)",
		p.FirstName, p.LastName, p.Address, p.Phone, p.Email)

	if err != nil {
		log.Printf("cannot save the patient, %s", err.Error())
		return err
	}

	return nil
}

func (s *Service) loadAll() []model.Patient {
	rows, err := s.db.Query("select id, first_name, last_name, address, phone, email, created_at from patient")

	if err != nil {
		log.Printf("cannot execute select query: %s", err.Error())
		return nil
	}
	defer rows.Close()
	var p []model.Patient
	for rows.Next() {
		var patient model.Patient
		err := rows.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone,
			&patient.Email, &patient.CreatedAt)
		if err != nil {
			log.Println("cannot read current row")
			return nil
		}
		p = append(p, patient)
	}

	return p
}

func (s *Service) loadOne(id int64) (model.Patient, error) {
	var patient model.Patient
	err := s.db.QueryRow(`select id, first_name, last_name, address, phone, email, created_at from patient
		where id = ?`, id).Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.Phone,
		&patient.Email, &patient.CreatedAt)

	if err != nil {
		log.Printf("cannot fetch patient")
		return patient, err
	}

	return patient, nil
}
