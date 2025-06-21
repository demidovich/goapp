package database

import (
	"database/sql"
	"embed"
	"regexp"
	"strings"
	"text/template"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Migrator struct {
	db *sql.DB
}

var createTableSQLTemplate = `
-- +goose Up
-- +goose StatementBegin
create sequence if not exists %s_id_seq start 1;
alter sequence %s_id_seq restart with 1;

create table %s (
	id bigint default nextval('%s_id_seq'::regclass) not null,
	created_at timestamp without time zone,
	updated_at timestamp without time zone,
	constraint %s_pkey primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists %s;
-- +goose StatementEnd
`

func NewMigrator(db *sql.DB) (*Migrator, error) {
	m := &Migrator{
		db: db,
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return m, err
	}

	goose.SetBaseFS(embedMigrations)
	goose.SetSequential(true)

	return m, nil
}

func (m *Migrator) Up() error {
	return goose.Up(m.db, "migrations")
}

func (m *Migrator) Down() error {
	return goose.Down(m.db, "migrations")
}

func (m *Migrator) Status() error {
	return goose.Status(m.db, "migrations")
}

func (m *Migrator) Create(migrationType, migrationName string) error {
	re := regexp.MustCompile(`(?i)^create[-_ ]{1,}([a-z_\d]+)[-_ ]{1,}table$`)
	match := re.FindStringSubmatch(migrationName)

	if len(match) < 2 || migrationType != "sql" {
		return goose.Create(m.db, "database/migrations", migrationName, migrationType)
	} else {
		tableName := match[1]
		return m.createTable(migrationType, migrationName, tableName)
	}
}

func (m *Migrator) createTable(migrationType, migrationName, tableName string) error {
	templateString := strings.ReplaceAll(createTableSQLTemplate, "%s", tableName)

	template, err := template.New("createTableSQLTemplate").Parse(templateString)
	if err != nil {
		return err
	}

	return goose.CreateWithTemplate(m.db, "database/migrations", template, migrationName, migrationType)
}
