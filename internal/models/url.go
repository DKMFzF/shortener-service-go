package models

type MetaInfo struct {
	Title       string `json:"title" valid:"optional"`
	Description string `json:"description" valid:"optional"`
}

type UrlToShort struct {
	Url string `json:"url" valid:"url"`
}

type Request struct {
	Meta *MetaInfo   `json:"meta" valid:"required"`
	Data *UrlToShort `json:"data" valid:"required"`
}

type Response struct {
	Data    *UrlToShort `json:"data" valid:"required"`
	Version string      `json:"version" valid:"required"`
}
