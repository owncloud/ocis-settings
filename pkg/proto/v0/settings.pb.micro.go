// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: settings.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for BundleService service

func NewBundleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "BundleService.SaveBundle",
			Path:    []string{"/api/v0/settings/bundle-save"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.GetBundle",
			Path:    []string{"/api/v0/settings/bundle-get"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.ListBundles",
			Path:    []string{"/api/v0/settings/bundles-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.AddSettingToBundle",
			Path:    []string{"/api/v0/settings/bundles-add-setting"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.RemoveSettingFromBundle",
			Path:    []string{"/api/v0/settings/bundles-remove-setting"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for BundleService service

type BundleService interface {
	SaveBundle(ctx context.Context, in *SaveBundleRequest, opts ...client.CallOption) (*SaveBundleResponse, error)
	GetBundle(ctx context.Context, in *GetBundleRequest, opts ...client.CallOption) (*GetBundleResponse, error)
	ListBundles(ctx context.Context, in *ListBundlesRequest, opts ...client.CallOption) (*ListBundlesResponse, error)
	AddSettingToBundle(ctx context.Context, in *AddSettingToBundleRequest, opts ...client.CallOption) (*AddSettingToBundleResponse, error)
	RemoveSettingFromBundle(ctx context.Context, in *RemoveSettingFromBundleRequest, opts ...client.CallOption) (*empty.Empty, error)
}

type bundleService struct {
	c    client.Client
	name string
}

func NewBundleService(name string, c client.Client) BundleService {
	return &bundleService{
		c:    c,
		name: name,
	}
}

func (c *bundleService) SaveBundle(ctx context.Context, in *SaveBundleRequest, opts ...client.CallOption) (*SaveBundleResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.SaveBundle", in)
	out := new(SaveBundleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) GetBundle(ctx context.Context, in *GetBundleRequest, opts ...client.CallOption) (*GetBundleResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.GetBundle", in)
	out := new(GetBundleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) ListBundles(ctx context.Context, in *ListBundlesRequest, opts ...client.CallOption) (*ListBundlesResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.ListBundles", in)
	out := new(ListBundlesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) AddSettingToBundle(ctx context.Context, in *AddSettingToBundleRequest, opts ...client.CallOption) (*AddSettingToBundleResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.AddSettingToBundle", in)
	out := new(AddSettingToBundleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) RemoveSettingFromBundle(ctx context.Context, in *RemoveSettingFromBundleRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "BundleService.RemoveSettingFromBundle", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BundleService service

type BundleServiceHandler interface {
	SaveBundle(context.Context, *SaveBundleRequest, *SaveBundleResponse) error
	GetBundle(context.Context, *GetBundleRequest, *GetBundleResponse) error
	ListBundles(context.Context, *ListBundlesRequest, *ListBundlesResponse) error
	AddSettingToBundle(context.Context, *AddSettingToBundleRequest, *AddSettingToBundleResponse) error
	RemoveSettingFromBundle(context.Context, *RemoveSettingFromBundleRequest, *empty.Empty) error
}

func RegisterBundleServiceHandler(s server.Server, hdlr BundleServiceHandler, opts ...server.HandlerOption) error {
	type bundleService interface {
		SaveBundle(ctx context.Context, in *SaveBundleRequest, out *SaveBundleResponse) error
		GetBundle(ctx context.Context, in *GetBundleRequest, out *GetBundleResponse) error
		ListBundles(ctx context.Context, in *ListBundlesRequest, out *ListBundlesResponse) error
		AddSettingToBundle(ctx context.Context, in *AddSettingToBundleRequest, out *AddSettingToBundleResponse) error
		RemoveSettingFromBundle(ctx context.Context, in *RemoveSettingFromBundleRequest, out *empty.Empty) error
	}
	type BundleService struct {
		bundleService
	}
	h := &bundleServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.SaveBundle",
		Path:    []string{"/api/v0/settings/bundle-save"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.GetBundle",
		Path:    []string{"/api/v0/settings/bundle-get"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.ListBundles",
		Path:    []string{"/api/v0/settings/bundles-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.AddSettingToBundle",
		Path:    []string{"/api/v0/settings/bundles-add-setting"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.RemoveSettingFromBundle",
		Path:    []string{"/api/v0/settings/bundles-remove-setting"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&BundleService{h}, opts...))
}

type bundleServiceHandler struct {
	BundleServiceHandler
}

func (h *bundleServiceHandler) SaveBundle(ctx context.Context, in *SaveBundleRequest, out *SaveBundleResponse) error {
	return h.BundleServiceHandler.SaveBundle(ctx, in, out)
}

func (h *bundleServiceHandler) GetBundle(ctx context.Context, in *GetBundleRequest, out *GetBundleResponse) error {
	return h.BundleServiceHandler.GetBundle(ctx, in, out)
}

func (h *bundleServiceHandler) ListBundles(ctx context.Context, in *ListBundlesRequest, out *ListBundlesResponse) error {
	return h.BundleServiceHandler.ListBundles(ctx, in, out)
}

func (h *bundleServiceHandler) AddSettingToBundle(ctx context.Context, in *AddSettingToBundleRequest, out *AddSettingToBundleResponse) error {
	return h.BundleServiceHandler.AddSettingToBundle(ctx, in, out)
}

func (h *bundleServiceHandler) RemoveSettingFromBundle(ctx context.Context, in *RemoveSettingFromBundleRequest, out *empty.Empty) error {
	return h.BundleServiceHandler.RemoveSettingFromBundle(ctx, in, out)
}

// Api Endpoints for ValueService service

func NewValueServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "ValueService.SaveValue",
			Path:    []string{"/api/v0/settings/values-save"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ValueService.GetValue",
			Path:    []string{"/api/v0/settings/values-get"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ValueService.ListValues",
			Path:    []string{"/api/v0/settings/values-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for ValueService service

type ValueService interface {
	SaveValue(ctx context.Context, in *SaveValueRequest, opts ...client.CallOption) (*SaveValueResponse, error)
	GetValue(ctx context.Context, in *GetValueRequest, opts ...client.CallOption) (*GetValueResponse, error)
	ListValues(ctx context.Context, in *ListValuesRequest, opts ...client.CallOption) (*ListValuesResponse, error)
}

type valueService struct {
	c    client.Client
	name string
}

func NewValueService(name string, c client.Client) ValueService {
	return &valueService{
		c:    c,
		name: name,
	}
}

func (c *valueService) SaveValue(ctx context.Context, in *SaveValueRequest, opts ...client.CallOption) (*SaveValueResponse, error) {
	req := c.c.NewRequest(c.name, "ValueService.SaveValue", in)
	out := new(SaveValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueService) GetValue(ctx context.Context, in *GetValueRequest, opts ...client.CallOption) (*GetValueResponse, error) {
	req := c.c.NewRequest(c.name, "ValueService.GetValue", in)
	out := new(GetValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueService) ListValues(ctx context.Context, in *ListValuesRequest, opts ...client.CallOption) (*ListValuesResponse, error) {
	req := c.c.NewRequest(c.name, "ValueService.ListValues", in)
	out := new(ListValuesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ValueService service

type ValueServiceHandler interface {
	SaveValue(context.Context, *SaveValueRequest, *SaveValueResponse) error
	GetValue(context.Context, *GetValueRequest, *GetValueResponse) error
	ListValues(context.Context, *ListValuesRequest, *ListValuesResponse) error
}

func RegisterValueServiceHandler(s server.Server, hdlr ValueServiceHandler, opts ...server.HandlerOption) error {
	type valueService interface {
		SaveValue(ctx context.Context, in *SaveValueRequest, out *SaveValueResponse) error
		GetValue(ctx context.Context, in *GetValueRequest, out *GetValueResponse) error
		ListValues(ctx context.Context, in *ListValuesRequest, out *ListValuesResponse) error
	}
	type ValueService struct {
		valueService
	}
	h := &valueServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ValueService.SaveValue",
		Path:    []string{"/api/v0/settings/values-save"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ValueService.GetValue",
		Path:    []string{"/api/v0/settings/values-get"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ValueService.ListValues",
		Path:    []string{"/api/v0/settings/values-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&ValueService{h}, opts...))
}

type valueServiceHandler struct {
	ValueServiceHandler
}

func (h *valueServiceHandler) SaveValue(ctx context.Context, in *SaveValueRequest, out *SaveValueResponse) error {
	return h.ValueServiceHandler.SaveValue(ctx, in, out)
}

func (h *valueServiceHandler) GetValue(ctx context.Context, in *GetValueRequest, out *GetValueResponse) error {
	return h.ValueServiceHandler.GetValue(ctx, in, out)
}

func (h *valueServiceHandler) ListValues(ctx context.Context, in *ListValuesRequest, out *ListValuesResponse) error {
	return h.ValueServiceHandler.ListValues(ctx, in, out)
}

// Api Endpoints for RoleService service

func NewRoleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "RoleService.ListRoles",
			Path:    []string{"/api/v0/settings/roles-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RoleService.ListRoleAssignments",
			Path:    []string{"/api/v0/settings/assignments-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RoleService.AssignRoleToUser",
			Path:    []string{"/api/v0/settings/assignments-add"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RoleService.RemoveRoleFromUser",
			Path:    []string{"/api/v0/settings/assignments-remove"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for RoleService service

type RoleService interface {
	ListRoles(ctx context.Context, in *ListBundlesRequest, opts ...client.CallOption) (*ListBundlesResponse, error)
	ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, opts ...client.CallOption) (*ListRoleAssignmentsResponse, error)
	AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, opts ...client.CallOption) (*AssignRoleToUserResponse, error)
	RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, opts ...client.CallOption) (*empty.Empty, error)
}

type roleService struct {
	c    client.Client
	name string
}

func NewRoleService(name string, c client.Client) RoleService {
	return &roleService{
		c:    c,
		name: name,
	}
}

func (c *roleService) ListRoles(ctx context.Context, in *ListBundlesRequest, opts ...client.CallOption) (*ListBundlesResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.ListRoles", in)
	out := new(ListBundlesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, opts ...client.CallOption) (*ListRoleAssignmentsResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.ListRoleAssignments", in)
	out := new(ListRoleAssignmentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, opts ...client.CallOption) (*AssignRoleToUserResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.AssignRoleToUser", in)
	out := new(AssignRoleToUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "RoleService.RemoveRoleFromUser", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoleService service

type RoleServiceHandler interface {
	ListRoles(context.Context, *ListBundlesRequest, *ListBundlesResponse) error
	ListRoleAssignments(context.Context, *ListRoleAssignmentsRequest, *ListRoleAssignmentsResponse) error
	AssignRoleToUser(context.Context, *AssignRoleToUserRequest, *AssignRoleToUserResponse) error
	RemoveRoleFromUser(context.Context, *RemoveRoleFromUserRequest, *empty.Empty) error
}

func RegisterRoleServiceHandler(s server.Server, hdlr RoleServiceHandler, opts ...server.HandlerOption) error {
	type roleService interface {
		ListRoles(ctx context.Context, in *ListBundlesRequest, out *ListBundlesResponse) error
		ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, out *ListRoleAssignmentsResponse) error
		AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, out *AssignRoleToUserResponse) error
		RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, out *empty.Empty) error
	}
	type RoleService struct {
		roleService
	}
	h := &roleServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.ListRoles",
		Path:    []string{"/api/v0/settings/roles-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.ListRoleAssignments",
		Path:    []string{"/api/v0/settings/assignments-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.AssignRoleToUser",
		Path:    []string{"/api/v0/settings/assignments-add"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.RemoveRoleFromUser",
		Path:    []string{"/api/v0/settings/assignments-remove"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&RoleService{h}, opts...))
}

type roleServiceHandler struct {
	RoleServiceHandler
}

func (h *roleServiceHandler) ListRoles(ctx context.Context, in *ListBundlesRequest, out *ListBundlesResponse) error {
	return h.RoleServiceHandler.ListRoles(ctx, in, out)
}

func (h *roleServiceHandler) ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, out *ListRoleAssignmentsResponse) error {
	return h.RoleServiceHandler.ListRoleAssignments(ctx, in, out)
}

func (h *roleServiceHandler) AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, out *AssignRoleToUserResponse) error {
	return h.RoleServiceHandler.AssignRoleToUser(ctx, in, out)
}

func (h *roleServiceHandler) RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, out *empty.Empty) error {
	return h.RoleServiceHandler.RemoveRoleFromUser(ctx, in, out)
}
