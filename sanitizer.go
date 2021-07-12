package sanitizer

import (
	"errors"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"reflect"
)

//Sanitizer is HTML sanitizer
type Sanitizer struct {
	TagName string
	policy  map[string]*bluemonday.Policy
}

// Field sanitizing HTML from string|[]string|[]*string
// val must be a pointer
// tag must be strict|html
//   strict remove all tags from string
//   html   remove unsafe HTML tags and attributes
func (s *Sanitizer) Field(val interface{}, tag string) error {
	value := reflect.ValueOf(val)
	if value.Kind() != reflect.Ptr {
		return errors.New("interface should be a pointer")
	}
	return s.field(value, tag)
}

func (s *Sanitizer) field(value reflect.Value, tag string) error {
	el := value
	if el.Kind() == reflect.Ptr {
		el = value.Elem()
	}
	if !el.CanSet() {
		return nil
	}
	switch el.Kind() {
	case reflect.String:
		v, err := s.stringField(el.String(), tag)
		if err != nil {
			return err
		}
		el.SetString(v)
	case reflect.Slice, reflect.Array:
		count := el.Len()
		for i := 0; i < count; i++ {
			row := el.Index(i)
			if row.Kind() == reflect.Ptr || row.Kind() == reflect.String {
				err := s.field(row, tag)
				if err != nil {
					return fmt.Errorf("[%d] %w", i, err)
				}
			}
		}
	}
	return nil
}

func (s *Sanitizer) stringField(val string, tag string) (string, error) {
	p, ok := s.policy[tag]
	if !ok {
		return "", errors.New("undefined tag: " + tag)
	}
	return p.Sanitize(val), nil
}

//Struct sanitizing HTML from struct|[]struct|[]*struct
// obj must be a pointer
// You can specify which fields will be cleared using the tag
//type Additional struct{
//	Value string `sanitize:"strict"`
//}
//type Example struct {
//	Title string `sanitize:"strict"`
//	Description []string `sanitize:"html"`
//	Value Additional `sanitize:"dive"`
//}
// All HTML tags and attributes will be removed from the Title field
// All unsafe HTML tags and attributes will be removed from each element of the Description field
// All HTML tags and attributes will be removed from the Value.Value field
func (s *Sanitizer) Struct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	if value.Kind() != reflect.Ptr {
		return errors.New("interface should be a pointer")
	}
	return s.structValue(value)
}

func (s *Sanitizer) structValue(value reflect.Value) error {
	elem := value
	if elem.Kind() == reflect.Ptr {
		elem = value.Elem()
	}
	if !elem.CanSet() {
		return nil
	}
	switch elem.Kind() {
	case reflect.Struct:
		return s.sanitizeStruct(elem)
	case reflect.Array, reflect.Slice:
		count := elem.Len()
		for i := 0; i < count; i++ {
			row := elem.Index(i)
			if row.Kind() == reflect.Ptr || row.Kind() == reflect.Struct {
				err := s.structValue(row)
				if err != nil {
					return fmt.Errorf("[%d] %w", i, err)
				}
			}

		}
	}
	return nil
}

func (s *Sanitizer) sanitizeStruct(val reflect.Value) error {
	el := val
	if el.Kind() == reflect.Ptr {
		el = val.Elem()
	}
	if !el.CanSet() {
		return nil
	}
	t := reflect.TypeOf(el.Interface())
	count := el.NumField()
	for i := 0; i < count; i++ {
		field := el.Field(i)
		if !field.CanSet() {
			continue
		}
		ft := t.Field(i)
		tag := ft.Tag.Get(s.TagName)
		if tag == "dive" || ft.Anonymous {
			err := s.structValue(field)
			if err != nil {
				return errors.New("field " + t.Name() + ": " + err.Error())
			}
			continue
		}
		if tag == "" {
			continue
		}
		err := s.field(field, tag)
		if err != nil {
			return errors.New("field " + t.Name() + ": " + err.Error())
		}
	}
	return nil
}

//SetPolicy add new policy
//How to create new policy: see https://pkg.go.dev/github.com/microcosm-cc/bluemonday
func (s *Sanitizer) SetPolicy(name string, policy *bluemonday.Policy) {
	s.policy[name] = policy
}

//NewSanitizer creating HTML sanitizer
//NOTE: - this method is not thread-safe
//   Do this once, and use for the life of the program
func NewSanitizer() *Sanitizer {
	return &Sanitizer{
		policy: map[string]*bluemonday.Policy{
			"strict": bluemonday.StrictPolicy(),
			"html":   bluemonday.UGCPolicy(),
		},
		TagName: "sanitize",
	}
}
