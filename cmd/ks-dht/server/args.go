package server

type HashArgs struct {
	Hash string `json:"hash" binding:"required"`
}
