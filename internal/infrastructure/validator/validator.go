package validator

import go_playground "github.com/go-playground/validator/v10"

type (
	validator struct {
		validator *go_playground.Validate
		err       error
		msg       []string
	}

	Interface interface {
		Validate(interface{}) error
		Messages() []string
	}
)

func New() Interface {
	return &validator{validator: go_playground.New()}
}

func (g *validator) Validate(i interface{}) error {
	if len(g.msg) > 0 {
		g.msg = nil
	}

	g.err = g.validator.Struct(i)

	if g.err != nil {
		return g.err
	}

	return nil
}

func (g *validator) Messages() []string {
	if g.err != nil {
		for _, err := range g.err.(go_playground.ValidationErrors) {
			g.msg = append(g.msg, err.Error())
		}
	}

	return g.msg
}
