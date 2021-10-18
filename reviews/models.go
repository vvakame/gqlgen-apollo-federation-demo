package accounts

type Review struct {
	ID         string `json:"id"`
	AuthorID   string `json:"authorID"`
	ProductUPC string `json:"productUPC"`
	Body       string `json:"body"`
}

func (Review) IsEntity() {}
