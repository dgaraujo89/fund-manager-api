package server

import (
	"github.com/diegogomesaraujo/fund-manager-api/internal/webresources"
)

var homeRoutes = webresources.HomeRoutes

type routes []webresources.Routes

var routesList = routes{
	homeRoutes,
}
