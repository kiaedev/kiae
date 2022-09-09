// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Container struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Status       string `json:"status"`
	Version      string `json:"version"`
	RestartCount int    `json:"restartCount"`
	CreatedAt    string `json:"createdAt"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Pod struct {
	Name         string       `json:"name"`
	Namespace    string       `json:"namespace"`
	Containers   []*Container `json:"containers"`
	RestartCount int          `json:"restartCount"`
	Status       string       `json:"status"`
	PodIP        string       `json:"podIP"`
	NodeIP       string       `json:"nodeIP"`
	CreatedAt    string       `json:"createdAt"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}