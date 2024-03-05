package shared

import (
	"runtime/debug"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rs/zerolog/log"
)

func LoadEnv(env string) error {
	e := env
	if e == "" {
		e = "dev"
	}
	// In production, these are injected by the container.
	if e == "dev" {
		err := godotenv.Load(".env.local")
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadBuildInfo(buildTimestamp string) (string, string, bool) {
	rev := "unknown"
	dirty := true
	timestamp := "unknown"
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, v := range buildInfo.Settings {
			switch v.Key {
			case "vcs.revision":
				if v.Value != "" && len(v.Value) > 7 {
					rev = v.Value[:7]
				}
			case "vcs.modified":
				if v.Value == "false" {
					dirty = false
				}
			}
		}
	}
	if i, err := strconv.ParseInt(buildTimestamp, 10, 64); err == nil {
		timestamp = time.Unix(i/1000, 0).Format("Mon Jan 2 15:04:05 MST 2006")
	}
	return rev, timestamp, dirty
}

type Endpoint struct {
	Subject string
	Handler micro.Handler
}

type MicroService struct {
	Name           string
	ServiceName    string
	Version        string
	Description    string
	BuildEnv       string
	BuildTimestamp string
	Endpoints      []Endpoint
}

func RunMicroservice(nc *nats.Conn, s MicroService) {
	// Load build info
	rev, ts, dirty := LoadBuildInfo(s.BuildTimestamp)

	log.Info().
		Str("rev", rev).
		Str("env", s.BuildEnv).
		Str("built", ts).
		Bool("dirty", dirty).
		Msg(s.Name)

	svc, err := micro.AddService(nc, micro.Config{
		Name:        s.ServiceName,
		Version:     s.Version,
		Description: s.Description,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("error adding service")
	}

	handlers := svc.AddGroup("svc")
	for _, e := range s.Endpoints {
		if err = handlers.AddEndpoint(e.Subject, e.Handler); err != nil {
			log.Fatal().Err(err).Msgf("error adding %s endpoint", e.Subject)
		}
	}
	log.Info().Msgf("%s %s listening on %s", svc.Info().Name, svc.Info().Version, nc.ConnectedAddr())
}
