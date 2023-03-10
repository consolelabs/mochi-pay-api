// Package appmain contains the common application initialization code for Social Payment servers.
package appmain

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/consolelabs/mochi-pay-api/internal/config"
	"github.com/consolelabs/mochi-pay-api/internal/logging"
)

var (
	logger = logrus.WithFields(logrus.Fields{
		"app":       "mochi-pay-api",
		"component": "app.main",
	})
)

// RunApplication starts and runs the given application forever.  For use in
// main functions to run the full application.
func RunApplication(serviceName string, bindService Bind) {
	c := make(chan os.Signal, 1)
	// SIGTERM is signaled by k8s when it wants a pod to stop.
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	readConfig := func() (config.View, error) {
		return config.Read()
	}

	a, err := newApplication(serviceName, bindService, readConfig, net.Listen)
	if err != nil {
		logger.Fatal(err)
	}

	go a.serve()
	logger.Info("Application started successfully")

	<-c
	err = a.stop()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Application stopped successfully.")
}

// Bind is a function which starts an application, and binds it to serving.
type Bind func(p *Params, b *Bindings) IServer

// Params are inputs to starting an application.
type Params struct {
	config      config.View
	logger      *logrus.Entry
	serviceName string
}

// Config provides the configuration for the application.
func (p *Params) Config() config.View {
	return p.config
}

// Logger provides a logger for the application.
func (p *Params) Logger() *logrus.Entry {
	return p.logger
}

type IServer interface {
	ListenAndServe() error
}

// ServiceName is a name for the currently running binary specified by
// RunApplication.
func (p *Params) ServiceName() string {
	return p.serviceName
}

// Bindings allows applications to bind various functions to the running servers.
type Bindings struct {
	a *App
}

// AddCloserErr specifies a function to be called when the application is being
// stopped.  Closers are called in reverse order.  The first error returned by
// a closer will be logged.
func (b *Bindings) AddCloserErr(c func() error) {
	b.a.closers = append(b.a.closers, c)
}

// App is used internally, and public only for apptest.  Do not use, and use apptest instead.
type App struct {
	closers []func() error
	srv     IServer
}

// newApplication is used internally, and public only for apptest.  Do not use, and use apptest instead.
func newApplication(serviceName string, bindService Bind, getCfg func() (config.View, error), listen func(network, address string) (net.Listener, error)) (*App, error) {
	a := &App{}

	cfg, err := getCfg()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatalf("cannot read configuration.")
	}
	logging.ConfigureLogging(cfg)

	p := &Params{
		config:      cfg,
		logger:      logger,
		serviceName: serviceName,
	}
	b := &Bindings{
		a: a,
	}

	a.srv = bindService(p, b)

	return a, nil
}

// stop is used internally, and public only for apptest.  Do not use, and use apptest instead.
func (a *App) stop() error {
	// Use closers in reverse order: Since dependencies are created before
	// their dependants, this helps ensure no dependencies are closed
	// unexpectedly.
	var firstErr error
	for i := len(a.closers) - 1; i >= 0; i-- {
		err := a.closers[i]()
		if firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (a *App) serve() error {
	return a.srv.ListenAndServe()
}
