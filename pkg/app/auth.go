package app

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"gopkg.in/validator.v2"
	"sabariram.com/goserverbase/baseapp"
	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/goserverbase/log"
	"sabariram.com/goserverbase/log/logwriter"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/config"
	"sabariram.com/rolebasedauth/pkg/model"
)

type RoleBasedAuthentication struct {
	b         *baseapp.BaseApp
	db        *mongo.Mongo
	log       *log.Logger
	validator *validator.Validator
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
	r.log.Info(ctx, "App Created", nil)
	r.registerBookStoreRoutes(r.b.GetRouter())
	r.log.Info(ctx, "Routes Registered", nil)
	r.log.Info(ctx, "Starting server on port - "+r.b.GetPort(), nil)
	return r, nil
}

func (rba *RoleBasedAuthentication) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")
	var err error
	if apiKey != "" {
		cur := rba.adminAuth.FindOne(r.Context(), map[string]interface{}{"apiKey": apiKey, "isActive": true})
		val := &model.Admin{}
		err = cur.Decode(val)
		if err == nil {
			r = r.WithContext(context.WithValue(r.Context(), constant.ActorIdKey, val.Name))
			r = r.WithContext(context.WithValue(r.Context(), constant.ClaimsKey, val.Claims))
			rba.b.ServeHTTP(w, r)
			return
		}
		err = errors.NewCustomError("INVALID_API_KEY", "invalid api key", nil)
	} else {
		err = errors.NewCustomError("MISSING_API_KEY", "requires api key", nil)
	}
	w.Header().Set(constant.HeaderContentType, constant.ContentTypeJSON)
	w.WriteHeader(http.StatusUnauthorized)
	b := err.Error()
	_, err = w.Write([]byte(b))
	if err != nil {
		panic(err)
	}

}

func (rba *RoleBasedAuthentication) GetLogger() *log.Logger {
	return rba.log
}

func (rba *RoleBasedAuthentication) GetPort() string {
	return rba.b.GetPort()
}
