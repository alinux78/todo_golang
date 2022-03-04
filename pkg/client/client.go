package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alinux78/todo/pkg/model"
)

type Client interface {
	GetAll() ([]model.TodoItem, error)
	Create(item model.TodoItemInput) (*model.TodoItem, error)
	//DISCUSS add new method
}

func NewClient(addr string) Client {
	return &clientImpl{
		addr: addr,
	}
}

type clientImpl struct {
	addr string
}

func (c *clientImpl) GetAll() ([]model.TodoItem, error) {
	endpoint := fmt.Sprintf("%s/items", c.addr)
	response, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(responseData))

	items := []model.TodoItem{}

	err = json.Unmarshal(responseData, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *clientImpl) Create(item model.TodoItemInput) (*model.TodoItem, error) {
	endpoint := fmt.Sprintf("%s/items", c.addr)
	jsonReq, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(endpoint, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(responseData))

	createdItem := model.TodoItem{}

	err = json.Unmarshal(responseData, &createdItem)
	if err != nil {
		return nil, err
	}

	return &createdItem, nil
}
