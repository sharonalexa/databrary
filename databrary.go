// main.go
package main

import (
	"net/http"

	"os"
	"path/filepath"

	"github.com/databrary/databrary/config"
	"github.com/databrary/databrary/logging"
	"github.com/databrary/databrary/db"
	"github.com/databrary/databrary/routes"
	"github.com/databrary/databrary/services/redis"
	"github.com/databrary/databrary/services/sessions"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
	"github.com/pressly/chi/docgen"
	"gopkg.in/alecthomas/kingpin.v2"
	"time"
	"fmt"
	"strings"
)

var (
	proj_root   = strings.Split(filepath.Join(os.Getenv("GOPATH"), "src/github.com/databrary/databrary/"), ":")[1]
	config_path = kingpin.Flag("config", "Path to config file").
		Default(filepath.Join(proj_root, "config/databrary_dev.toml")).
		Short('c').
		String()
)

func init() {
	// cmd line flags
	kingpin.Version("0.0.0")
	kingpin.Parse()

	if config_path, err := filepath.Abs(*config_path); err != nil {
		panic("command line config file path error")
	} else {
		log.InitLgr(config.InitConf(config_path))
	}

	err := db.InitDB(config.GetConf())
	if err != nil {
		panic(err.Error())
	}

	redis.InitRedisStore(config.GetConf())
}

func main() {
	// New permissions middleware
	conf := config.GetConf()
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          []string{"localhost:3000", "localhost:3444", "www.petri.li:3000"},
		HostsProxyHeaders:     []string{"X-Forwarded-Host"},
		SSLRedirect:           true,
		SSLHost:               "",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
		IsDevelopment:         true,
	})
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(log.NewStructuredLogger(log.Logger))
	r.Use(middleware.Recoverer)
	r.Use(secureMiddleware.Handler) // TODO turn back on
	r.Use(middleware.Timeout(60 * time.Second))

	rateLimiter, err := routes.NewRateLimiter()
	if err != nil {
		log.WrapErrLogFatal(err, "couldn't create rate limiter")
	}

	r.Use(rateLimiter.RateLimit)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //[]string{"http://localhost:3000", "https://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		Debug:            true,
	})

	r.Use(c.Handler)
	r.Use(sessions.NewSessionManager())

	r.Use(middleware.StripSlashes)
	r.Mount("/api", routes.Api())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nothing here yet"))
	})

	r.FileServer("/public", http.Dir("public"))
	GenerateApi(r)
	addr := ":3444"
	fmt.Printf("serving on https://%s/\n", addr)

	certPath := conf.GetString("ssl.cert")
	keyPath := conf.GetString("ssl.key")
	err = http.ListenAndServeTLS(addr, certPath, keyPath, r)
	fmt.Println(err)
}

func GenerateApi(r chi.Router) {
	m := docgen.MarkdownOpts{
		ProjectPath:        proj_root,
		Intro:              "Databrary 2.0 API",
		ForceRelativeLinks: true,
	}

	f, _ := os.Create(filepath.Join(proj_root, "api.md"))
	defer f.Close()
	f.WriteString(docgen.MarkdownRoutesDoc(r, m))
	f.Sync()
}