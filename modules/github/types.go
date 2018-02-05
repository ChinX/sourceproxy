package github

var (
	provider    = "github"
	authURL     = "https://github.com/login/oauth/authorize"
	tokenURL    = "https://github.com/login/oauth/access_token"
	userURL     = "https://api.github.com/user"
	reposURL    = "https://api.github.com/user/repos"
	branchesURL = "https://api.github.com/repos/%s/%s/branches"
	commitsURL  = "https://api.github.com/repos/%s/%s/commits"
	accessKey = "Authorization"
	tokenPrefix = "token "
)
