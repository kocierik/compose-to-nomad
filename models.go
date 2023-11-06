package models

type Job struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Group []Group `json:"group"`
	Task  []Task  `json:"task"`
}

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Task []Task `json:"task"`
}

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Driver string `json:"driver"`
	Config Config `json:"config"`
}

type Config struct {
	Image string   `json:"image"`
	Port  []string `json:"port"`
	Env   Env      `json:"env"`
}

type Env struct {
	Value []string `json:"value"`
}
