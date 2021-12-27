package main

import (
        "fmt"
        "html/template"
        "log"
        "net/http"
        "os"
        "time"
        "embed"
)

var resources embed.FS

func logging(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                start := time.Now()
                req := fmt.Sprintf("%s %s", r.Method, r.URL)
                log.Println(req)
                next.ServeHTTP(w, r)
                log.Println(req, "completed in", time.Now().Sub(start))
        })
}

// templates references the specified templates and caches the parsed results
// to help speed up response times.
var templates = template.Must(template.ParseFiles("/http/templates/base.html"))

// index is the handler responsible for rending the index page for the site.
func index() http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                b := struct {
                        Title        template.HTML
                        Name string
                }{
                        Title:        template.HTML("Business &verbar; Landing"),
                        Name: "Trufflehog,",
                }
                err := templates.ExecuteTemplate(w, "base", &b)
                if err != nil {
                        http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
                        return
                }
                w.WriteHeader(http.StatusOK)
        })
}

func main() {
        mux := http.NewServeMux()
        mux.Handle("/", logging(index()))

        port, ok := os.LookupEnv("PORT")
        if !ok {
                port = "9001"
        }

        addr := fmt.Sprintf(":%s", port)
        server := http.Server{
                Addr:         addr,
                Handler:      mux,
                ReadTimeout:  15 * time.Second,
                WriteTimeout: 15 * time.Second,
                IdleTimeout:  15 * time.Second,
        }
        log.Println("main: running simple server on port", port)
        if err := server.ListenAndServe(); err != nil {
                log.Fatalf("main: couldn't start simple server: %v\n", err)
        }
}

