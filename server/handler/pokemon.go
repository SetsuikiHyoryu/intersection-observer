package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Pokemon struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Picture string `json:"picture"`
}

func GetPokemons(response http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	page := queryParams.Get("page")
	offset, err := strconv.Atoi(page)
	if err != nil {
		http.Error(response, "转换页数失败：%v", http.StatusBadRequest)
		return
	}

	pokemonUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?offset=%d&limit=20", offset)
	result, err := http.Get(pokemonUrl)
	if err != nil {
		http.Error(response, "获取宝可梦失败：%v", http.StatusInternalServerError)
		return
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		http.Error(response, "读取宝可梦响应体失败：%v", http.StatusInternalServerError)
		return
	}

	pokemons := &struct {
		Count   int       `json:"count"`
		Results []Pokemon `json:"results"`
	}{}

	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		http.Error(response, "反序列化宝可梦响应体失败：%v", http.StatusInternalServerError)
		return
	}

	// 如果 `for i, v := range` 这样去做，v 将是复制的，对其元素的修改将不影响响应所使用的原实例。
	for index := range pokemons.Results {
		pokemon := &pokemons.Results[index]

		// 从 https://pokeapi.co/api/v2/pokemon/1/ 中取出 1(id) 使用
		parsedUrl, err := url.Parse(pokemon.Url)
		if err != nil {
			http.Error(response, "宝可梦 URL 解析失败：%v", http.StatusInternalServerError)
			return

		}

		parts := strings.Split(parsedUrl.Path, "/") // /api/v2/pokemon/20/
		id := parts[4]                              // [ "", "api", "v2", "pokemon", "1" ]

		pokemon.Id = id
		pokemon.Picture = fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%s.png", id)
	}

	response.Header().Set("Content-Type", "application/json")
	responseBody, err := json.Marshal(pokemons)
	if err != nil {
		http.Error(response, "序列化宝可梦响应体失败：%v", http.StatusInternalServerError)
		return
	}

	response.Write(responseBody)
}
