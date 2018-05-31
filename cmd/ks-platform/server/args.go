package server

type PathArgs struct {
	Path string `form:"path" json:"path" binding:"required"`
}
