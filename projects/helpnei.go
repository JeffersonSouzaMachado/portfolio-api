package projects


func HelpneiPortuguese() ProjectResponse {
	return ProjectResponse{

		ID:                 4,
		CardImage:          "helpneiMainLogo",
		ShortCompanyName:   "Helpnei",
		ShortDescription:   "A Helpnei é uma plataforma de impacto social que conecta empresas patrocinadoras a pessoas em situação de vulnerabilidade econômica.",
		CompanyFullName:    "Helpnei +Renda para todos",
		CompanyDescription: "A Helpnei é uma plataforma de impacto social que conecta empresas patrocinadoras a pessoas em situação de vulnerabilidade econômica que desejam empreender. Por meio de um modelo de assinatura mensal, empresas patrocinam a formação de novos empreendedores, que recebem treinamento, suporte contínuo e acesso a uma loja online integrada a grandes marcas parceiras, como Magalu, Natura, Avon e Shopee.",
		AppOverview:        "Atuei no desenvolvimento do app desde a prototipagem, participando de todo o processo — da concepção criativa até o deploy nas lojas Apple Store e Play Store. Estive envolvido desde a primeira tela, com desenhos e fluxos criados no Miro e no Figma. No desenvolvimento, implementei autenticação, banco de dados e armazenamento com Firebase (Auth, Firestore e Database), além de funcionalidades como geração e envio automático de PDFs por e-mail e geolocalização. Também integrei o Sentry para monitoramento de erros e implementei o método de pagamento via Stark Bank. Posteriormente, conduzi a migração da infraestrutura para um backend próprio.",
		AppChallenge:       "Durante o cadastro, com autorização do usuário, coletávamos dados de geolocalização e informações do condomínio, armazenando-os em uma tabela para uso posterior na plataforma. Conforme a base de usuários cresceu, o volume de chamadas à API do Google Maps (usada para buscar localização e dados do local) aumentou proporcionalmente, elevando significativamente os custos operacionais e evidenciando a necessidade de repensar essa estratégia de integração.",
		AppSolution:        "Reestruturei o fluxo de busca em duas etapas: a cada novo cadastro, buscávamos apenas a latitude e longitude do usuário (gratuito). Com essas coordenadas, consultávamos uma tabela de cache com os condomínios já mapeados na região. Caso já existissem dados para aquele ponto, nenhuma nova chamada era feita. Caso contrário, buscávamos os condomínios num raio de 10 km — dado que já retornava em uma única chamada — e armazenávamos o resultado no cache para reutilização em cadastros futuros na mesma área.",
		TechStack:          HelpneiProjectTechStack(),
		AppMockups:         HelpneiAppMockups(),
	}
}

func HelpneiEnglish() ProjectResponse {
	return ProjectResponse{

		ID:                 4,
		CardImage:          "helpneiMainLogo",
		ShortCompanyName:   "Helpnei",
		ShortDescription:   "Helpnei is a social impact platform that connects sponsoring companies with people facing economic vulnerability.",
		CompanyFullName:    "Helpnei +Renda para todos",
		CompanyDescription: "Helpnei is a social impact platform that connects sponsoring companies with people in situations of economic vulnerability who want to become entrepreneurs. Through a monthly subscription model, companies sponsor the training of new entrepreneurs, who receive coaching, ongoing support, and access to an online store integrated with major partner brands such as Magalu, Natura, Avon, and Shopee.",
		AppOverview:        "I worked on the app's development from the prototyping stage, taking part in the entire process — from creative conception to deployment on the Apple App Store and Play Store. I was involved from the very first screen, with designs and flows created in Miro and Figma. In development, I implemented authentication, database, and storage using Firebase (Auth, Firestore, and Database), as well as features like automatic PDF generation and email delivery, and geolocation. I also integrated Sentry for error monitoring and implemented the payment method via Stark Bank. I later led the migration of the infrastructure to a proprietary backend.",
		AppChallenge:       "During registration, with the user's authorization, we collected geolocation data and building/condominium information, storing it in a table for later use on the platform. As the user base grew, the number of calls to the Google Maps API (used to look up location and site data) increased proportionally, significantly raising operational costs and highlighting the need to rethink this integration strategy.",
		AppSolution:        "I restructured the lookup flow into two steps: for each new registration, we would fetch only the user's latitude and longitude (free of charge). Using those coordinates, we would query a cache table of buildings/condominiums already mapped in the region. If data already existed for that point, no new call was made. Otherwise, we would search for buildings within a 10 km radius — a query that already returned in a single call — and store the result in the cache for reuse in future registrations in the same area.",
		TechStack:          HelpneiProjectTechStack(),
		AppMockups:         HelpneiAppMockups(),
	}
}

func HelpneiAppMockups() []string {
	return []string{
		"helpnei_1",
		"helpnei_2",
		"helpnei_3",
	}
}

func HelpneiProjectTechStack() []TechStackResponse {
	return []TechStackResponse{
		{
			Icon:      "code",
			Stack:     "Flutter/Dart",
			AssetType: "icon",
		},
		{
			Icon:      "database",
			Stack:     "Firebase",
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
			Icon:      "bug",
			Stack:     "Sentry",
			AssetType: "icon",
		},
		{
			Icon:      "cloudUpload",
			Stack:     "Codemagic",
			AssetType: "icon",
		},
		{
			Icon:      "notepadText",
			Stack:     "Jira",
			AssetType: "icon",
		},
		{
			Icon:      "penTool",
			Stack:     "Miro",
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
	}
}
