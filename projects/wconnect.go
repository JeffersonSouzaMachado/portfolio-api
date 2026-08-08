package projects

func WconnectPortuguese() ProjectResponse {
	return ProjectResponse{

		ID:                 2,
		CardImage:          "pulsetrip_main",
		ShortCompanyName:   "WConnect",
		ShortDescription:   "O aplicativo Pulsetrip tem como finalidade automatizar processos para reservas de vôos, hoteis e veiculos.",
		CompanyFullName:    "Pulsetrip - WConnect",
		CompanyDescription: "A Pulse Trip nasceu a partir da necessidade de se preparar para o futuro e providenciar uma plataforma para dar ao cliente mais automação, simplificando os processos de viagens. Voltada para empresas, a plataforma permite a reserva de voos, hotéis e carros, com validação de cada locação pela própria empresa contratante, contemplando reservas nacionais e internacionais e oferecendo preços diferenciados negociados especialmente para o público corporativo.",

		AppOverview: "Atuei no desenvolvimento de novas funcionalidades da plataforma tanto para Web quanto para iOS e Android, utilizando Flutter Web com GetX como gerenciador de estado. Trabalhei na refatoração de código legado e na criação de telas e widgets personalizados a partir de protótipos no Figma, em colaboração direta com as equipes de UI/UX e Backend. Também atuei na integração com sistemas de busca de passagens aéreas e hotéis, contribuindo tanto para a evolução de funcionalidades existentes quanto para o desenvolvimento de novas features do produto.",

		AppChallenge: "Como o aplicativo era construído em Flutter, originalmente pensado para mobile, adaptar as telas para a versão Web trouxe desafios significativos de responsividade. Componentes e widgets que funcionavam bem em telas de smartphone não se comportavamento adequadamente em resoluções maiores, como desktops e tablets, gerando problemas de layout quebrado, elementos desproporcionais e experiência inconsistente entre as plataformas.",
		AppSolution:  "Desenvolvi widgets responsivos personalizados, utilizando LayoutBuilder e MediaQuery para adaptar dinamicamente o layout conforme o tamanho da tela, definindo breakpoints específicos para mobile, tablet e desktop. Refatorei componentes reutilizados em múltiplas telas para que se ajustassem de forma flexível, evitando duplicação de código entre as versões Web e mobile. Isso resultou em uma experiência mais consistente entre plataformas, mantendo uma única base de código Flutter Web para os três formatos.",
		TechStack:    WconnectProjectTechStack(),
		AppMockups:   WconnectAppMockups(),
	}
}

func WconnectEnglish() ProjectResponse {
	return ProjectResponse{

		ID:                 2,
		CardImage:          "pulsetrip_main",
		ShortCompanyName:   "WConnect",
		ShortDescription:   "The Pulsetrip app is designed to automate booking processes for flights, hotels, and vehicles.",
		CompanyFullName:    "Pulsetrip - WConnect",
		CompanyDescription: "Pulse Trip was born from the need to prepare for the future and provide a platform that gives clients more automation, simplifying travel processes. Built for companies, the platform allows booking flights, hotels, and cars, with each reservation validated by the hiring company, covering both domestic and international bookings, and offering special negotiated rates for corporate clients.",

		AppOverview: "I worked on developing new features for the platform across Web, iOS, and Android, using Flutter Web with GetX as the state management solution. I refactored legacy code and created custom screens and widgets based on Figma prototypes, working closely with the UI/UX and Backend teams. I also worked on integrating flight and hotel search systems, contributing both to the evolution of existing features and the development of new ones.",

		AppChallenge: "Since the app was originally built in Flutter with a mobile-first approach, adapting the screens for the Web version brought significant responsiveness challenges. Components and widgets that worked well on smartphone screens didn't behave properly at larger resolutions, such as desktops and tablets, leading to broken layouts, disproportionate elements, and an inconsistent experience across platforms.",
		AppSolution:  "I developed custom responsive widgets using LayoutBuilder and MediaQuery to dynamically adapt the layout based on screen size, defining specific breakpoints for mobile, tablet, and desktop. I refactored components reused across multiple screens so they could adjust flexibly, avoiding code duplication between the Web and mobile versions. This resulted in a more consistent experience across platforms while maintaining a single Flutter Web codebase for all three formats.",
		TechStack:    WconnectProjectTechStack(),
		AppMockups:   WconnectAppMockups(),
	}
}

func WconnectAppMockups() []string {
	return []string{
		"pulsetrip_1",
		"pulsetrip_2",
		"pulsetrip_3",
		"pulsetrip_4",
	}
}

func WconnectProjectTechStack() []TechStackResponse {
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
			Stack:     "BitBucket",
			AssetType: "icon",
		},
		{
			Icon:      "notepadText",
			Stack:     "Jira",
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
			Stack:     "Web",
			AssetType: "icon",
		},
	}
	
}
