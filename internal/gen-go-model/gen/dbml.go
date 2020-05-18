package gen

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func parseDBML(from string) (*core.DBML, error) {
	r, err := dbmlReader(from)
	if err != nil {
		return nil, err
	}
	p := parser.NewParser(scanner.NewScanner(r))
	return p.Parse()
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
