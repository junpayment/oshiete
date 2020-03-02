package services

import (
	"bytes"
	"fmt"
	"github.com/junpayment/oshiete/models/iruka"
	"io/ioutil"
	"text/template"

	"github.com/junpayment/oshiete/models"
)

type Templete struct{}

func (s *Templete) OutEte(answer *models.EteAnswer) (string, error) {
	letter := `
検索結果です！
{{ .Answer }}
{{ $length := len .EsaAnswer.List }}{{ if gt $length 0 }}
esa検索結果
{{ $list := trimslice .EsaAnswer.List 3 }}
{{ range $i, $v := $list }}
{{ $v.Name }}
{{ $v.URL }}
{{ end }}
{{ end }}
`
	tmp := template.New("answer")
	tmp.Funcs(template.FuncMap{
		"trimslice": func(s []*struct {
			Name        string
			Description string
			URL         string
		}, end int) []*struct {
			Name        string
			Description string
			URL         string
		} {
			if len(s) < end {
				return s
			}
			return s[:end]
		},
	})
	tmp, err := tmp.Parse(letter)
	if err != nil {
		return "", fmt.Errorf(`tmp, err := tmp.Parse(letter): %w`, err)
	}
	w := &bytes.Buffer{}
	err = tmp.Execute(w, answer)
	if err != nil {
		return "", fmt.Errorf(`err := tmp.Execute(&buf, answer): %w`, err)
	}
	buf, err := ioutil.ReadAll(w)
	if err != nil {
		return "", fmt.Errorf(`buf, err := ioutil.ReadAll(w): %w`, err)
	}
	return string(buf), nil
}

func (s *Templete) OutIruka(states []*iruka.State) (string, error) {
	letter := `
在席状況です！
{{ range $i, $v := . }}
:{{ $v.Name }}: : {{ $v.Status }} : {{ $v.Message }}
{{ end }}
`
	tmp := template.New("answer")
	tmp, err := tmp.Parse(letter)
	if err != nil {
		return "", fmt.Errorf(`tmp, err := tmp.Parse(letter): %w`, err)
	}
	w := &bytes.Buffer{}
	err = tmp.Execute(w, states)
	if err != nil {
		return "", fmt.Errorf(`err := tmp.Execute(&buf, answer): %w`, err)
	}
	buf, err := ioutil.ReadAll(w)
	if err != nil {
		return "", fmt.Errorf(`buf, err := ioutil.ReadAll(w): %w`, err)
	}
	return string(buf), nil
}
