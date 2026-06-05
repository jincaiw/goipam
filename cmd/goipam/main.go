package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/jincaiw/goipam/internal/api"
	"github.com/jincaiw/goipam/internal/auth"
	"github.com/jincaiw/goipam/internal/config"
	"github.com/jincaiw/goipam/internal/db"
	"github.com/jincaiw/goipam/internal/db/sqlc"
	"github.com/jincaiw/goipam/internal/rbac"
	"github.com/jincaiw/goipam/internal/repository"
	"github.com/jincaiw/goipam/internal/server"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	configPath := flag.String("config", "", "path to configuration file")
	showVersion := flag.Bool("version", false, "print version information and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("goipam %s (%s, %s)\n", version, commit, date)
		return
	}

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		slog.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	// Setup structured logging
	setupLogger(cfg.Log)

	slog.Info("starting GoIPAM",
		"listen", cfg.Server.Listen,
		"driver", cfg.Database.Driver,
	)

	// Connect to database
	database, err := db.Connect(&cfg.Database)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	// Run migrations
	if err := db.RunMigrations(database, &cfg.Database); err != nil {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}

	// Initialize SQLC queries
	queries := sqlc.New(database)

	// Seed initial data
	if err := db.SeedData(database, queries); err != nil {
		slog.Error("failed to seed data", "error", err)
		os.Exit(1)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(queries)
	sectionRepo := repository.NewSectionRepository(queries)
	folderRepo := repository.NewFolderRepository(queries)
	subnetRepo := repository.NewSubnetRepository(queries, database)
	addressRepo := repository.NewAddressRepository(queries)
	vlanRepo := repository.NewVlanRepository(queries)
	vrfRepo := repository.NewVrfRepository(queries)
	deviceRepo := repository.NewDeviceRepository(queries)
	locationRepo := repository.NewLocationRepository(queries)
	natRepo := repository.NewNatRepository(queries)
	nameserverRepo := repository.NewNameserverRepository(queries)
	scanRepo := repository.NewScanRepository(queries)
	requestRepo := repository.NewRequestRepository(queries)
	auditRepo := repository.NewAuditRepository(queries)
	searchRepo := repository.NewSearchRepository(queries)
	roleRepo := repository.NewRoleRepository(queries)
	groupRepo := repository.NewGroupRepository(queries)
	settingsRepo := repository.NewSettingsRepository(database)

	// Initialize services
	authService := auth.NewAuthService(userRepo, &cfg.Security)
	rbacMW := rbac.NewMiddleware(&authValidator{authService: authService})

	// Initialize handlers
	authHandler := api.NewAuthHandler(authService)
	userHandler := api.NewUserHandler(userRepo)
	sectionHandler := api.NewSectionHandler(sectionRepo, folderRepo, subnetRepo, addressRepo)
	folderHandler := api.NewFolderHandler(folderRepo)
	subnetHandler := api.NewSubnetHandler(subnetRepo, addressRepo)
	addressHandler := api.NewAddressHandler(addressRepo, subnetRepo)
	vlanHandler := api.NewVlanHandler(vlanRepo)
	vrfHandler := api.NewVrfHandler(vrfRepo)
	deviceHandler := api.NewDeviceHandler(deviceRepo)
	locationHandler := api.NewLocationHandler(locationRepo, deviceRepo)
	natHandler := api.NewNatHandler(natRepo)
	nameserverHandler := api.NewNameserverHandler(nameserverRepo)
	scanHandler := api.NewScanHandler(scanRepo, subnetRepo, addressRepo)
	requestHandler := api.NewRequestHandler(requestRepo, subnetRepo, addressRepo)
	searchHandler := api.NewSearchHandler(searchRepo)
	auditHandler := api.NewAuditHandler(auditRepo)
	importExportHandler := api.NewImportExportHandler(addressRepo, subnetRepo)
	dashboardHandler := api.NewDashboardHandler(database)
	roleHandler := api.NewRoleHandler(roleRepo)
	groupHandler := api.NewGroupHandler(groupRepo)
	settingsHandler := api.NewSettingsHandler(settingsRepo)

	// Create and start HTTP server
	srv := server.New(cfg,
		server.WithAuthHandler(authHandler),
		server.WithUserHandler(userHandler),
		server.WithSectionHandler(sectionHandler),
		server.WithFolderHandler(folderHandler),
		server.WithSubnetHandler(subnetHandler),
		server.WithAddressHandler(addressHandler),
		server.WithVlanHandler(vlanHandler),
		server.WithVrfHandler(vrfHandler),
		server.WithDeviceHandler(deviceHandler),
		server.WithLocationHandler(locationHandler),
		server.WithNatHandler(natHandler),
		server.WithNameserverHandler(nameserverHandler),
		server.WithScanHandler(scanHandler),
		server.WithRequestHandler(requestHandler),
		server.WithSearchHandler(searchHandler),
		server.WithAuditHandler(auditHandler),
		server.WithImportExportHandler(importExportHandler),
		server.WithDashboardHandler(dashboardHandler),
		server.WithRoleHandler(roleHandler),
		server.WithGroupHandler(groupHandler),
		server.WithSettingsHandler(settingsHandler),
		server.WithRBACMiddleware(rbacMW),
	)

	slog.Info("GoIPAM server listening", "addr", cfg.Server.Listen)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}

func setupLogger(cfg config.LogConfig) {
	var level slog.Level
	switch cfg.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if cfg.Format == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))
}
