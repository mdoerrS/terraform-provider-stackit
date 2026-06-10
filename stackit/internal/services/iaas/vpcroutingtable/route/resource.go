package route

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	// iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2alpha1api"

	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/core"
	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/services/iaas/vpcroutingtable/shared"
	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/utils"
	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/validate"
)

var (
	_ resource.Resource = &vpcRoutingTableRouteResource{}
)

type Model struct {
	Id             types.String `tfsdk:"id"` // needed by TF
	Region         types.String `tfsdk:"region"`
	ProjectId      types.String `tfsdk:"project_id"`
	VpcId          types.String `tfsdk:"vpc_id"`
	RoutingTableId types.String `tfsdk:"routing_table_id"`

	*shared.StaticRouteModel // XXX: name of route model
}

// NewVpcRoutingTableRouteResource is a helper function to simplify the provider implementation.
func NewVpcRoutingTableRouteResource() resource.Resource {
	return &vpcRoutingTableRouteResource{}
}

// vpcRoutingTableRouteResource is the resource implementation.
type vpcRoutingTableRouteResource struct {
	// client       *iaas.APIClient
	// providerData core.ProviderData
}

// Metadata returns the resource type name.
func (v *vpcRoutingTableRouteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_routing_table_route"
}

// Schema implements resource.Resource.
func (v *vpcRoutingTableRouteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "VPC routing table route resource schema.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Terraform's internal resource ID. It is structured as \"`project_id`,`vpc_id`,`region`,`routing_table_id`,`route_id`\".",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				Description: "STACKIT project ID to which the VPC is associated.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"vpc_id": schema.StringAttribute{
				Description: "The VPC ID to which the route is associated.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"region": schema.StringAttribute{
				Description: "The resource region. If not defined, the provider region is used.",
				Optional:    true,
				// must be computed to allow for storing the override value from the provider
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"routing_table_id": schema.StringAttribute{
				Description: "The routing table ID of the regional routing table route.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"route_id": schema.StringAttribute{
				Description: "The ID of the regional routing table route.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"next_hop": schema.SingleNestedAttribute{
				Description: "Next hop destination.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: fmt.Sprintf("Type of the next hop. %s %s", utils.FormatPossibleValues("blackhole", "internet", "ipv4", "ipv6"), "Only `ipv4` supported currently."),
						Required:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"value": schema.StringAttribute{
						Description: "Either IPv4 or IPv6 (not set for blackhole and internet). Only IPv4 supported currently.",
						Optional:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Validators: []validator.String{
							validate.IP(false),
						},
					},
				},
			},
			"destination": schema.SingleNestedAttribute{
				Description: "Destination of the regional routing table route.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: fmt.Sprintf("CIDRV type. %s %s", utils.FormatPossibleValues("cidrv4", "cidrv6"), "Only `cidrv4` is supported currently."),
						Required:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"value": schema.StringAttribute{
						Description: "A CIDR string.",
						Required:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Validators: []validator.String{
							validate.CIDR(),
						},
					},
				},
			},
			"labels": schema.MapAttribute{
				Description: "Labels are key-value string pairs, which can be attached to a resource container",
				ElementType: types.StringType,
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Date-time when the regional routing table route was created.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Date-time when the regional routing table route was last updated.",
				Computed:    true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (v *vpcRoutingTableRouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error create VPC regional routing table route", "not implemented yet")
}

// Read refreshes the Terraform state with the latest data.
func (v *vpcRoutingTableRouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error read VPC regional routing table route", "not implemented yet")
}

// Update updates the resource and sets the updated Terraform state on success.
func (v *vpcRoutingTableRouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error update VPC regional routing table route", "not implemented yet")
}

// Delete deletes the resource and removes the Terraform state on success.
func (v *vpcRoutingTableRouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error delete VPC regional routing table route", "not implemented yet")
}
