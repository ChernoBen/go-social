package models

//representação da tabela Articles que guarda publicações de usuarios
type Articles struct {
	ID         uint64 `json:"id,omitempty"`
	Title      uint64 `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"author_id,omitempty"`
	AuthorNick uint64 `json:"author_nick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreatedAt  uint64 `json:"createdat,omitempty"`
}
