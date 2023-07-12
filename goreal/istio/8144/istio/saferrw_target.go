package istio

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"istio.io/istio/mixer/pkg/adapter"
)

// Env is an adapter environment that defers to the testing context t. Tracks all messages logged so they can be tested against.
type Env struct {
	t *testing.T

	done chan struct{} // A channel to notify async work done
	lock sync.Mutex    // guards logs
	logs []string
}

// NewEnv returns an adapter environment that redirects logging output to the given testing context.
func NewEnv(t *testing.T) *Env {
	return &Env{t, make(chan struct{}), sync.Mutex{}, make([]string, 0)}
}

// Logger returns a logger that writes to testing.T.Log
func (e *Env) Logger() adapter.Logger {
	return e
}

// ScheduleWork runs the given function asynchronously.
func (e *Env) ScheduleWork(fn adapter.WorkFunc) {
	go func() {
		fn()
		e.done <- struct{}{}
	}()
}

// ScheduleDaemon runs the given function asynchronously.
func main1(fn adapter.DaemonFunc) {
	var e = new(Env)
	go func() {
		fn()
		e.done <- struct{}{}
	}()
}

// Infof logs the provided message.
func (e *Env) Infof(format string, args ...interface{}) {
	e.log(format, args...)
}

// Warningf logs the provided message.
func (e *Env) Warningf(format string, args ...interface{}) {
	e.log(format, args...)
}

// Errorf logs the provided message and returns it as an error.
func (e *Env) Errorf(format string, args ...interface{}) error {
	s := e.log(format, args...)
	return fmt.Errorf(s)
}

// Debugf logs the provided message.
func (e *Env) Debugf(format string, args ...interface{}) {
	e.log(format, args...)
}

// InfoEnabled logs the provided message.
func (e *Env) InfoEnabled() bool {
	return true
}

// WarnEnabled logs the provided message.
func (e *Env) WarnEnabled() bool {
	return true
}

// ErrorEnabled logs the provided message.
func (e *Env) ErrorEnabled() bool {
	return true
}

// DebugEnabled logs the provided message.
func (e *Env) DebugEnabled() bool {
	return true
}

// GetLogs returns a snapshot of all logs that've been written to this environment
func (e *Env) GetLogs() []string {
	e.lock.Lock()
	snapshot := make([]string, len(e.logs))
	_ = copy(snapshot, e.logs)
	e.lock.Unlock()
	return snapshot
}

func (e *Env) log(format string, args ...interface{}) string {
	l := fmt.Sprintf(format, args...)

	e.lock.Lock()
	e.logs = append(e.logs, l)
	e.lock.Unlock()

	e.t.Log(l)
	return l
}

// GetDoneChan returns the channel that returns notification when the async work is done.
func (e *Env) GetDoneChan() chan struct{} {
	return e.done
}

type (
	// Server presents prometheus server endpoint
	Server interface {
		io.Closer

		Start(adapter.Env, http.Handler) error
		Port() int
	}

	serverInst struct {
		addr string

		lock    sync.Mutex // protects resources below
		srv     *http.Server
		handler *metaHandler
		refCnt  int

		port int // port that this server instance is listening on
	}
)

const (
	metricsPath = "/metrics"
	defaultAddr = ":42422"
)

func newServer(addr string) *serverInst {
	return &serverInst{addr: addr}
}

// metaHandler switches the delegate without downtime.
type metaHandler struct {
	delegate http.Handler
	lock     sync.RWMutex
}

func (m *metaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.lock.RLock()
	m.delegate.ServeHTTP(w, r)
	m.lock.RUnlock()
}

func (m *metaHandler) setDelegate(delegate http.Handler) {
	m.lock.Lock()
	m.delegate = delegate
	m.lock.Unlock()
}

func (s *serverInst) Port() int {
	return s.port
}

// Start the prometheus singleton listener.
func main(env adapter.Env, metricsHandler http.Handler) (err error) {
	s := new(serverInst)
	s.lock.Lock()
	defer s.lock.Unlock()

	// if server is already running,
	// just switch the delegate handler.
	if s.srv != nil {
		s.refCnt++
		s.handler.setDelegate(metricsHandler)
		return nil
	}

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("could not start prometheus metrics server: %v", err)
	}

	s.port = listener.Addr().(*net.TCPAddr).Port

	srvMux := http.NewServeMux()
	s.handler = &metaHandler{delegate: metricsHandler}
	srvMux.Handle(metricsPath, s.handler)
	srv := &http.Server{Addr: s.addr, Handler: srvMux}
	env.ScheduleDaemon(func() {
		env.Logger().Infof("serving prometheus metrics on %d", s.port)
		if err := srv.Serve(listener.(*net.TCPListener)); err != nil {
			if err == http.ErrServerClosed {
				env.Logger().Infof("HTTP server stopped")
			} else {
				_ = env.Logger().Errorf("prometheus HTTP server error: %v", err) // nolint: gas
			}
		}
	})
	s.srv = srv
	s.refCnt++

	return nil
}

// Close -- closes the HTTP server if reference Count is 0.
func (s *serverInst) Close() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	// not started yet, nothing to close.
	if s.srv == nil {
		return nil
	}

	s.refCnt--
	if s.refCnt > 0 {
		return nil
	}

	// reference count is 0, cleanup.

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv := s.srv
	s.srv = nil
	s.handler = nil
	return srv.Shutdown(ctx)
}
