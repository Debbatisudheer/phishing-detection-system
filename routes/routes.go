package routes

func SetupRoutes() {

	RegisterAuthRoutes()

	RegisterAnalysisRoutes()

	RegisterDashboardRoutes()

	RegisterCaseRoutes()

	RegisterIncidentRoutes()

	RegisterCampaignRoutes()

	RegisterThreatIntelRoutes()

	RegisterSandboxRoutes()

	RegisterSystemRoutes()

	RegisterEmailRoutes()

	RegisterWebsocketRoutes()

}