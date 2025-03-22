package main

import "fmt"

type Match struct {
	observers         []Observer
	selectedChampions map[string]Champion
}

func NewMatch() *Match {
	return &Match{
		observers:         make([]Observer, 0),
		selectedChampions: make(map[string]Champion),
	}
}

func (m *Match) Attach(observer Observer) {
	m.observers = append(m.observers, observer)
}

func (m *Match) Detach(observer Observer) {
	for i, obs := range m.observers {
		if obs == observer {
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			break
		}
	}
}
func (m *Match) Notify(champion Champion) {
	for _, observer := range m.observers {
		observer.Update(champion)
	}
}

func (m *Match) SelectChampion(playerName string, champion Champion) {
	if _, exists := m.selectedChampions[champion.Name]; exists {
		fmt.Printf("Tướng %s đã được chọn trước đó!\n", champion.Name)
		return
	}
	m.selectedChampions[champion.Name] = champion
	fmt.Printf("%s đã khóa %s\n", playerName, champion.Name)
	m.Notify(champion)
}
