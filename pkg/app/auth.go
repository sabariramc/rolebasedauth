package app

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"gopkg.in/validator.v2"
	"sabariram.com/goserverbase/baseapp"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/log"
	"sabariram.com/goserverbase/log/logwriter"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/config"
)

type RoleBasedAuthentication struct {
	b         *baseapp.BaseApp
	db        *mongo.Mongo
	log       *log.Logger
	validator *validator.Validator
	adminAuth *mongo.Collection
}

func GetDefaultApp() (*RoleBasedAuthentication, error) {
	c := config.NewConfig()
	hostParams := &log.HostParams{
		Host:        c.App.Host,
		Version:     "1.0",
		ServiceName: c.App.ServiceName,
	}
	errLog := logwriter.NewConsoleWriter(*hostParams)
	graylogger, err := logwriter.NewGraylogUDP(*hostParams, errLog, logwriter.Endpoint{
		Address: c.Logger.GrayLog.Address,
		Port:    c.Logger.GrayLog.Port,
	})
	if err != nil {
		return nil, err
	}
	runtime.GOMAXPROCS(c.Runtime.GoMaxProcs)
	consoleLogger := logwriter.NewConsoleWriter(*hostParams)
	lmux := log.NewChanneledLogMultipluxer(uint8(c.Logger.BufferSize), consoleLogger, graylogger)
	return GetApp(c, lmux, consoleLogger, utils.IST)
}

func GetApp(c *config.Config, lMux log.LogMultipluxer, auditLog log.AuditLogWriter, timeZone *time.Location) (*RoleBasedAuthentication, error) {
	r := &RoleBasedAuthentication{
		b: baseapp.NewBaseApp(baseapp.ServerConfig{
			LoggerConfig: c.Logger,
			AppConfig:    c.App,
		}, lMux, auditLog, timeZone),
		validator: validator.NewValidator(),
	}
	ctx := r.b.GetCorrelationContext(context.Background(), log.GetDefaultCorrelationParams(c.App.ServiceName))
	r.log = r.b.GetLogger()
	conn, err := mongo.NewMongo(ctx, r.log, *c.Mongo)
	if err != nil {
		return nil, err
	}
	r.db = conn
	r.registerValidator()
	r.adminAuth = conn.NewCollection("")
	r.log.Info(ctx, "App Created", nil)
	r.registerBookStoreRoutes(r.b.GetRouter())
	r.log.Info(ctx, "Routes Registered", nil)
	r.log.Info(ctx, "Starting server on port - "+r.b.GetPort(), nil)
	return r, nil
}

func (rba *RoleBasedAuthentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rba.b.ServeHTTP(w, r)
}

func (rba *RoleBasedAuthentication) GetLogger() *log.Logger {
	return rba.log
}

func (rba *RoleBasedAuthentication) GetPort() string {
	return rba.b.GetPort()
}
