package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozoncp/ocp-template-api/internal/config"
	"github.com/ozoncp/ocp-template-api/internal/database"
	"github.com/ozoncp/ocp-template-api/internal/server"
	"github.com/ozoncp/ocp-template-api/internal/tracer"
	"github.com/pressly/goose/v3"
)

var (
	batchSize uint = 2
)

func main() {

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	migration := flag.String("migration", "", "Defines the migration start option")
	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	var err error
	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}

	if *migration != "" {
		if err := migrate(db.DB, *migration); err != nil {
			log.Fatal().Err(err).Msg("Failed migrating")
		}
	}

	if err := tracer.InitTracing("ocp_template_api"); err != nil {
		log.Fatal().Err(err).Msg("Failed init tracing")
	}

	if err := server.NewGrpcServer(db, batchSize).Start(&cfg); err != nil {
		log.Fatal().Err(err).Msg("Failed creating gRPC server")
	}

	db.Close()
}

func migrate(db *sql.DB, command string) error {
	switch command {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Error().Err(err).Msg("Migration failed")
			return err
		}
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Error().Err(err).Msg("Migration failed")
			return err
		}

	default:
		log.Warn().Msgf("Invalid command for 'migration' flag: '%v'", command)
	}
	return nil
}
