package sanitizer

import (
	"github.com/microcosm-cc/bluemonday"
	"reflect"
	"testing"
)

func TestNewSanitizer(t *testing.T) {
	t.Run("NewSanitizer", func(t *testing.T) {
		got := NewSanitizer()
		if got == nil {
			t.Errorf("NewSanitizer() expected: pointer, got: nil")
			return
		}
	})
}

func TestSanitizer_Field(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		val interface{}
		tag string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sanitizer{
				TagName: tt.fields.TagName,
				policy:  tt.fields.policy,
			}
			if err := s.Field(tt.args.val, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("Field() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSanitizer_SetPolicy(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		name   string
		policy *bluemonday.Policy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//s := &Sanitizer{
			//	TagName: tt.fields.TagName,
			//	policy:  tt.fields.policy,
			//}
		})
	}
}

func TestSanitizer_Struct(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sanitizer{
				TagName: tt.fields.TagName,
				policy:  tt.fields.policy,
			}
			if err := s.Struct(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSanitizer_field(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		value reflect.Value
		tag   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sanitizer{
				TagName: tt.fields.TagName,
				policy:  tt.fields.policy,
			}
			if err := s.field(tt.args.value, tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("field() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSanitizer_sanitizeStruct(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sanitizer{
				TagName: tt.fields.TagName,
				policy:  tt.fields.policy,
			}
			if err := s.sanitizeStruct(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("sanitizeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSanitizer_stringField(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		val string
		tag string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sanitizer{
				TagName: tt.fields.TagName,
				policy:  tt.fields.policy,
			}
			got, err := s.stringField(tt.args.val, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("stringField() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizer_structValue(t *testing.T) {
	type fields struct {
		TagName string
		policy  map[string]*bluemonday.Policy
	}
	type args struct {
		value reflect.Value
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sanitizer{
				TagName: tt.fields.TagName,
				policy:  tt.fields.policy,
			}
			if err := s.structValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("structValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
