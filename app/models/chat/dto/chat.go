package dto

type CreateDialogRequest struct {
	UserID       string   `validate:"required"`
	Name         string   `json:"name"`
	Participants []string `json:"participants" validate:"required"`
}

type CreateDialogResponse struct {
	DialogID string `json:"dialog_id"`
}

type GetDialogsRequest struct {
	UserID string `query:"id"`
}

type GetDialogsInfoResponseBody struct {
	DialogInfo []DialogInfo `json:"dialog_info"`
}

type Message struct {
	ID        string `json:"id"`
	DialogID  string `json:"dialog_id"`
	AuthorID  string `json:"author_id"`
	Body      string `json:"body"`
	CreatedAt int64  `json:"created_at"`
}

type GetDialogRequest struct {
	DialogID string `query:"id"`
}

type GetDialogResponse struct {
	DialogInfo  DialogInfo `json:"dialog"`
	Messages    []Message  `json:"messages"`
	Total       int64      `json:"total"`
	AmountPages int64      `json:"amount_pages"`
}

type DialogInfo struct {
	DialogID     string   `json:"dialog_id"`
	Name         string   `json:"name"`
	Participants []string `json:"participants"`
	CreatedAt    int64    `json:"created_at"`
}
