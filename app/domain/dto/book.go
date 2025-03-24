package dto

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

type BookResponse struct {
	ID string `json:"id"`
}

type ListBooksResponseItem struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    ISBN   string `json:"isbn"`
}

type ListBooksResponse struct {
	Data []ListBooksResponseItem `json:"data"`
}
