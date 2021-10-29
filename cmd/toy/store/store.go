package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/brunocalza/textile-stack/cmd/toy/store/internal/db"
	"github.com/brunocalza/textile-stack/cmd/toy/store/migrations"
	pb "github.com/brunocalza/textile-stack/gen/proto/person"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Store is a store for authentication information.
type Store struct {
	conn *sql.DB
	db   *db.Queries
}

// New returns a new Store.
func New(postgresURI string) (*Store, error) {
	as := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	conn, err := MigrateAndConnectToDB(postgresURI, as)
	if err != nil {
		return nil, fmt.Errorf("initializing db connection: %s", err)
	}

	s := &Store{
		conn: conn,
		db:   db.New(conn),
	}

	return s, nil
}

// ListPeople retrieves all persons
func (s *Store) ListPeople(ctx context.Context) ([]db.Person, bool, error) {
	people, err := s.db.ListPeople(ctx)
	if err == sql.ErrNoRows {
		return []db.Person{}, false, nil
	}
	if err != nil {
		return []db.Person{}, false, fmt.Errorf("db list people: %s", err)
	}

	return people, true, nil
}

// CreatePerson creates a new person
func (s *Store) CreatePerson(ctx context.Context, person pb.Person, data []byte) error {
	var email sql.NullString
	if person.Email == nil {
		email.Valid = false
	} else {
		email.String = *person.Email
		email.Valid = true
	}

	params := db.CreatePersonParams{
		ID:     person.Id,
		Name:   person.Name,
		Email:  email,
		PbData: data,
	}
	err := s.db.CreatePerson(ctx, params)
	if err != nil {
		return fmt.Errorf("db create person: %s", err)
	}

	return nil
}

// MigrateAndConnectToDB run db migrations and return a ready to use connection to the Postgres database.
func MigrateAndConnectToDB(postgresURI string, as *bindata.AssetSource) (*sql.DB, error) {
	// To avoid dealing with time zone issues, we just enforce UTC timezone
	if !strings.Contains(postgresURI, "timezone=UTC") {
		return nil, errors.New("timezone=UTC is required in postgres URI")
	}
	d, err := bindata.WithInstance(as)
	if err != nil {
		return nil, fmt.Errorf("creating source driver: %s", err)
	}
	m, err := migrate.NewWithSourceInstance("go-bindata", d, postgresURI)
	if err != nil {
		return nil, fmt.Errorf("creating migration: %s", err)
	}
	version, dirty, err := m.Version()
	log.Printf("current version %d, dirty %v, err: %v", version, dirty, err)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("running migration up: %s", err)
	}
	conn, err := sql.Open("pgx", postgresURI)
	if err != nil {
		return nil, fmt.Errorf("creating pgx connection: %s", err)
	}

	return conn, nil
}
