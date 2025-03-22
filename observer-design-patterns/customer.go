package main

import "fmt"

type Observer interface {
	Update(champion Champion)
}

type Player struct {
	Name string
}

func (p *Player) Update(champion Champion) {
	fmt.Printf("%s nhận thông báo: %s đã chọn %s (vai trò: %s)\n",
		p.Name, champion.Name, champion.Name, champion.Role)
}
