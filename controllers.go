package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func GetPokemon(identifier string) (Pokemon, error) {

	var pokemon Pokemon
	var errorPokemon error = nil

	c := colly.NewCollector()

	c.OnXML(`/html/body/div[4]/section[1]/div[2]/div/span`, func(e *colly.XMLElement) {
		pokemon.Number = strings.Replace(e.Text, "N.º", "", -1)
	})

	c.OnXML(`/html/body/div[4]/section[1]/div[2]/div`, func(e *colly.XMLElement) {
		pokemon.Name = strings.TrimSpace(strings.Split(strings.TrimSpace(e.Text), "N.º")[0])
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[1]/div[1]/div/img`, func(e *colly.XMLElement) {
		if e.Attr("class") == "active" {
			pokemon.ImageUrl = e.Attr("src")
		}
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[1]/p[2]`, func(e *colly.XMLElement) {
		pokemon.Description = strings.TrimSpace(e.Text)
	})

	//HEIGHT
	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[3]/div[1]/div[1]/ul/li[1]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Height = strings.TrimSpace(e.Text)
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[4]/div[1]/div[1]/ul/li[1]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Height = strings.TrimSpace(e.Text)
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[5]/div[1]/div[1]/ul/li[1]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Height = strings.TrimSpace(e.Text)
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[6]/div[1]/div[1]/ul/li[1]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Height = strings.TrimSpace(e.Text)
	})

	//WEIGHT
	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[3]/div[1]/div[1]/ul/li[2]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Weight = strings.TrimSpace(e.Text)
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[4]/div[1]/div[1]/ul/li[2]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Weight = strings.TrimSpace(e.Text)
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[5]/div[1]/div[1]/ul/li[2]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Weight = strings.TrimSpace(e.Text)
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[6]/div[1]/div[1]/ul/li[2]/span[2]`, func(e *colly.XMLElement) {
		pokemon.Weight = strings.TrimSpace(e.Text)
	})

	//HABILIDADES
	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[3]/div[1]/div[2]/ul/li[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Abilities = append(pokemon.Abilities, Ability{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[4]/div[1]/div[2]/ul/li[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Abilities = append(pokemon.Abilities, Ability{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[5]/div[1]/div[2]/ul/li[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Abilities = append(pokemon.Abilities, Ability{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[6]/div[1]/div[2]/ul/li[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Abilities = append(pokemon.Abilities, Ability{
			Name: strings.TrimSpace(e.Text),
		})
	})

	//TIPOS
	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[3]/div[2]/div[1]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Types = append(pokemon.Types, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[4]/div[2]/div[1]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Types = append(pokemon.Types, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[5]/div[2]/div[1]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Types = append(pokemon.Types, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[6]/div[2]/div[1]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Types = append(pokemon.Types, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	//DEBILIDADES
	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[3]/div[2]/div[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Weaknesses = append(pokemon.Weaknesses, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[4]/div[2]/div[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Weaknesses = append(pokemon.Weaknesses, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[5]/div[6]/div[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Weaknesses = append(pokemon.Weaknesses, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnXML(`/html/body/div[4]/section[3]/div[2]/div/div[6]/div[2]/div[2]/ul/li`, func(e *colly.XMLElement) {
		pokemon.Weaknesses = append(pokemon.Weaknesses, Type{
			Name: strings.TrimSpace(e.Text),
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		errorPokemon = err
	})

	c.Visit("https://www.pokemon.com/el/pokedex/" + identifier)

	if errorPokemon != nil {
		return Pokemon{}, errorPokemon
	}

	return pokemon, nil
}
