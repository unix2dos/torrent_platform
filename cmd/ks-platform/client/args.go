package client

type HashArgs struct {
	Hash string `form:"hash" json:"hash" binding:"required"`
}
