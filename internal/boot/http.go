package boot

import (
	"log"
	"net/http"

	"print-apple/internal/config"

	appleData "print-apple/internal/data/apple"
	server "print-apple/internal/delivery/http"
	appleHandler "print-apple/internal/delivery/http/apple"
	appleService "print-apple/internal/service/apple"

	firebaseclient "print-apple/pkg/firebaseClient"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s   server.Server         // HTTP Server Object
		ad  appleData.Data        // User domain data layer
		as  appleService.Service  // User domain service layer
		ah  *appleHandler.Handler // User domain handler
		cfg *config.Config        // Configuration object
		fb  *firebaseclient.Client
	)

	// Get configuration
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()

	fb, err = firebaseclient.NewClient(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Apple domain initialization
	ad = appleData.New(fb)
	as = appleService.New(ad)
	ah = appleHandler.New(as)

	// Inject service used on handler
	s = server.Server{
		Apple: ah,
	}

	// Error Handling
	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
