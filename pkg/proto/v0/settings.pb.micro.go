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
			Name:    "BundleService.SaveSettingsBundle",
			Path:    []string{"/api/v0/settings/bundle-save"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.GetSettingsBundle",
			Path:    []string{"/api/v0/settings/bundle-get"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.ListSettingsBundles",
			Path:    []string{"/api/v0/settings/bundles-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.AddSettingToSettingsBundle",
			Path:    []string{"/api/v0/settings/bundles-add-setting"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "BundleService.RemoveSettingFromSettingsBundle",
			Path:    []string{"/api/v0/settings/bundles-remove-setting"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for BundleService service

type BundleService interface {
	SaveSettingsBundle(ctx context.Context, in *SaveSettingsBundleRequest, opts ...client.CallOption) (*SaveSettingsBundleResponse, error)
	GetSettingsBundle(ctx context.Context, in *GetSettingsBundleRequest, opts ...client.CallOption) (*GetSettingsBundleResponse, error)
	ListSettingsBundles(ctx context.Context, in *ListSettingsBundlesRequest, opts ...client.CallOption) (*ListSettingsBundlesResponse, error)
	AddSettingToSettingsBundle(ctx context.Context, in *AddSettingToSettingsBundleRequest, opts ...client.CallOption) (*empty.Empty, error)
	RemoveSettingFromSettingsBundle(ctx context.Context, in *RemoveSettingFromSettingsBundleRequest, opts ...client.CallOption) (*empty.Empty, error)
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

func (c *bundleService) SaveSettingsBundle(ctx context.Context, in *SaveSettingsBundleRequest, opts ...client.CallOption) (*SaveSettingsBundleResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.SaveSettingsBundle", in)
	out := new(SaveSettingsBundleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) GetSettingsBundle(ctx context.Context, in *GetSettingsBundleRequest, opts ...client.CallOption) (*GetSettingsBundleResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.GetSettingsBundle", in)
	out := new(GetSettingsBundleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) ListSettingsBundles(ctx context.Context, in *ListSettingsBundlesRequest, opts ...client.CallOption) (*ListSettingsBundlesResponse, error) {
	req := c.c.NewRequest(c.name, "BundleService.ListSettingsBundles", in)
	out := new(ListSettingsBundlesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) AddSettingToSettingsBundle(ctx context.Context, in *AddSettingToSettingsBundleRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "BundleService.AddSettingToSettingsBundle", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundleService) RemoveSettingFromSettingsBundle(ctx context.Context, in *RemoveSettingFromSettingsBundleRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "BundleService.RemoveSettingFromSettingsBundle", in)
	out := new(empty.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BundleService service

type BundleServiceHandler interface {
	SaveSettingsBundle(context.Context, *SaveSettingsBundleRequest, *SaveSettingsBundleResponse) error
	GetSettingsBundle(context.Context, *GetSettingsBundleRequest, *GetSettingsBundleResponse) error
	ListSettingsBundles(context.Context, *ListSettingsBundlesRequest, *ListSettingsBundlesResponse) error
	AddSettingToSettingsBundle(context.Context, *AddSettingToSettingsBundleRequest, *empty.Empty) error
	RemoveSettingFromSettingsBundle(context.Context, *RemoveSettingFromSettingsBundleRequest, *empty.Empty) error
}

func RegisterBundleServiceHandler(s server.Server, hdlr BundleServiceHandler, opts ...server.HandlerOption) error {
	type bundleService interface {
		SaveSettingsBundle(ctx context.Context, in *SaveSettingsBundleRequest, out *SaveSettingsBundleResponse) error
		GetSettingsBundle(ctx context.Context, in *GetSettingsBundleRequest, out *GetSettingsBundleResponse) error
		ListSettingsBundles(ctx context.Context, in *ListSettingsBundlesRequest, out *ListSettingsBundlesResponse) error
		AddSettingToSettingsBundle(ctx context.Context, in *AddSettingToSettingsBundleRequest, out *empty.Empty) error
		RemoveSettingFromSettingsBundle(ctx context.Context, in *RemoveSettingFromSettingsBundleRequest, out *empty.Empty) error
	}
	type BundleService struct {
		bundleService
	}
	h := &bundleServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.SaveSettingsBundle",
		Path:    []string{"/api/v0/settings/bundle-save"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.GetSettingsBundle",
		Path:    []string{"/api/v0/settings/bundle-get"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.ListSettingsBundles",
		Path:    []string{"/api/v0/settings/bundles-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.AddSettingToSettingsBundle",
		Path:    []string{"/api/v0/settings/bundles-add-setting"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "BundleService.RemoveSettingFromSettingsBundle",
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

func (h *bundleServiceHandler) SaveSettingsBundle(ctx context.Context, in *SaveSettingsBundleRequest, out *SaveSettingsBundleResponse) error {
	return h.BundleServiceHandler.SaveSettingsBundle(ctx, in, out)
}

func (h *bundleServiceHandler) GetSettingsBundle(ctx context.Context, in *GetSettingsBundleRequest, out *GetSettingsBundleResponse) error {
	return h.BundleServiceHandler.GetSettingsBundle(ctx, in, out)
}

func (h *bundleServiceHandler) ListSettingsBundles(ctx context.Context, in *ListSettingsBundlesRequest, out *ListSettingsBundlesResponse) error {
	return h.BundleServiceHandler.ListSettingsBundles(ctx, in, out)
}

func (h *bundleServiceHandler) AddSettingToSettingsBundle(ctx context.Context, in *AddSettingToSettingsBundleRequest, out *empty.Empty) error {
	return h.BundleServiceHandler.AddSettingToSettingsBundle(ctx, in, out)
}

func (h *bundleServiceHandler) RemoveSettingFromSettingsBundle(ctx context.Context, in *RemoveSettingFromSettingsBundleRequest, out *empty.Empty) error {
	return h.BundleServiceHandler.RemoveSettingFromSettingsBundle(ctx, in, out)
}

// Api Endpoints for ValueService service

func NewValueServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "ValueService.SaveSettingsValue",
			Path:    []string{"/api/v0/settings/value-save"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ValueService.GetSettingsValue",
			Path:    []string{"/api/v0/settings/value-get"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "ValueService.ListSettingsValues",
			Path:    []string{"/api/v0/settings/values-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for ValueService service

type ValueService interface {
	SaveSettingsValue(ctx context.Context, in *SaveSettingsValueRequest, opts ...client.CallOption) (*SaveSettingsValueResponse, error)
	GetSettingsValue(ctx context.Context, in *GetSettingsValueRequest, opts ...client.CallOption) (*GetSettingsValueResponse, error)
	ListSettingsValues(ctx context.Context, in *ListSettingsValuesRequest, opts ...client.CallOption) (*ListSettingsValuesResponse, error)
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

func (c *valueService) SaveSettingsValue(ctx context.Context, in *SaveSettingsValueRequest, opts ...client.CallOption) (*SaveSettingsValueResponse, error) {
	req := c.c.NewRequest(c.name, "ValueService.SaveSettingsValue", in)
	out := new(SaveSettingsValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueService) GetSettingsValue(ctx context.Context, in *GetSettingsValueRequest, opts ...client.CallOption) (*GetSettingsValueResponse, error) {
	req := c.c.NewRequest(c.name, "ValueService.GetSettingsValue", in)
	out := new(GetSettingsValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueService) ListSettingsValues(ctx context.Context, in *ListSettingsValuesRequest, opts ...client.CallOption) (*ListSettingsValuesResponse, error) {
	req := c.c.NewRequest(c.name, "ValueService.ListSettingsValues", in)
	out := new(ListSettingsValuesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ValueService service

type ValueServiceHandler interface {
	SaveSettingsValue(context.Context, *SaveSettingsValueRequest, *SaveSettingsValueResponse) error
	GetSettingsValue(context.Context, *GetSettingsValueRequest, *GetSettingsValueResponse) error
	ListSettingsValues(context.Context, *ListSettingsValuesRequest, *ListSettingsValuesResponse) error
}

func RegisterValueServiceHandler(s server.Server, hdlr ValueServiceHandler, opts ...server.HandlerOption) error {
	type valueService interface {
		SaveSettingsValue(ctx context.Context, in *SaveSettingsValueRequest, out *SaveSettingsValueResponse) error
		GetSettingsValue(ctx context.Context, in *GetSettingsValueRequest, out *GetSettingsValueResponse) error
		ListSettingsValues(ctx context.Context, in *ListSettingsValuesRequest, out *ListSettingsValuesResponse) error
	}
	type ValueService struct {
		valueService
	}
	h := &valueServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ValueService.SaveSettingsValue",
		Path:    []string{"/api/v0/settings/value-save"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ValueService.GetSettingsValue",
		Path:    []string{"/api/v0/settings/value-get"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "ValueService.ListSettingsValues",
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

func (h *valueServiceHandler) SaveSettingsValue(ctx context.Context, in *SaveSettingsValueRequest, out *SaveSettingsValueResponse) error {
	return h.ValueServiceHandler.SaveSettingsValue(ctx, in, out)
}

func (h *valueServiceHandler) GetSettingsValue(ctx context.Context, in *GetSettingsValueRequest, out *GetSettingsValueResponse) error {
	return h.ValueServiceHandler.GetSettingsValue(ctx, in, out)
}

func (h *valueServiceHandler) ListSettingsValues(ctx context.Context, in *ListSettingsValuesRequest, out *ListSettingsValuesResponse) error {
	return h.ValueServiceHandler.ListSettingsValues(ctx, in, out)
}

// Api Endpoints for RoleService service

func NewRoleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "RoleService.ListRoleAssignments",
			Path:    []string{"/api/v0/settings/roles-list"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RoleService.AssignRoleToUser",
			Path:    []string{"/api/v0/settings/roles-assign"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "RoleService.RemoveRoleFromUser",
			Path:    []string{"/api/v0/settings/roles-remove"},
			Method:  []string{"POST"},
			Body:    "*",
			Handler: "rpc",
		},
	}
}

// Client API for RoleService service

type RoleService interface {
	ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, opts ...client.CallOption) (*UserRoleAssignments, error)
	AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, opts ...client.CallOption) (*empty.Empty, error)
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

func (c *roleService) ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, opts ...client.CallOption) (*UserRoleAssignments, error) {
	req := c.c.NewRequest(c.name, "RoleService.ListRoleAssignments", in)
	out := new(UserRoleAssignments)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, opts ...client.CallOption) (*empty.Empty, error) {
	req := c.c.NewRequest(c.name, "RoleService.AssignRoleToUser", in)
	out := new(empty.Empty)
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
	ListRoleAssignments(context.Context, *ListRoleAssignmentsRequest, *UserRoleAssignments) error
	AssignRoleToUser(context.Context, *AssignRoleToUserRequest, *empty.Empty) error
	RemoveRoleFromUser(context.Context, *RemoveRoleFromUserRequest, *empty.Empty) error
}

func RegisterRoleServiceHandler(s server.Server, hdlr RoleServiceHandler, opts ...server.HandlerOption) error {
	type roleService interface {
		ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, out *UserRoleAssignments) error
		AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, out *empty.Empty) error
		RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, out *empty.Empty) error
	}
	type RoleService struct {
		roleService
	}
	h := &roleServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.ListRoleAssignments",
		Path:    []string{"/api/v0/settings/roles-list"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.AssignRoleToUser",
		Path:    []string{"/api/v0/settings/roles-assign"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "RoleService.RemoveRoleFromUser",
		Path:    []string{"/api/v0/settings/roles-remove"},
		Method:  []string{"POST"},
		Body:    "*",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&RoleService{h}, opts...))
}

type roleServiceHandler struct {
	RoleServiceHandler
}

func (h *roleServiceHandler) ListRoleAssignments(ctx context.Context, in *ListRoleAssignmentsRequest, out *UserRoleAssignments) error {
	return h.RoleServiceHandler.ListRoleAssignments(ctx, in, out)
}

func (h *roleServiceHandler) AssignRoleToUser(ctx context.Context, in *AssignRoleToUserRequest, out *empty.Empty) error {
	return h.RoleServiceHandler.AssignRoleToUser(ctx, in, out)
}

func (h *roleServiceHandler) RemoveRoleFromUser(ctx context.Context, in *RemoveRoleFromUserRequest, out *empty.Empty) error {
	return h.RoleServiceHandler.RemoveRoleFromUser(ctx, in, out)
}
