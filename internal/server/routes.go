package server

import (
	"github.com/diegogomesaraujo/fund-manager-api/internal/webresources"
)

type routes []webresources.Routes

var routesList = routes{
	// webresources.HomeRoutes,
	webresources.StocksRoutes,
}
