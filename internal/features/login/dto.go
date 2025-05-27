package login

type LoginRequest struct {
	ClientID    string `json:"client_id" binding:"required"`
	RID 	   string `json:"rid" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CRSFSes string `json:"csrf_ses" binding:"required"`
}

type LoginResponse struct {
	


}