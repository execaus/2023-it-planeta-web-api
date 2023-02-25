package models

type GetAnimalTypeOutput struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

type CreateAnimalTypeInput struct {
	Type string `json:"type" binding:"required,excludesall=' ',printascii"`
}

type CreateAnimalTypeOutput struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}
