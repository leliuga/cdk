package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/leliuga/cdk/service/middleware/requestid"
	"github.com/leliuga/cdk/types"
	"k8s.io/klog/v2"
)

const (
	DefaultDomain          = "leliuga.com"
	DefaultApplicationName = "Application"
	DefaultVendor          = `Leliuga`
)

// NewService creates a new service.
func NewService(options *Options) *Service {
	s := &Service{
		Options: options,
		App: fiber.New(fiber.Config{
			ServerHeader:                 strings.ToLower(options.Name),
			StrictRouting:                false,
			CaseSensitive:                false,
			Immutable:                    false,
			UnescapePath:                 false,
			ETag:                         false,
			BodyLimit:                    options.BodyLimit,
			Concurrency:                  options.Concurrency,
			Views:                        options.Views,
			PassLocalsToViews:            true,
			ReadTimeout:                  options.ReadTimeout,
			WriteTimeout:                 options.WriteTimeout,
			IdleTimeout:                  options.IdleTimeout,
			ReadBufferSize:               options.ReadBufferSize,
			WriteBufferSize:              options.WriteBufferSize,
			CompressedFileSuffix:         DefaultCompressedFileSuffix,
			GETOnly:                      false,
			ErrorHandler:                 options.ErrorHandler,
			DisableKeepalive:             false,
			DisableDefaultDate:           false,
			DisableDefaultContentType:    false,
			DisableHeaderNormalizing:     false,
			DisableStartupMessage:        true,
			AppName:                      options.Name,
			StreamRequestBody:            false,
			DisablePreParseMultipartForm: false,
			ReduceMemoryUsage:            false,
			JSONEncoder:                  json.Marshal,
			JSONDecoder:                  json.Unmarshal,
			Network:                      options.Network,
			EnableTrustedProxyCheck:      options.EnableTrustedProxyCheck,
			TrustedProxies:               options.TrustedProxies,
			EnableIPValidation:           false,
			EnablePrintRoutes:            options.EnablePrintRoutes,
			ColorScheme:                  fiber.DefaultColors,
			RequestMethods:               fiber.DefaultMethods,
		}),
	}

	s.Use(
		recover.New(),
		compress.New(compress.Config{
			Level: compress.LevelBestSpeed,
		}),
		requestid.New(),
		etag.New(),
	)

	s.Get(DefaultPathMonitoring, func(c *fiber.Ctx) error {
		return c.JSON(types.Map[string]{
			"status": "ok",
		})
	})

	if options.Handlers != nil {
		s.Handlers.Init(s)
	}

	return s
}

// Serve the service
func (s *Service) Serve() error {
	if err := s.Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	<-ch

	klog.InfoS("the service is shutting down...", "name", s.Options.Name, "port", s.Port)
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	if err := s.Shutdown(); err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

// Start the service
func (s *Service) Start() error {
	klog.InfoS("the service is serving", "name", s.Options.Name, "port", s.Port)
	go func() {
		address := fmt.Sprintf(":%d", s.Port)

		if s.CertificateFile != "" && s.CertificateKeyFile != "" {
			if err := s.ListenTLS(address, s.CertificateFile, s.CertificateKeyFile); err != nil {
				klog.ErrorS(err, "failed to start the service", "name", s.Options.Name, "port", s.Port)
				os.Exit(1)
			}
		} else {
			if err := s.Listen(address); err != nil {
				klog.ErrorS(err, "failed to start the service", "name", s.Options.Name, "port", s.Port)
				os.Exit(1)
			}
		}
	}()

	return nil
}
