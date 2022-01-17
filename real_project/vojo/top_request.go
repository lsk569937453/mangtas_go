package vojo

type TopRequst struct {
	SourceText *string `form:"sourceText" json:"sourceText" binding:"required"`
}
