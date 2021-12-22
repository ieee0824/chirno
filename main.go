package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
)

type Config struct {
	BackendURL string   `json:"backend_url"`
	Command    []string `json:"command"`
	Port       string   `json:"port"`
}

type HelathChecker struct{}

func (h *HelathChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It works!")
}

func genProxy(cnf *Config) *httputil.ReverseProxy {
	if cnf == nil || cnf.BackendURL == "" {
		return nil
	}
	u, err := url.Parse(cnf.BackendURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	director := func(request *http.Request) {
		url := *request.URL
		url.Scheme = u.Scheme
		url.Host = u.Host
		req, err := http.NewRequest(request.Method, url.String(), request.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		req.Header = request.Header
		*request = *req
	}
	return &httputil.ReverseProxy{Director: director}
}

func runTargetProcess(cnf *Config) error {
	if cnf == nil {
		return fmt.Errorf("no command")
	}
	if len(cnf.Command) == 0 {
		return fmt.Errorf("no command")
	}

	var cmd *exec.Cmd
	if len(cnf.Command) == 1 {
		cmd = exec.Command(cnf.Command[0])
	} else {
		cmd = exec.Command(cnf.Command[0], cnf.Command[1:]...)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	fileName := flag.String("f", "./conf.json", "conf file name")
	flag.Parse()
	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalln("can not read conf.json")
	}
	cnf := &Config{}
	if err := json.NewDecoder(f).Decode(cnf); err != nil {
		log.Fatalln("can not read conf.json")
	}

	go func(cnf *Config) {
		if err := runTargetProcess(cnf); err != nil {
			log.Fatal(err.Error())
		}
		os.Exit(0)
	}(cnf)

	http.Handle("/health_check", &HelathChecker{})
	proxy := genProxy(cnf)
	if proxy != nil {
		http.Handle("/", proxy)
	}
	if err := http.ListenAndServe(":"+cnf.Port, nil); err != nil {
		log.Fatal(err.Error())
	}
}
