package element

import (
	"fmt"
	"strings"
)

type Element struct {
	ID      string
	Session executable
}

type executable interface {
	Execute(endpoint, method string, body, result interface{}) error
}

func (e *Element) GetID() string {
	return e.ID
}

func (e *Element) GetText() (string, error) {
	var text string
	if err := e.Session.Execute(e.url()+"/text", "GET", nil, &text); err != nil {
		return "", err
	}
	return text, nil
}

func (e *Element) GetAttribute(attribute string) (string, error) {
	var value string
	if err := e.Session.Execute(fmt.Sprintf("%s/attribute/%s", e.url(), attribute), "GET", nil, &value); err != nil {
		return "", err
	}
	return value, nil
}

func (e *Element) GetCSS(property string) (string, error) {
	var value string
	if err := e.Session.Execute(fmt.Sprintf("%s/css/%s", e.url(), property), "GET", nil, &value); err != nil {
		return "", err
	}
	return value, nil
}

func (e *Element) Click() error {
	return e.Session.Execute(e.url()+"/click", "POST", nil, &struct{}{})
}

func (e *Element) Clear() error {
	return e.Session.Execute(e.url()+"/clear", "POST", nil, &struct{}{})
}

func (e *Element) Value(text string) error {
	splitText := strings.Split(text, "")
	request := struct {
		Value []string `json:"value"`
	}{splitText}
	return e.Session.Execute(e.url()+"/value", "POST", request, &struct{}{})
}

func (e *Element) IsSelected() (bool, error) {
	var selected bool
	if err := e.Session.Execute(e.url()+"/selected", "GET", nil, &selected); err != nil {
		return false, err
	}
	return selected, nil
}

func (e *Element) Submit() error {
	return e.Session.Execute(e.url()+"/submit", "POST", nil, &struct{}{})
}

func (e *Element) url() string {
	return "element/" + e.ID
}
