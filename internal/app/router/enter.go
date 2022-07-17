package router

import "template/internal/app/router/user"

var (
	FrontRouterGroupApp    = new(FrontRouterGroup)
	MerchantRouterGroupApp = new(MerchantRouterGroup)
)

type FrontRouterGroup struct {
	User user.FrontUserRouterGroup
}

type MerchantRouterGroup struct {
	User user.MerchantRouterGroup
}
