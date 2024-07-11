package main

type Recipe struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type recipeStore interface {
	Add(name string, recipe Recipe) error
	Get(name string) (Recipe, error)
	List() (map[string]Recipe, error)
	Update(name string, recipe Recipe) error
	Remove(name string) error
}
