package documents

type CPF string

func (c CPF) String() string {
	return string(c)
}
