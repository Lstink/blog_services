package model

type ArticleTag struct {
	*Model
	ArticleId string `json:"article_id"`
	TagId     uint8  `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
