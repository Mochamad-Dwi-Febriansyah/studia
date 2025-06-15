package http

type CreateJurnalRequest struct {
	Activity string `json:"activity" binding:"required"`
	Description string 	`json:"description" binding:"required"`
}
 
type UpdateJurnalRequest struct {
    Activity    *string `json:"activity"`
    Description *string `json:"description"`
}

// Anda juga bisa menambahkan struct untuk response di sini jika diperlukan.
// Contoh:
// type JurnalResponse struct {
//     ID          uuid.UUID `json:"id"`
//     Activity    string    `json:"activity"`
//     Description string    `json:"description"`
//     Status      string    `json:"status"`
//     CreatedAt   time.Time `json:"created_at"`
// }