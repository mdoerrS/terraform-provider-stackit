package vpcroutingtable

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/core"
	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/validate"
)

var (
	_ resource.Resource = &vpcRoutingTableResource{}
)

type Model struct {
	Id             types.String `tfsdk:"id"` // needed by TF
	Region         types.String `tfsdk:"region"`
	ProjectId      types.String `tfsdk:"project_id"`
	VpcId          types.String `tfsdk:"vpc_id"`
	RoutingTableId types.String `tfsdk:"routing_table_id"`

	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Labels        types.Map    `tfsdk:"labels"`
	DynamicRoutes types.Bool   `tfsdk:"dynamic_routes"`
	SystemRoutes  types.Bool   `tfsdk:"system_routes"`
}

// NewVpcRoutingTable is a helper function to simplify the provider implementation.
func NewVpcRoutingTableResource() resource.Resource {
	return &vpcRoutingTableResource{}
}

// routingTableResource is the resource implementation.
type vpcRoutingTableResource struct {
	// client       *iaas.APIClient
	// providerData core.ProviderData
}

// Metadata returns the resource type name.
func (v *vpcRoutingTableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_routing_table"
}

// Schema implements resource.Resource.
func (v *vpcRoutingTableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Routing table resource schema.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Terraform's internal resource ID. It is structured as \"`project_id`,`vpc_id`,`region`,`routing_table_id`\".",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
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
			"project_id": schema.StringAttribute{
				Description: "STACKIT project ID to which the routing table is associated.",
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
				Description: "The network area ID to which the routing table is associated.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"routing_table_id": schema.StringAttribute{
				Description: "The routing tables ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the routing table.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(127),
				},
			},
			"description": schema.StringAttribute{
				Description: "Description of the routing table.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"labels": schema.MapAttribute{
				Description: "Labels are key-value string pairs which can be attached to a resource container",
				ElementType: types.StringType,
				Optional:    true,
			},
			"dynamic_routes": schema.BoolAttribute{
				Description: "This controls whether dynamic routes are propagated to this routing table",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"system_routes": schema.BoolAttribute{
				Description: "This allows installation of automatic system routes for connectivity between projects in the same VPC.",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (v *vpcRoutingTableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error create VPC regional routing table", "not implemented yet")
}

// Read refreshes the Terraform state with the latest data.
func (v *vpcRoutingTableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error read VPC regional routing table", "not implemented yet")
}

// Update updates the resource and sets the updated Terraform state on success.
func (v *vpcRoutingTableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error update VPC regional routing table", "not implemented yet")
}

// Delete deletes the resource and removes the Terraform state on success.
func (v *vpcRoutingTableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error delete VPC regional routing table", "not implemented yet")
}
