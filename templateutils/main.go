package templateutils

type Register struct {
	data map[interface{}]interface{}
}

func NewRegister() *Register {
	return &Register{map[interface{}]interface{}{}}
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
