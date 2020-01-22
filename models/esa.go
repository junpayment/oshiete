package models

type EsaAnswer struct {
	List []*struct {
		Name        string
		Description string
		URL         string
	}
}
