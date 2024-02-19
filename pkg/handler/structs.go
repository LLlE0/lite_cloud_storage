package handler

//struct to store credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//struct to store folder contents
type FolderContent struct {
	Fname []string `json:"files"`
}
