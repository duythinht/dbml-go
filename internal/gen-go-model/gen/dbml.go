package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/duythinht/dbml-go/core"
	"github.com/duythinht/dbml-go/parser"
	"github.com/duythinht/dbml-go/scanner"
)

const (
	dbdiagramURLPattern    = `https://dbdiagram.io/d/(\w+)`
	dbdiagramAPIURLPattern = `https://api.dbdiagram.io/query/%s`
)

var re = regexp.MustCompile(dbdiagramURLPattern)

func parseDBML(from string, recursive bool) (ret []*core.DBML) {
	files := collectFiles(from, recursive)
	for _, f := range files {
		r, err := dbmlReader(f)
		if err != nil {
			fmt.Printf("Error read file %s: %s", f, err)
			continue
		}

		p := parser.NewParser(scanner.NewScanner(r))
		dbml, err := p.Parse()
		if err != nil {
			fmt.Printf("Error parse file %s: %s", f, err)
			continue
		}
		ret = append(ret, dbml)
	}
	return
}

func collectFiles(from string, recursive bool) []string {
	stat, err := os.Stat(from)
	if err != nil {
		fmt.Printf("Invalid from parameter %s", err)
		return []string{}
	}

	// single file
	if !stat.IsDir() {
		return []string{from}
	}

	// directory
	files := []string{}
	filepath.Walk(from, func(path string, info os.FileInfo, err error) error {
		if path != from && info.IsDir() && !recursive {
			return filepath.SkipDir
		}
		files = append(files, path)
		return nil
	})
	return files
}

func dbmlReader(from string) (io.Reader, error) {
	m := re.FindStringSubmatch(from)

	if len(m) < 1 {
		return os.Open(from)
	}

	dbdURL := fmt.Sprintf(dbdiagramAPIURLPattern, m[1])
	resp, err := http.Get(dbdURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s NOT FOUND", dbdURL)
	}
	bodyJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := new(struct {
		ID      string `json:"_id"`
		Content string `json:"content"`
	})
	err = json.Unmarshal(bodyJSON, body)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(body.Content), nil
}
