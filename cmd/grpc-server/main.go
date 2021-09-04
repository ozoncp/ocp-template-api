package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozoncp/ocp-template-api/internal/config"

	pb "github.com/ozoncp/ocp-template-api/pkg/ocp-template-api"

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

	if err := RunServer(cfg.Grpc.Host, cfg.Grpc.Port); err != nil {
		log.Fatal().Err(err).Msg("Failed creating gRPC server")
	}

	db.Close()
}

func RunServer(host string, port int) error {
	listenOn := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	pb.RegisterOcpTemplateApiServiceServer(server, &ocpTemplateApiServiceServer{})
	log.Info().Msgf("Listening on %d", listenOn)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

type ocpTemplateApiServiceServer struct {
	pb.UnimplementedOcpTemplateApiServiceServer
}

func (s *ocpTemplateApiServiceServer) CreateTemplateV1(
	ctx context.Context,
	req *pb.CreateTemplateV1Request,
) (*pb.CreateTemplateV1Response, error) {

	return &pb.CreateTemplateV1Response{}, nil
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
