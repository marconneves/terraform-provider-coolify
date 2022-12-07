package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Item struct {
	Name		string
	Description string
	Tags		[]string
}


// GetAll Retrieves all of the Items from the server
func (c *Client) GetAll() (*map[string]Item, error) {
	body, err := c.httpRequest("item", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Item{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

// GetItem gets an item with a specific name from the server
func (c *Client) GetItem(name string) (*Item, error) {
	body, err := c.httpRequest(fmt.Sprintf("item/%v", name), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	item := &Item{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// NewItem creates a new Item
func (c *Client) NewItem(item *Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest("item", "POST", buf)
	if err != nil {
		return err
	}
	return nil
}

// UpdateItem updates the values of an item
func (c *Client) UpdateItem(item *Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("item/%s", item.Name), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem removes an item from the server
func (c *Client) DeleteItem(itemName string) error {
	_, err := c.httpRequest(fmt.Sprintf("item/%s", itemName), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}