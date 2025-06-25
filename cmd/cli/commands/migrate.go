package commands

import (
	"errors"
	"fmt"
	"goapp/config"
	"goapp/database"
	"goapp/pkg/console"
	"goapp/pkg/postgres"
	"regexp"

	"github.com/spf13/cobra"
)

func InitMakeMigration(root *cobra.Command, cfg *config.Config) {
	cmd := &cobra.Command{
		Use:   "make:migration",
		Short: "Создать файл миграции",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Printf("\nCreate new migration\n\n")

			nameValidator := func(value string) error {
				re := regexp.MustCompile(`^\s*[a-z_]{1}[a-z0-9_]{0,61}\s*$`)
				if value == "" || re.MatchString(value) {
					return nil
				} else {
					return errors.New("the value must satisfy a regular expression \"[a-z_]{1}[a-z0-9_]{0,61}\"")
				}
			}

			migrationType := console.Select("Migration Type", []string{"sql", "go"})
			migrationName := console.InputStringValidated("Migration Name", nameValidator)

			mg := newMigrator(cfg.Postgres)
			err := mg.Create(migrationType, migrationName)
			fmt.Println("")

			return err
		},
	}

	root.AddCommand(cmd)
}

func InitMigrate(root *cobra.Command, cfg *config.Config) {
	cmd := &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"migrate:up"},
		Short:   "Применить миграции базы данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := setPostgresDbnameFlag(cmd, cfg); err != nil {
				return err
			}

			fmt.Printf("\nMigrate database %s...\n", cfg.Postgres.Dbname)

			mg := newMigrator(cfg.Postgres)
			if err := mg.Up(); err == nil {
				fmt.Printf("\nMigration complete\n\n")
			} else {
				fmt.Printf("\nMigration error: %v\n\n", err)
			}

			return nil
		},
	}
	cmd.Flags().String("dbname", "", "База данных, применяется для создание тестовой базы")

	root.AddCommand(cmd)
}

func InitMigrateDown(rootCmd *cobra.Command, cfg *config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate:down",
		Short: "Откатить миграции базы данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := setPostgresDbnameFlag(cmd, cfg); err != nil {
				return err
			}

			fmt.Printf("\nRun migrate down on database %s...\n\n", cfg.Postgres.Dbname)

			mg := newMigrator(cfg.Postgres)
			if err := mg.Down(); err == nil {
				fmt.Printf("\nMigration complete\n\n")
			} else {
				fmt.Printf("\nMigration error: %v\n\n", err)
			}

			return nil
		},
	}
	cmd.Flags().String("dbname", "", "База данных, применяется для создание тестовой базы")

	rootCmd.AddCommand(cmd)
}

func InitMigrateStatus(root *cobra.Command, cfg *config.Config) {
	cmd := &cobra.Command{
		Use:   "migrate:status",
		Short: "Состояние миграций базы данных",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := setPostgresDbnameFlag(cmd, cfg); err != nil {
				return err
			}

			fmt.Printf("\nMigrate status on database %s\n\n", cfg.Postgres.Dbname)

			mg := newMigrator(cfg.Postgres)
			err := mg.Status()
			fmt.Println("")

			return err
		},
	}
	cmd.Flags().String("dbname", "", "База данных, применяется для создание тестовой базы")

	root.AddCommand(cmd)
}

func newMigrator(cfg postgres.Config) *database.Migrator {
	cfg.ConnPoolEnabled = false
	db, err := postgres.NewConnection(cfg)
	if err != nil {
		panic(err)
	}

	migrator, err := database.NewMigrator(db.DB)
	if err != nil {
		panic(err)
	}

	return migrator
}

func setPostgresDbnameFlag(cmd *cobra.Command, cfg *config.Config) error {
	value, _ := cmd.Flags().GetString("dbname")
	if value == "" {
		return nil
	}

	re := regexp.MustCompile(`^[a-z_]{1}[a-z0-9_]{0,61}$`)
	if !re.MatchString(value) {
		return errors.New("the dbname must satisfy a regular expression \"^[a-z_]{1}[a-z0-9_]{0,61}$\"")
	}

	cfg.Postgres.Dbname = value
	return nil
}
