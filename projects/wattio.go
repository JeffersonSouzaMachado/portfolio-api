package projects

func WattioPortuguese() ProjectResponse {
	return ProjectResponse{

		ID:                 2,
		CardImage:          "wattio_main",
		ShortCompanyName:   "Wattio",
		CompanyFullName:    "Wattio",
		ShortDescription:   "A Wattio é uma plataforma completa de gestão e automação de usinas em Geração Distribuída, oferecendo aos clientes finais um app para acompanhar créditos de energia e economia na conta de luz.",
		CompanyDescription: "A Wattio nasceu para facilitar o crescimento da Energia por Assinatura no Brasil, oferecendo uma solução completa que integra e automatiza toda a operação de usinas na Geração Distribuída. A plataforma atende geradoras, integradores e desenvolvedores, cobrindo todo o fluxo operacional: simulação e adesão de clientes, distribuição de créditos em conformidade com a Lei 14.300, leitura e auditoria automatizada de contas de energia junto a mais de 20 distribuidoras do país, rateio de despesas e geração de faturas com assinatura digital.",

		AppOverview: "Atuei como Mobile Developer na Wattio, sendo responsável pela unificação de 10 aplicativos em uma única base de código Flutter, além de conduzir toda a esteira de CI/CD do projeto com GitHub Actions e AWS Amplify. Implementei o gerenciamento de estado com GetX, criei um mecanismo de force-update para controlar versões em produção e apliquei testes automatizados para aumentar a confiabilidade das entregas. Também documentei a arquitetura da aplicação para apoiar o onboarding do time e trabalhei em conjunto com áreas multidisciplinares ao longo de todo o desenvolvimento.",

		AppChallenge: "A empresa possuía 10 aplicativos White Label separados, um para cada cliente, todos com a mesma base funcional mas com pequenas customizações visuais e de regras de negócio. Isso gerava duplicação massiva de código, dificultava a manutenção, tornava a publicação de novas versões um processo lento e repetitivo, e aumentava significativamente o risco de inconsistências entre os apps, já que qualquer correção ou nova funcionalidade precisava ser replicada manualmente em cada um dos 10 repositórios.",

		AppSolution: "Liderei a migração dos 10 aplicativos para um único monorepo, utilizando Flutter Flavors para gerenciar as configurações específicas de cada cliente (ícones, cores, nomes e endpoints) e feature flags para controlar quais funcionalidades estariam disponíveis em cada versão. Implementei pipelines de CI/CD com GitHub Actions integrados ao AWS Amplify, automatizando o build e deploy de todas as flavors a partir de uma única base de código. Isso eliminou a duplicação de código, reduziu drasticamente o tempo de entrega de novas features e correções, e garantiu consistência entre todos os apps do portfólio.",
		TechStack:   WattioProjectTechStack(),
		AppMockups:  WattioAppMockups(),
	}
}

func WattioEnglish() ProjectResponse {
	return ProjectResponse{

		ID:                 2,
		CardImage:          "wattio_main",
		ShortCompanyName:   "Wattio",
		CompanyFullName:    "Wattio",
		ShortDescription:   "Wattio is a complete management and automation platform for Distributed Generation power plants, offering end customers an app to track energy credits and savings on their electricity bill.",
		CompanyDescription: "Wattio was created to accelerate the growth of Energy by Subscription in Brazil, offering a complete solution that integrates and automates the entire operation of Distributed Generation power plants. The platform serves power plant operators, integrators, and developers, covering the entire operational flow: customer simulation and onboarding, credit distribution in compliance with Brazilian Law 14.300, automated reading and auditing of energy bills across more than 20 utility companies nationwide, expense apportionment, and invoice generation with digital signature.",

		AppOverview: "Worked as a Mobile Developer at Wattio, responsible for unifying 10 applications into a single Flutter codebase, as well as leading the project's entire CI/CD pipeline using GitHub Actions and AWS Amplify. Implemented state management with GetX, built a force-update mechanism to control app versions in production, and applied automated tests to improve delivery reliability. Also documented the application's architecture to support team onboarding and collaborated with cross-functional teams throughout development.",

		AppChallenge: "The company had 10 separate White Label applications, one for each client, all sharing the same core functionality but with small visual and business rule customizations. This generated massive code duplication, made maintenance difficult, turned the release of new versions into a slow and repetitive process, and significantly increased the risk of inconsistencies between apps, since any fix or new feature had to be manually replicated across each of the 10 repositories.",

		AppSolution: "Led the migration of the 10 applications into a single monorepo, using Flutter Flavors to manage client-specific configurations (icons, colors, names, and endpoints) and feature flags to control which functionalities were available in each version. Implemented CI/CD pipelines with GitHub Actions integrated with AWS Amplify, automating the build and deployment of all flavors from a single codebase. This eliminated code duplication, drastically reduced the delivery time for new features and fixes, and ensured consistency across the entire app portfolio.",
		TechStack:   WattioProjectTechStack(),
		AppMockups:  WattioAppMockups(),
	}
}

func WattioAppMockups() []string {
	return []string{
		"wattio_1",
		"wattio_2",
		"wattio_3",
		"wattio_4",
	}
}

func WattioProjectTechStack() []TechStackResponse {
	return []TechStackResponse{
		{
			Icon:      "code",
			Stack:     "Flutter/Dart",
			AssetType: "icon",
		},
		{
			Icon:      "database",
			Stack:     "AWS",
			AssetType: "icon",
		},
		{
			Icon:      "mapPin",
			Stack:     "Geo Location",
			AssetType: "icon",
		},
		{
			Icon:      "gitGraph",
			Stack:     "Github",
			AssetType: "icon",
		},
		{
			Icon:      "gitCommitHorizontal",
			Stack:     "AWS",
			AssetType: "icon",
		},
		{
			Icon:      "notepadText",
			Stack:     "Notion",
			AssetType: "icon",
		},
		{
			Icon:      "penTool",
			Stack:     "Figma",
			AssetType: "icon",
		},
		{
			Icon:      "playStore",
			Stack:     "Play Store",
			AssetType: "svg",
		},
		{
			Icon:      "appStore",
			Stack:     "App Store",
			AssetType: "svg",
		},
		{
			Icon:      "globeCheck",
			Stack:     "Amplify",
			AssetType: "icon",
		},
		{
			Icon:      "bookCopy",
			Stack:     "White Label",
			AssetType: "icon",
		},
		{
			Icon:      "calendarSync",
			Stack:     "Realtime",
			AssetType: "icon",
		},
		{
			Icon:      "globeOff",
			Stack:     "Offline First",
			AssetType: "icon",
		},
	}

}
