package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozoncp/ocp-template-api/internal/config"

	"github.com/pressly/goose/v3"
)

func NewPostgres(dsn, driver string) *sqlx.DB {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to create database connection")

		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatal().Err(err).Msgf("failed ping the database")

		return nil
	}

	return db
}

func main() {
	migration := flag.String("migration", "", "Defines the migration start option")
	flag.Parse()

	const configYML = "config.yml"

	if err := config.ReadConfigYML(configYML); err != nil {
		log.Fatal().
			Err(err).
			Msg("Reading configuration")
	}

	cfg := config.GetConfigInstance()

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

	db := NewPostgres(dsn, cfg.Database.Driver)

	if *migration != "" {
		Migrate(db.DB, *migration)
	}

	InitTracing("ocp_template_api")

	// if err := server.NewGrpcServer(db).Start(); err != nil {
	// 	log.Fatal().Err(err).Msg("Failed creating gRPC server")
	// }

	db.Close()
}

func Migrate(db *sql.DB, command string) {
	switch command {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatal().Err(err).Msg("Migration failed")
		}
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatal().Err(err).Msg("Migration failed")
		}

	default:
		log.Warn().Msgf("Invalid command for 'migration' flag: '%v'", command)
	}
}

func InitTracing(serviceName string) {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		log.Fatal().Err(err).Msgf("Jaeger Tracer initialization error")
	}

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
}
