package vpcregion

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	_ resource.Resource = &vpcRegionResource{}
)

type Model struct {
	Id        types.String `tfsdk:"id"` // needed by TF
	Region    types.String `tfsdk:"region"`
	ProjectId types.String `tfsdk:"project_id"`
	VpcId     types.String `tfsdk:"vpc_id"`

	DefaultRoutingTable types.String `tfsdk:"default_routing_table"`
	Ipv4                *ipv4Model   `tfsdk:"ipv4"`
}

type ipv4Model struct {
	DefaultNameservers  types.List   `tfsdk:"default_nameservers"`
	DefaultNetworkRange types.String `tfsdk:"default_network_range"`
}

// NewVpcResource is a helper function to simplify the provider implementation.
func NewVpcRegionResource() resource.Resource {
	return &vpcRegionResource{}
}

// vpcRegionResource is the resource implementation.
type vpcRegionResource struct {
	// client                *iaas.APIClient
	// resourceManagerClient *resourcemanager.APIClient
}

// Metadata returns the resource type name.
func (r *vpcRegionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpc_region"
}

func (r *vpcRegionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "VPC region configuration resource schema.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Terraform's internal resource identifier. It is structured as \"`project_id`,`vpc_id`,`region`\".",
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
			"ipv4": schema.SingleNestedAttribute{
				Description: "The regional IPv4 config of a VPC.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"default_nameservers": schema.ListAttribute{
						Description: "List of IPv4 DNS Servers/Nameservers.",
						ElementType: types.StringType,
						Required:    true,
						Validators: []validator.List{
							listvalidator.SizeAtMost(3),
							listvalidator.ValueStringsAre(validate.IP(false)),
						},
					},
					"default_network_range": schema.StringAttribute{
						Description: "The ID of the default IPv4 network range.",
						Optional:    true,
						Validators: []validator.String{
							validate.UUID(),
							validate.NoSeparator(),
						},
					},
				},
			},
			"default_routing_table": schema.StringAttribute{
				Description: "The ID of the default routing table.",
				Optional:    true,
				Validators: []validator.String{
					validate.UUID(),
					validate.NoSeparator(),
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *vpcRegionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error create VPC region configuration", "not implemented yet")
}

// Read refreshes the Terraform state with the latest data.
func (r *vpcRegionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error read VPC region configuration", "not implemented yet")
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *vpcRegionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error update VPC region configuration", "not implemented yet")
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *vpcRegionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) { // nolint:all // function signature required by Terraform
	core.LogAndAddError(ctx, &resp.Diagnostics, "Error delete VPC region configuration", "not implemented yet")
}
