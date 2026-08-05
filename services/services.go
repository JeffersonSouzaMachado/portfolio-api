package services

import "portfolio-api/projects"

func GetProjects(language string) []projects.ProjectResponse {

	switch language {
	case "pt":
		return projects.ProjectsPortuguese()

	case "en":
		return projects.ProjectsEnglish()

	default:
		return projects.ProjectsEnglish()
	}

}
