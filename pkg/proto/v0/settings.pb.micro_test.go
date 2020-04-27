package proto_test

import (
	"context"
	"testing"
	"encoding/json"

	svc "github.com/owncloud/ocis-settings/pkg/service/v0"
	"github.com/owncloud/ocis-pkg/v2/service/grpc"
	"log"
	"github.com/owncloud/ocis-settings/pkg/proto/v0"
	"github.com/owncloud/ocis-settings/pkg/config"
	"github.com/stretchr/testify/assert"
)

var service = grpc.Service{}

func init() {
	service = grpc.NewService(
		grpc.Namespace("com.owncloud.api"),
		grpc.Name("settings"),
		grpc.Address("localhost:9992"),
	)

	cfg := config.New()
	err := proto.RegisterBundleServiceHandler(service.Server(), svc.NewService(cfg))
	if err != nil {
		log.Fatalf("could not register HelloHandler: %v", err)
	}
	service.Server().Start()
}

type CustomError struct {
	Id     string
	Code   int
	Detail string
	Status string
}

type TestStruct struct {
	testDataName  string
	Key           string
	DisplayName   string
	Extension     string
	expectedError CustomError
}

func TestCreateSettingsBundle(t *testing.T) {
	var tests = []TestStruct{
		{
			"ASCII",
			"simple-key",
			"simple-display-name",
			"simple-extension-name",
			CustomError{},
		},
		{
			"space in values",
			"simple key",
			"simple display name",
			"simple extension name",
			CustomError{},
		},
		{
			"UNICODE",
			"सिम्पलेछबि",
			"सिम्पले-name",
			"सिम्पले-extension-name",
			CustomError{},
		},
		{
			"key missing",
			"",
			"simple-display-name",
			"simple-extension-name",
			CustomError{
				Id:     "go.micro.client",
				Code:   500,
				Detail: "key cannot be empty",
				Status: "Internal Server Error",
			},
		},
		{
			"display name missing",
			"simple-key",
			"",
			"simple-extension-name",
			CustomError{},
		},
		{
			"extension name missing",
			"simple-key",
			"simple-display-name",
			"",
			CustomError{
				Id:     "go.micro.client",
				Code:   500,
				Detail: "key cannot be empty",
				Status: "Internal Server Error",
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.testDataName, func(t *testing.T) {
			bundle := proto.SettingsBundle{
				Key:         testCase.Key,
				DisplayName: testCase.DisplayName,
				Extension:   testCase.Extension,
				Settings:    nil,
			}
			createRequest := proto.CreateSettingsBundleRequest{
				SettingsBundle: &bundle,
			}

			client := service.Client()
			cl := proto.NewBundleService("com.owncloud.api.settings", client)

			cresponse, err := cl.CreateSettingsBundle(context.Background(), &createRequest)
			if err != nil || (CustomError{} != testCase.expectedError) {
				var errorData CustomError
				json.Unmarshal([]byte(err.Error()), &errorData)
				assert.Equal(t, testCase.expectedError.Id, errorData.Id)
				assert.Equal(t, testCase.expectedError.Code, errorData.Code)
				assert.Equal(t, testCase.expectedError.Detail, errorData.Detail)
				assert.Equal(t, testCase.expectedError.Status, errorData.Status)
			} else {
				assert.Equal(t, testCase.Key, cresponse.SettingsBundle.Key)
				assert.Equal(t, testCase.DisplayName, cresponse.SettingsBundle.DisplayName)
				assert.Equal(t, testCase.Extension, cresponse.SettingsBundle.Extension)
			}
		})
	}
}

func TestListSettingsBundlesWithNoSettings(t *testing.T) {
	var tests = []TestStruct{
		{
			"ASCII",
			"simple-key",
			"simple-display-name",
			"simple-extension-name",
			CustomError{},
		},
		{
			"space in values",
			"simple key",
			"simple display name",
			"simple extension name",
			CustomError{},
		},
		{
			"UNICODE",
			"सिम्पलेछबि",
			"सिम्पले-name",
			"सिम्पले-extension-name",
			CustomError{},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.testDataName, func(t *testing.T) {
			bundle := proto.SettingsBundle{
				Key:         testCase.Key,
				DisplayName: testCase.DisplayName,
				Extension:   testCase.Extension,
				Settings:    nil,
			}
			createRequest := proto.CreateSettingsBundleRequest{
				SettingsBundle: &bundle,
			}

			client := service.Client()
			cl := proto.NewBundleService("com.owncloud.api.settings", client)

			_, err := cl.CreateSettingsBundle(context.Background(), &createRequest)
			assert.NoError(t, err)
			request := proto.ListSettingsBundlesRequest{Extension: testCase.Extension}
			response, err := cl.ListSettingsBundles(context.Background(), &request)
			assert.NoError(t, err)
			assert.Equal(t, testCase.Key, response.SettingsBundles[0].Key)
			assert.Equal(t, "", response.SettingsBundles[0].DisplayName)
			assert.Equal(t, testCase.Extension, response.SettingsBundles[0].Extension)
		})
	}
}

func TestListSettingsBundlesWithSettings(t *testing.T) {
	var tests = []TestStruct{
		{
			"ASCII",
			"simple-key",
			"simple-display-name",
			"simple-extension-name",
			CustomError{},
		},
		{
			"space in values",
			"simple key",
			"simple display name",
			"simple extension name",
			CustomError{},
		},
		{
			"UNICODE",
			"सिम्पलेछबि",
			"सिम्पले-name",
			"सिम्पले-extension-name",
			CustomError{},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.testDataName, func(t *testing.T) {
			var settings []*proto.Setting

			settings = append(settings, &proto.Setting{
				Key:         "setting.Key1",
				DisplayName: "setting.DisplayName1",
				Description: "setting.Description1",
			})
			settings = append(settings, &proto.Setting{
				Key:         "setting.Key2",
				DisplayName: "setting.DisplayName2",
				Description: "setting.Description2",
			})
			bundle := proto.SettingsBundle{
				Key:         testCase.Key,
				Extension:   testCase.Extension,
				Settings:    settings,
			}
			createRequest := proto.CreateSettingsBundleRequest{
				SettingsBundle: &bundle,
			}

			client := service.Client()
			cl := proto.NewBundleService("com.owncloud.api.settings", client)

			_, err := cl.CreateSettingsBundle(context.Background(), &createRequest)
			assert.NoError(t, err)
			request := proto.ListSettingsBundlesRequest{Extension: testCase.Extension}
			response, err := cl.ListSettingsBundles(context.Background(), &request)
			assert.NoError(t, err)
			assert.Equal(t, testCase.Key, response.SettingsBundles[0].Key)
			assert.Equal(t, "", response.SettingsBundles[0].DisplayName)
			assert.Equal(t, testCase.Extension, response.SettingsBundles[0].Extension)
		})
	}
}

//func xTestRubbish(t *testing.T) {
//	var settings []*proto.Setting
//
//
//	settings = append(settings, &proto.Setting	{
//		Key:                  "settingKey",
//		DisplayName:          "settingDisplay",
//		Description:          "settingDesc",
//		Values:               nil,
//	})
//	bundle:=proto.SettingsBundle{
//		Key:                  "testkey",
//		DisplayName:          "test-name",
//		Extension:            "test-extension",
//		Settings:             settings,
//	}
//	createRequest := proto.CreateSettingsBundleRequest{
//		SettingsBundle:       &bundle,
//	}
//	request := proto.ListSettingsBundlesRequest{Extension: "test-extension"}
//	client := service.Client()
//	cl := proto.NewBundleService("com.owncloud.api.settings", client)
//
//	cresponse, _ := cl.CreateSettingsBundle(context.Background(), &createRequest)
//	log.Printf(cresponse.String())
//	response, _ := cl.ListSettingsBundles(context.Background(), &request)
//	log.Printf(response.String())
//
//	assert.Equal(t, "", response.String())
//}
