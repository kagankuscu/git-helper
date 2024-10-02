package listfiles

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
    dotChar = " * "
)

var (
    subtleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
    dotStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
)

type model struct {
    header string
    choices *Output
    cursor int
    selected map[int]string
    isSelected bool
    textInput textinput.Model
}

type Output struct {
    Output []string
    Selected []string
    Message string
}

func (o *Output) update(val []string) {
    o.Selected = val
}

func (o *Output) updateMessage(val string) {
    o.Message = val
}

func InitialListFiles(output *Output, header string) model {
	ti := textinput.New()
	ti.Placeholder = "Please enter a commit message."
	ti.Focus()
	ti.CharLimit = 156

    return model {
        header: header,
        choices: output,
        selected: make(map[int]string),
        isSelected: false,
        textInput: ti,
    }
}

func (m model) Init() tea.Cmd  {
    return nil 
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    if !m.isSelected {
        switch msg := msg.(type) {
        case tea.KeyMsg:
            switch msg.String() {
            case "ctrl+c", tea.KeyEsc.String():
                return m, tea.Quit
            case "up", "k":
                if m.cursor > 0 {
                    m.cursor--
                }
                return m, nil
            case "down", "j":
                if m.cursor < len(m.choices.Output)-1 {
                    m.cursor++
                }
                return m, nil
            case " ":
                _, ok := m.selected[m.cursor]
                if ok {
                    delete(m.selected, m.cursor)
                } else {
                    m.selected[m.cursor] = m.choices.Output[m.cursor]
                }
                return m, nil
            case "enter":
                var selectedFiles []string
                for _, f := range m.selected {
                    selectedFiles = append(selectedFiles, f)
                }
                m.choices.update(selectedFiles)
                if len(selectedFiles) > 0 {
                    m.isSelected = true
                }
                return m, nil
            }
        }
    }
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", tea.KeyEsc.String():
            return m, tea.Quit
        case "enter":
            m.choices.updateMessage(m.textInput.Value())
            return m, tea.Quit
        }
    }

    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m model) View() string {
    if m.isSelected {
        var selectedFiles string
        for i, file := range m.choices.Selected {
            selectedFiles += fmt.Sprintf("%d- %s\n", i + 1, file)
        }
        return fmt.Sprintf(
            "%s\n\n%s\n\n%s\n\n%s\n\n%s",
            "Selected files:",
            selectedFiles,
            "Please add a commit message",
            m.textInput.View(),
            subtleStyle.Render("enter: commit") + dotStyle + subtleStyle.Render("esc: quit"),
        ) + "\n"
    }

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

    s += "\n" + subtleStyle.Render("j/k, up/down: select" ) + dotStyle +
        subtleStyle.Render("space: choose" ) + dotStyle +
        subtleStyle.Render("enter: next") + dotStyle +
        subtleStyle.Render("esc: quit")
    return s
}
