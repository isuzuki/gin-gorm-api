package main

type Item struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	hidden1     string
	Hidden2     string `json:"-"`
}
