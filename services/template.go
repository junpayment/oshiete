package services

import (
	"bytes"
	"fmt"
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
{{ range $i, $v := .EsaAnswer.List }}
{{ $v.Name }}
{{ $v.URL }}
{{ end }}
{{ end }}
`
	tmp := template.New("answer")
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
