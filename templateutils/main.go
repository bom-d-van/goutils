package templateutils

import (
	htmltmpl "html/template"
	"text/template"
)

type Register struct {
	data map[interface{}]interface{}
}

func NewRegister() *Register {
	return &Register{map[interface{}]interface{}{}}
}

// AddRegisterFunc provides an easy to add function for getting a Register to your text/template.Template
func AddRegisterFunc(tmpl *template.Template, name string) *template.Template {
	if name == "" {
		name = "getregister"
	}
	tmpl.Funcs(map[string]interface{}{
		name: func() *Register {
			return NewRegister()
		},
	})
	return tmpl
}

// AddRegisterFunc provides an easy to add function for getting a Register to your html/template.Template
func AddHtmlRegisterFunc(tmpl *htmltmpl.Template, name string) *htmltmpl.Template {
	if name == "" {
		name = "getregister"
	}
	tmpl.Funcs(map[string]interface{}{
		name: func() *Register {
			return NewRegister()
		},
	})
	return tmpl
}

func (r *Register) Set(key, val interface{}) string {
	r.data[key] = val
	return ""
}

func (r *Register) Get(key interface{}) interface{} {
	return r.data[key]
}

func (r *Register) GetString(key interface{}) string {
	return r.data[key].(string)
}

func (r *Register) GetFloat32(key interface{}) float32 {
	return r.data[key].(float32)
}

func (r *Register) GetFloat64(key interface{}) float64 {
	return r.data[key].(float64)
}

func (r *Register) GetInt(key interface{}) int {
	return r.data[key].(int)
}

func (r *Register) GetInt64(key interface{}) int64 {
	return r.data[key].(int64)
}

func (r *Register) GetRune(key interface{}) rune {
	return r.data[key].(rune)
}

func (r *Register) GetByte(key interface{}) byte {
	return r.data[key].(byte)
}

func (r *Register) GetBytes(key interface{}) []byte {
	return r.data[key].([]byte)
}
