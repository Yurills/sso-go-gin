package login

type Service struct {
	//access user repo
	//create token maker
}

func (s *Service) Login(req LoginRequest) () {
	//1. verify required parameter
	//verify csrf_ses
	//verify auth_request_code not expired
	// verify username password

	//2. create code random string uniq

}