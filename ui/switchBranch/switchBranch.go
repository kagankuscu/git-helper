package switchBranch

import (
	"fmt"
	"git-helper/utils"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
    dotChar = " * "
)

var (
    currentBranchStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
    subtleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
    dotStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
)

type tickMsg time.Time

type model struct {
	cursor   int
	choices  []string
	choice string
    current string
}

func InitialModel() model {
    out, err := exec.Command("git", "branch").Output()
    utils.CheckError(err)

    var branches []string
    var current string

    for _, branch := range strings.Split(strings.TrimSpace(string(out)), "\n") {
        if !strings.Contains(branch, "*") {
            branches = append(branches, strings.TrimSpace(branch))
        } else {
            current = strings.Split(branch, " ")[1]
        }
    }

	return model{
		choices: branches,
        current: current,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Switch Branch")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
            return m, nil
        case "enter":
		    m.choice = m.choices[m.cursor]	
            err := exec.Command("git", "checkout", m.choice).Run()
            utils.CheckError(err)
            return m, tick
		}
    case tickMsg:
        return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
    if m.choice != "" {
        return "Switched to branch " + fmt.Sprintf("'%s'\n", currentBranchStyle.Render(m.choice))
    }

	s := "Choose branch you want to switch.\n"
    s += "Current branch: " + currentBranchStyle.Render(m.current) + "\n\n"

	for i, choice := range m.choices {
		cursor := "( )"
		if m.cursor == i {
			cursor = "{•)"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

    s += "\n" + subtleStyle.Render("j/k, up/down: select" ) + dotStyle +
        subtleStyle.Render("enter: choice") + dotStyle +
        subtleStyle.Render("esc: quit")

	return s
}

func tick() tea.Msg {
    time.Sleep(2 * time.Second)
    return tickMsg{}
}
