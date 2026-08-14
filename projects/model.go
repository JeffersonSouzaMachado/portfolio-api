package projects

type ProjectResponse struct {
	ID                 int                 `json:"id"`
	CardImage          string              `json:"cardImage"`
	ShortCompanyName   string              `json:"shortCompanyName"`
	ShortDescription   string              `json:"shortDescription"`
	CompanyFullName    string              `json:"companyFullName"`
	CompanyDescription string              `json:"companyDescription"`
	AppOverview        string              `json:"appOverview"`
	AppChallenge       string              `json:"appChallenge"`
	AppSolution        string              `json:"appSolution"`
	TechStack          []TechStackResponse `json:"techStack"`
	AppMockups         []string            `json:"appMockups"`
	ProjectInfo        ProjectInfoResponse `json:"projectInfo"`
}

type TechStackResponse struct {
	Icon      string `json:"icon"`
	Stack     string `json:"stack"`
	AssetType string `json:"assetType"`
}

type ProjectInfoResponse struct {
	Downloads    int16  `json:"downloads"`
	PlayStoreUrl string `json:"playStoreUrl"`
	AppStoreUrl  string `json:"appStoreUrl"`
	Rate         string `json:"rate"`
}
