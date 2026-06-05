package vpcnetworkrange

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	// iaas "github.com/stackitcloud/stackit-sdk-go/services/iaas/v2alpha1api"
	// resourcemanager "github.com/stackitcloud/stackit-sdk-go/services/resourcemanager/v0api"
	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/core"
	"github.com/stackitcloud/terraform-provider-stackit/stackit/internal/validate"
)

var (
	_ resource.Resource                   = &vpcNetworkRangeResource{}
	_ resource.ResourceWithValidateConfig = &vpcNetworkRangeResource{}
)

type Model struct {
	Id             types.String `tfsdk:"id"` // needed by TF
	Region         types.String `tfsdk:"region"`
	ProjectId      types.String `tfsdk:"project_id"`
	VpcId          types.String `tfsdk:"vpc_id"`
	NetworkRangeId types.String `tfsdk:"network_range_id"`

	Ipv4 *ipv4Model `tfsdk:"ipv4"`
}

type ipv4Model struct {
	DefaultPrefixLength types.Int64  `tfsdk:"default_prefix_length"`
	MaxPrefixLength     types.Int64  `tfsdk:"max_prefix_length"`
	MinPrefixLength     types.Int64  `tfsdk:"min_prefix_length"`
	Nameservers         types.List   `tfsdk:"nameservers"`
	Prefix              types.String `tfsdk:"prefix"`
	Description         types.String `tfsdk:"description"`
	Labels              types.Map    `tfsdk:"labels"`
}

type vpcNetworkRangeResource struct {
	// client                *iaas.APIClient
	// resourceManagerClient *resourcemanager.APIClient
}

// NewVpcNetworkRangeResource is a helper function to simplify the provider implementation.
func NewVpcNetworkRangeResource() resource.Resource {
	return &vpcNetworkRangeResource{}
}

// Metadata returns the resource type name.
func (v *vpcNetworkRangeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_network_range"
}

// Schema defines the schema for the resource.
func (v *vpcNetworkRangeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	description := "VPC Network Range resource schema."

	resp.Schema = schema.Schema{
		Description: description,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Terraform's internal resource ID. It is structured as \"`project_id`,`vpc_id`,`region`,`network_range_id`\".",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				Description: "STACKIT project ID to which the network range is associated.",
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
				Description: "The VPC ID.",
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
			"network_range_id": schema.StringAttribute{
				Description: "The network range ID",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
			"ipv4": schema.SingleNestedAttribute{
				Description: "The regional IPv4 config of a network range.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"default_prefix_length": schema.Int64Attribute{
						Description: "The default prefix length for networks in the network area.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(24),
							int64validator.AtMost(29),
						},
						Default: int64default.StaticInt64(25),
					},
					"max_prefix_length": schema.Int64Attribute{
						Description: "The maximal prefix length for networks in the network area.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(24),
							int64validator.AtMost(29),
						},
						Default: int64default.StaticInt64(29),
					},
					"min_prefix_length": schema.Int64Attribute{
						Description: "The minimal prefix length for networks in the network area.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(8),
							int64validator.AtMost(29),
						},
						Default: int64default.StaticInt64(24),
					},
					"nameservers": schema.ListAttribute{
						Description: "List of DNS Servers/Nameservers.",
						Optional:    true,
						ElementType: types.StringType,
						Validators: []validator.List{
							listvalidator.SizeAtMost(3),
						},
					},
					"prefix": schema.StringAttribute{
						Description: "Classless Inter-Domain Routing (CIDR).",
						Required:    true,
						Validators: []validator.String{
							validate.CIDR(),
						},
					},
					"description": schema.StringAttribute{
						Description: "The description of the VPC network range.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
					},
					"labels": schema.MapAttribute{
						Description: "Labels are key-value string pairs which can be attached to a resource container",
						ElementType: types.StringType,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (v *vpcNetworkRangeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var resourceModel Model
	resp.Diagnostics.Append(req.Config.Get(ctx, &resourceModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if resourceModel.Ipv4 == nil {
		core.LogAndAddError(ctx, &resp.Diagnostics, "Error configuring VPC Networkrange", "'Ipv4' must be configured.")
	}
}

// Create creates the resource and sets the initial Terraform state.
func (v *vpcNetworkRangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error create VPC region network range", "not implemented yet")
}

// Read refreshes the Terraform state with the latest data.
func (v *vpcNetworkRangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error read VPC region network range", "not implemented yet")
}

// Update updates the resource and sets the updated Terraform state on success.
func (v *vpcNetworkRangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error update VPC region network range", "not implemented yet")
}

// Delete deletes the resource and removes the Terraform state on success.
func (v *vpcNetworkRangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error delete VPC region network range", "not implemented yet")
}
