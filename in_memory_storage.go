package main

import "errors"

type InMemoryStorage struct {
	recipes map[string]Recipe
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{map[string]Recipe{}}
}

func (i *InMemoryStorage) Add(name string, recipe Recipe) error {
	_, ok := i.recipes[name]
	if ok {
		return errors.New("item already exist")
	}
	i.recipes[name] = recipe
	return nil
}

func (i *InMemoryStorage) Get(name string) (Recipe, error) {
	recipe, ok := i.recipes[name]
	if !ok {
		return Recipe{}, errors.New("not found")
	}
	return recipe, nil
}

func (i *InMemoryStorage) List() (map[string]Recipe, error) {
	return i.recipes, nil
}

func (i *InMemoryStorage) Update(name string, recipe Recipe) error {
	_, err := i.Get(name)
	if err != nil {
		return errors.New("no such item")
	}
	i.recipes[name] = recipe
	return nil
}

func (i *InMemoryStorage) Remove(name string) error {
	_, err := i.Get(name)
	if err != nil {
		return errors.New("not found")
	}
	delete(i.recipes, name)
	return nil
}
