package shared

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// XXX: name of route model
type StaticRouteModel struct {
	RouteId     types.String       `tfsdk:"route_id"`
	Destination *DestinationModel  `tfsdk:"destination"`
	Nexthop     *RouteNexthopModel `tfsdk:"next_hop"`
	Labels      types.Map          `tfsdk:"labels"`
	CreatedAt   types.String       `tfsdk:"created_at"`
	UpdatedAt   types.String       `tfsdk:"updated_at"`
}

type RouteNexthopModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type DestinationModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}
