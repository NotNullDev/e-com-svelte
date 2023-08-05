package api

import "github.com/notnulldev/e-com-svelte/ent"

type AppApi struct {
	client *ent.Client
}

func NewAppApi(client *ent.Client) *AppApi {
	return &AppApi{
		client: client,
	}
}

type ErrorMessage struct {
	Message string `json:"message"`
}
