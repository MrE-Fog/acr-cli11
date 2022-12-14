// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	acr "github.com/Azure/acr-cli/acr"

	autorest "github.com/Azure/go-autorest/autorest"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BaseClientAPI is an autogenerated mock type for the BaseClientAPI type
type BaseClientAPI struct {
	mock.Mock
}

// CancelBlobUpload provides a mock function with given fields: ctx, name, UUID
func (_m *BaseClientAPI) CancelBlobUpload(ctx context.Context, name string, UUID string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, UUID)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, UUID)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckBlobExistence provides a mock function with given fields: ctx, name, digest
func (_m *BaseClientAPI) CheckBlobExistence(ctx context.Context, name string, digest string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, digest)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, digest)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateManifest provides a mock function with given fields: ctx, name, reference, payload
func (_m *BaseClientAPI) CreateManifest(ctx context.Context, name string, reference string, payload acr.Manifest) (acr.SetObject, error) {
	ret := _m.Called(ctx, name, reference, payload)

	var r0 acr.SetObject
	if rf, ok := ret.Get(0).(func(context.Context, string, string, acr.Manifest) acr.SetObject); ok {
		r0 = rf(ctx, name, reference, payload)
	} else {
		r0 = ret.Get(0).(acr.SetObject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, acr.Manifest) error); ok {
		r1 = rf(ctx, name, reference, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcrManifestMetadata provides a mock function with given fields: ctx, name, reference, metadata
func (_m *BaseClientAPI) DeleteAcrManifestMetadata(ctx context.Context, name string, reference string, metadata string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference, metadata)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, reference, metadata)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, reference, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcrRepository provides a mock function with given fields: ctx, name
func (_m *BaseClientAPI) DeleteAcrRepository(ctx context.Context, name string) (acr.DeletedRepository, error) {
	ret := _m.Called(ctx, name)

	var r0 acr.DeletedRepository
	if rf, ok := ret.Get(0).(func(context.Context, string) acr.DeletedRepository); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(acr.DeletedRepository)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcrRepositoryMetadata provides a mock function with given fields: ctx, name, metadata
func (_m *BaseClientAPI) DeleteAcrRepositoryMetadata(ctx context.Context, name string, metadata string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, metadata)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, metadata)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcrTag provides a mock function with given fields: ctx, name, reference
func (_m *BaseClientAPI) DeleteAcrTag(ctx context.Context, name string, reference string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, reference)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAcrTagMetadata provides a mock function with given fields: ctx, name, reference, metadata
func (_m *BaseClientAPI) DeleteAcrTagMetadata(ctx context.Context, name string, reference string, metadata string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference, metadata)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, reference, metadata)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, reference, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteManifest provides a mock function with given fields: ctx, name, reference
func (_m *BaseClientAPI) DeleteManifest(ctx context.Context, name string, reference string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, reference)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EndBlobUpload provides a mock function with given fields: ctx, digest, name, UUID
func (_m *BaseClientAPI) EndBlobUpload(ctx context.Context, digest string, name string, UUID string) (autorest.Response, error) {
	ret := _m.Called(ctx, digest, name, UUID)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) autorest.Response); ok {
		r0 = rf(ctx, digest, name, UUID)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, digest, name, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrAccessToken provides a mock function with given fields: ctx, service, scope, refreshToken
func (_m *BaseClientAPI) GetAcrAccessToken(ctx context.Context, service string, scope string, refreshToken string) (acr.AccessToken, error) {
	ret := _m.Called(ctx, service, scope, refreshToken)

	var r0 acr.AccessToken
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) acr.AccessToken); ok {
		r0 = rf(ctx, service, scope, refreshToken)
	} else {
		r0 = ret.Get(0).(acr.AccessToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, service, scope, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrAccessTokenFromLogin provides a mock function with given fields: ctx, service, scope
func (_m *BaseClientAPI) GetAcrAccessTokenFromLogin(ctx context.Context, service string, scope string) (acr.AccessToken, error) {
	ret := _m.Called(ctx, service, scope)

	var r0 acr.AccessToken
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.AccessToken); ok {
		r0 = rf(ctx, service, scope)
	} else {
		r0 = ret.Get(0).(acr.AccessToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, service, scope)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrManifestAttributes provides a mock function with given fields: ctx, name, reference
func (_m *BaseClientAPI) GetAcrManifestAttributes(ctx context.Context, name string, reference string) (acr.ManifestAttributes, error) {
	ret := _m.Called(ctx, name, reference)

	var r0 acr.ManifestAttributes
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.ManifestAttributes); ok {
		r0 = rf(ctx, name, reference)
	} else {
		r0 = ret.Get(0).(acr.ManifestAttributes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrManifestMetadata provides a mock function with given fields: ctx, name, reference, metadata
func (_m *BaseClientAPI) GetAcrManifestMetadata(ctx context.Context, name string, reference string, metadata string) (acr.SetObject, error) {
	ret := _m.Called(ctx, name, reference, metadata)

	var r0 acr.SetObject
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) acr.SetObject); ok {
		r0 = rf(ctx, name, reference, metadata)
	} else {
		r0 = ret.Get(0).(acr.SetObject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, reference, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrManifests provides a mock function with given fields: ctx, name, last, n, orderby
func (_m *BaseClientAPI) GetAcrManifests(ctx context.Context, name string, last string, n *int32, orderby string) (acr.Manifests, error) {
	ret := _m.Called(ctx, name, last, n, orderby)

	var r0 acr.Manifests
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int32, string) acr.Manifests); ok {
		r0 = rf(ctx, name, last, n, orderby)
	} else {
		r0 = ret.Get(0).(acr.Manifests)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *int32, string) error); ok {
		r1 = rf(ctx, name, last, n, orderby)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrRefreshTokenFromExchange provides a mock function with given fields: ctx, grantType, service, tenant, refreshToken, accessToken
func (_m *BaseClientAPI) GetAcrRefreshTokenFromExchange(ctx context.Context, grantType string, service string, tenant string, refreshToken string, accessToken string) (acr.RefreshToken, error) {
	ret := _m.Called(ctx, grantType, service, tenant, refreshToken, accessToken)

	var r0 acr.RefreshToken
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) acr.RefreshToken); ok {
		r0 = rf(ctx, grantType, service, tenant, refreshToken, accessToken)
	} else {
		r0 = ret.Get(0).(acr.RefreshToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string) error); ok {
		r1 = rf(ctx, grantType, service, tenant, refreshToken, accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrRepositories provides a mock function with given fields: ctx, last, n
func (_m *BaseClientAPI) GetAcrRepositories(ctx context.Context, last string, n *int32) (acr.Repositories, error) {
	ret := _m.Called(ctx, last, n)

	var r0 acr.Repositories
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32) acr.Repositories); ok {
		r0 = rf(ctx, last, n)
	} else {
		r0 = ret.Get(0).(acr.Repositories)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *int32) error); ok {
		r1 = rf(ctx, last, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrRepositoryAttributes provides a mock function with given fields: ctx, name
func (_m *BaseClientAPI) GetAcrRepositoryAttributes(ctx context.Context, name string) (acr.RepositoryAttributes, error) {
	ret := _m.Called(ctx, name)

	var r0 acr.RepositoryAttributes
	if rf, ok := ret.Get(0).(func(context.Context, string) acr.RepositoryAttributes); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(acr.RepositoryAttributes)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrRepositoryMetadata provides a mock function with given fields: ctx, name, metadata
func (_m *BaseClientAPI) GetAcrRepositoryMetadata(ctx context.Context, name string, metadata string) (acr.SetObject, error) {
	ret := _m.Called(ctx, name, metadata)

	var r0 acr.SetObject
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.SetObject); ok {
		r0 = rf(ctx, name, metadata)
	} else {
		r0 = ret.Get(0).(acr.SetObject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrTagAttributes provides a mock function with given fields: ctx, name, reference
func (_m *BaseClientAPI) GetAcrTagAttributes(ctx context.Context, name string, reference string) (acr.TagAttributesType, error) {
	ret := _m.Called(ctx, name, reference)

	var r0 acr.TagAttributesType
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.TagAttributesType); ok {
		r0 = rf(ctx, name, reference)
	} else {
		r0 = ret.Get(0).(acr.TagAttributesType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrTagMetadata provides a mock function with given fields: ctx, name, reference, metadata
func (_m *BaseClientAPI) GetAcrTagMetadata(ctx context.Context, name string, reference string, metadata string) (acr.SetObject, error) {
	ret := _m.Called(ctx, name, reference, metadata)

	var r0 acr.SetObject
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) acr.SetObject); ok {
		r0 = rf(ctx, name, reference, metadata)
	} else {
		r0 = ret.Get(0).(acr.SetObject)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, reference, metadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcrTags provides a mock function with given fields: ctx, name, last, n, orderby, digest
func (_m *BaseClientAPI) GetAcrTags(ctx context.Context, name string, last string, n *int32, orderby string, digest string) (acr.RepositoryTagsType, error) {
	ret := _m.Called(ctx, name, last, n, orderby, digest)

	var r0 acr.RepositoryTagsType
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *int32, string, string) acr.RepositoryTagsType); ok {
		r0 = rf(ctx, name, last, n, orderby, digest)
	} else {
		r0 = ret.Get(0).(acr.RepositoryTagsType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *int32, string, string) error); ok {
		r1 = rf(ctx, name, last, n, orderby, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlob provides a mock function with given fields: ctx, name, digest
func (_m *BaseClientAPI) GetBlob(ctx context.Context, name string, digest string) (acr.String, error) {
	ret := _m.Called(ctx, name, digest)

	var r0 acr.String
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.String); ok {
		r0 = rf(ctx, name, digest)
	} else {
		r0 = ret.Get(0).(acr.String)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlobUploadStatus provides a mock function with given fields: ctx, name, UUID
func (_m *BaseClientAPI) GetBlobUploadStatus(ctx context.Context, name string, UUID string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, UUID)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, UUID)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDockerRegistryV2Support provides a mock function with given fields: ctx
func (_m *BaseClientAPI) GetDockerRegistryV2Support(ctx context.Context) (autorest.Response, error) {
	ret := _m.Called(ctx)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context) autorest.Response); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManifest provides a mock function with given fields: ctx, name, reference, accept
func (_m *BaseClientAPI) GetManifest(ctx context.Context, name string, reference string, accept string) (acr.Manifest, error) {
	ret := _m.Called(ctx, name, reference, accept)

	var r0 acr.Manifest
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) acr.Manifest); ok {
		r0 = rf(ctx, name, reference, accept)
	} else {
		r0 = ret.Get(0).(acr.Manifest)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, name, reference, accept)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepositories provides a mock function with given fields: ctx, last, n
func (_m *BaseClientAPI) GetRepositories(ctx context.Context, last string, n *int32) (acr.Repositories, error) {
	ret := _m.Called(ctx, last, n)

	var r0 acr.Repositories
	if rf, ok := ret.Get(0).(func(context.Context, string, *int32) acr.Repositories); ok {
		r0 = rf(ctx, last, n)
	} else {
		r0 = ret.Get(0).(acr.Repositories)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *int32) error); ok {
		r1 = rf(ctx, last, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagList provides a mock function with given fields: ctx, name
func (_m *BaseClientAPI) GetTagList(ctx context.Context, name string) (acr.RepositoryTags, error) {
	ret := _m.Called(ctx, name)

	var r0 acr.RepositoryTags
	if rf, ok := ret.Get(0).(func(context.Context, string) acr.RepositoryTags); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(acr.RepositoryTags)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManifestMetadata provides a mock function with given fields: ctx, name, reference
func (_m *BaseClientAPI) ListManifestMetadata(ctx context.Context, name string, reference string) (acr.ManifestMetadataList, error) {
	ret := _m.Called(ctx, name, reference)

	var r0 acr.ManifestMetadataList
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.ManifestMetadataList); ok {
		r0 = rf(ctx, name, reference)
	} else {
		r0 = ret.Get(0).(acr.ManifestMetadataList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositoryMetadata provides a mock function with given fields: ctx, name
func (_m *BaseClientAPI) ListRepositoryMetadata(ctx context.Context, name string) (acr.RepositoryMetadata, error) {
	ret := _m.Called(ctx, name)

	var r0 acr.RepositoryMetadata
	if rf, ok := ret.Get(0).(func(context.Context, string) acr.RepositoryMetadata); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(acr.RepositoryMetadata)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagMetadata provides a mock function with given fields: ctx, name, reference
func (_m *BaseClientAPI) ListTagMetadata(ctx context.Context, name string, reference string) (acr.TagMetadataList, error) {
	ret := _m.Called(ctx, name, reference)

	var r0 acr.TagMetadataList
	if rf, ok := ret.Get(0).(func(context.Context, string, string) acr.TagMetadataList); ok {
		r0 = rf(ctx, name, reference)
	} else {
		r0 = ret.Get(0).(acr.TagMetadataList)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartBlobUpload provides a mock function with given fields: ctx, name, digest, from, mount
func (_m *BaseClientAPI) StartBlobUpload(ctx context.Context, name string, digest string, from string, mount string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, digest, from, mount)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, digest, from, mount)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, name, digest, from, mount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrManifestAttributes provides a mock function with given fields: ctx, name, reference, value
func (_m *BaseClientAPI) UpdateAcrManifestAttributes(ctx context.Context, name string, reference string, value *acr.ChangeableAttributes) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference, value)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *acr.ChangeableAttributes) autorest.Response); ok {
		r0 = rf(ctx, name, reference, value)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *acr.ChangeableAttributes) error); ok {
		r1 = rf(ctx, name, reference, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrManifestMetadata provides a mock function with given fields: ctx, name, reference, metadata, value
func (_m *BaseClientAPI) UpdateAcrManifestMetadata(ctx context.Context, name string, reference string, metadata string, value *interface{}) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference, metadata, value)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *interface{}) autorest.Response); ok {
		r0 = rf(ctx, name, reference, metadata, value)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *interface{}) error); ok {
		r1 = rf(ctx, name, reference, metadata, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrRepositoryAttributes provides a mock function with given fields: ctx, name, value
func (_m *BaseClientAPI) UpdateAcrRepositoryAttributes(ctx context.Context, name string, value *acr.ChangeableAttributes) (autorest.Response, error) {
	ret := _m.Called(ctx, name, value)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, *acr.ChangeableAttributes) autorest.Response); ok {
		r0 = rf(ctx, name, value)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *acr.ChangeableAttributes) error); ok {
		r1 = rf(ctx, name, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrRepositoryMetadata provides a mock function with given fields: ctx, name, metadata, value
func (_m *BaseClientAPI) UpdateAcrRepositoryMetadata(ctx context.Context, name string, metadata string, value *interface{}) (autorest.Response, error) {
	ret := _m.Called(ctx, name, metadata, value)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *interface{}) autorest.Response); ok {
		r0 = rf(ctx, name, metadata, value)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *interface{}) error); ok {
		r1 = rf(ctx, name, metadata, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrTagAttributes provides a mock function with given fields: ctx, name, reference, value
func (_m *BaseClientAPI) UpdateAcrTagAttributes(ctx context.Context, name string, reference string, value *acr.ChangeableAttributes) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference, value)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *acr.ChangeableAttributes) autorest.Response); ok {
		r0 = rf(ctx, name, reference, value)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *acr.ChangeableAttributes) error); ok {
		r1 = rf(ctx, name, reference, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcrTagMetadata provides a mock function with given fields: ctx, name, reference, metadata, value
func (_m *BaseClientAPI) UpdateAcrTagMetadata(ctx context.Context, name string, reference string, metadata string, value *interface{}) (autorest.Response, error) {
	ret := _m.Called(ctx, name, reference, metadata, value)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *interface{}) autorest.Response); ok {
		r0 = rf(ctx, name, reference, metadata, value)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *interface{}) error); ok {
		r1 = rf(ctx, name, reference, metadata, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadBlobContent provides a mock function with given fields: ctx, name, UUID
func (_m *BaseClientAPI) UploadBlobContent(ctx context.Context, name string, UUID string) (autorest.Response, error) {
	ret := _m.Called(ctx, name, UUID)

	var r0 autorest.Response
	if rf, ok := ret.Get(0).(func(context.Context, string, string) autorest.Response); ok {
		r0 = rf(ctx, name, UUID)
	} else {
		r0 = ret.Get(0).(autorest.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
