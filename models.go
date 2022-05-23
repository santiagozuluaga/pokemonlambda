package main

type Ability struct {
	Name string `json:"name"`
}

type Type struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Number      string    `json:"number"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	Types       []Type    `json:"types"`
	Weaknesses  []Type    `json:"weaknesses"`
	Weight      string    `json:"weight"`
	Height      string    `json:"height"`
	Abilities   []Ability `json:"abilities"`
}
