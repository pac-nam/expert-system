package structures

type Rule struct {
	Premice		[]byte
	Conclusion	[]byte
	Used		bool
}

func (r Rule)String() string {
	return string(r.Premice) + "=>" + string(r.Conclusion)
}