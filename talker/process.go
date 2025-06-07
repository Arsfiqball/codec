package talker

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

// Process is a process that can be run.
// This struct is used by the Serve function (check out the example in the Serve function).
type Process struct {
	Start       Callback     // Start is a callback that runs when the process starts.
	Live        Callback     // Live is a callback that runs periodically to check if the process is still alive.
	Ready       Callback     // Ready is a callback that runs periodically to check if the process is ready to serve requests.
	Stop        Callback     // Stop is a callback that runs when the process stops.
	Logger      *slog.Logger // Logger is the logger used by the process.
	MonitorAddr string       // MonitorAddr is the address used by the process to serve health check requests.
}

func emptyCallback(ctx context.Context) error {
	return nil
}

func sanitizeProcess(proc Process) Process {
	if proc.Start == nil {
		proc.Start = emptyCallback
	}

	if proc.Live == nil {
		proc.Live = emptyCallback
	}

	if proc.Ready == nil {
		proc.Ready = emptyCallback
	}

	if proc.Stop == nil {
		proc.Stop = emptyCallback
	}

	if proc.Logger == nil {
		proc.Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	}

	if proc.MonitorAddr == "" {
		proc.MonitorAddr = ":0" // Random port
	}

	return proc
}

func callbackToHealthCheckHandler(cb Callback) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := cb(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

// Serve runs the process.
// Example:
//
//	proc := talker.Process{
//		Start: func(ctx context.Context) error {
//			// ... do something, like starting a server
//			return nil
//		},
//		Live: func(ctx context.Context) error {
//			// ... do something, like checking if the server is still alive
//			return nil
//		},
//		Ready: func(ctx context.Context) error {
//			// ... do something, like checking if the server is ready to serve requests
//			return nil
//		},
//		Stop: func(ctx context.Context) error {
//			// ... do something, like stopping the server
//			return nil
//		},
//		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
//		MonitorAddr: ":8086", // Monitor address, default is ":0" (random port)
//	}
//
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM) // Listen to SIGINT and SIGTERM (Ctrl+C and docker stop)
//	talker.Serve(proc, sig) // Start the process, and stop it when the signal is received
func Serve(proc Process, stopSignal chan os.Signal) {
	proc = sanitizeProcess(proc)

	proc.Logger.Info("Start process")

	mainCtx, mainCancel := context.WithCancel(context.Background())

	// Health check server
	go func() {
		mux := http.NewServeMux()

		mux.HandleFunc("/live", callbackToHealthCheckHandler(proc.Live))
		mux.HandleFunc("/ready", callbackToHealthCheckHandler(proc.Ready))

		server := http.Server{
			Addr:    proc.MonitorAddr,
			Handler: mux,
		}

		listener, err := net.Listen("tcp", server.Addr)
		if err != nil {
			proc.Logger.Error(err.Error())
			return
		}

		defer listener.Close() // Ensure listener is closed after Serve() returns

		proc.Logger.Info("Monitor address: " + listener.Addr().String())

		go func() {
			<-stopSignal

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			err := server.Shutdown(ctx)
			if err != nil {
				proc.Logger.Error(err.Error())
			}
		}()

		err = server.Serve(listener)
		if err != nil {
			proc.Logger.Error(err.Error())
		}
	}()

	// Stop process when stopSignal is received
	go func() {
		<-stopSignal

		proc.Logger.Info("Stop process")

		stopCtx, stopCancel := context.WithCancel(context.Background())

		err := proc.Stop(stopCtx)
		if err != nil {
			proc.Logger.Error(err.Error())
		}

		stopCancel()
		mainCancel()
	}()

	// Start process
	err := proc.Start(mainCtx)
	if err != nil {
		proc.Logger.Error(err.Error())
	}

	// Block until mainCtx is canceled
	<-mainCtx.Done()
}
