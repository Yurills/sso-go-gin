package register

type RegisterRequest struct {
	// ClientID    string `json:"client_id" binding:"required"`
	// RID 	   string `json:"rid" binding:"required"`
	Username string `json:"username" form:"usernane" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	
	// CRSFSes string `json:"csrf_ses" binding:"required"`
}

// type registerResponse struct {
// 	Username string `json:"username"`
// }
