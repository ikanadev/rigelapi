package handlers

type Teacher struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type School struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}

type Dpto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Provincia struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Municipio struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Grade struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Period struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Area struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type Year struct {
	ID    string `json:"id"`
	Value int    `json:"value"`
}

type Class struct {
	ID       string `json:"id"`
	Parallel string `json:"parallel"`
}
