package structures

type Rule struct {
	Premice		string
	Conclusion	string
}

func (r Rule) String() string {
	return r.Premice + "=>" + r.Conclusion
}
