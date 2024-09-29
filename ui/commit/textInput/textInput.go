package textInput

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Output struct {
    Output string
}

func (o *Output) update(val string) {
    o.Output = val
}

type (
	errMsg error
)

type model struct {
    header string
	textInput textinput.Model
    output *Output
	err       error
}

func InitialModel(output *Output, header string) model {
	ti := textinput.New()
	ti.Placeholder = "Please enter a commit message."
	ti.Focus()
	ti.CharLimit = 156

	return model{
        header: header,
		textInput: ti,
		err:       nil,
        output: output,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
            m.output.update(m.textInput.Value())
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
        m.header,
		m.textInput.View(),
		"(enter to commit)",
	) + "\n"
}
