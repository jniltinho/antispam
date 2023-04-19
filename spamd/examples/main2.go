package main

// StatusCode StatusCode
// StatusMsg  string
// Version    string
// Score      float64
// BaseScore  float64
// IsSpam     bool
// Headers    textproto.MIMEHeader
// Msg        *Msg
// Rules      []map[string]string

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/jniltinho/antispam/spamd"
	"github.com/jniltinho/antispam/spamd/response"
	flag "github.com/spf13/pflag"
)

var (
	cfg     *Config
	cmdName string
)

// Config holds the configuration
type Config struct {
	Address        string
	Protocol       string
	User           string
	UseCompression bool
	Filename       string
}

func d(r *response.Response) {

	println(strings.Repeat("==", 85))
	log.Printf("RequestMethod => %v\n", r.RequestMethod)
	log.Printf("StatusCode => %v\n", r.StatusCode)
	log.Printf("StatusMsg => %v\n", r.StatusMsg)
	log.Printf("Version => %v\n", r.Version)
	log.Printf("Score => %v\n", r.Score)
	log.Printf("BaseScore => %v\n", r.BaseScore)
	log.Printf("IsSpam => %v\n", r.IsSpam)
	log.Printf("Headers => %v\n", r.Headers)
	log.Printf("Msg => %v\n", r.Msg)
	log.Printf("Msg.Header => %v\n", r.Msg.Header)
	log.Printf("Msg.Body => %s", r.Msg.Body)
	log.Printf("Msg.Raw => %s", r.Raw)
	log.Printf("Rules => %v\n", r.Rules)
	println(strings.Repeat("==", 85))

}

func init() {
	cfg = &Config{}
	cmdName = path.Base(os.Args[0])

	flag.StringVarP(&cfg.Protocol, "proto", "S", "unix", "Protocol family (unix or tcp)")
	flag.StringVarP(&cfg.Address, "addr", "H", "127.0.0.1:783", "Bind to address or unix domain socket")
	flag.StringVarP(&cfg.User, "user", "U", "clamav", "User for spamd to process this message under.")
	flag.BoolVarP(&cfg.UseCompression, "use-compression", "Z", false, "Compress mail message sent to spamd.")
	flag.StringVarP(&cfg.Filename, "file-name", "F", "", "Specify filename .eml file to scan.")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", cmdName)
	fmt.Fprint(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.ErrHelp = errors.New("")
	flag.CommandLine.SortFlags = false
	flag.Parse()
	ctx := context.Background()
	m := []byte("Subject: Hello\r\n\r\nHey there!\r\n")

	if cfg.Filename != "" {
		var err error
		m, err = os.ReadFile(cfg.Filename)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if cfg.Protocol == "unix" && cfg.Address == "127.0.0.1:783" {
		cfg.Address = "/var/run/clamav/spamd-socket"
	}
	c, err := spamd.NewClient(cfg.Protocol, cfg.Address, cfg.User, cfg.UseCompression)
	if err != nil {
		log.Println(err)
		return
	}
	c.EnableRawBody()
	ir := bytes.NewReader(m)
	r, e := c.Symbols(ctx, ir)
	if e != nil {
		log.Println(e)
		return
	}
	d(r)
}
