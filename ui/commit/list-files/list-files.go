package listfiles

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    header string
    choices *Output
    cursor int
    selected map[int]string
}

type Output struct {
    Output []string
    Selected []string
}

func (o *Output) update(val []string) {
    o.Selected = val
}

func InitialListFiles(output *Output, header string) model {
    return model {
        header: header,
        choices: output,
        selected: make(map[int]string),
    }
}

func (m model) Init() tea.Cmd  {
    return nil 
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.choices.Output)-1 {
                m.cursor++
            }
        case " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = m.choices.Output[m.cursor]
            }
        case "enter":
            var selectedFiles []string
            for _, f := range m.selected {
                selectedFiles = append(selectedFiles, f)
            }
            m.choices.update(selectedFiles)
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    s := fmt.Sprintf("%s\n\n", m.header)

    for i, choice := range m.choices.Output {
        cursor := ""
        if m.cursor == i {
            cursor = ">"
        }

        checked := ""
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\n Press q to quit.\n"
    return s
}
