// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/nuomizi-fw/stargazer/service"
)

// Ensure, that DownloaderServiceMock does implement service.DownloaderService.
// If this is not the case, regenerate this file with moq.
var _ service.DownloaderService = &DownloaderServiceMock{}

// DownloaderServiceMock is a mock implementation of service.DownloaderService.
//
//	func TestSomethingThatUsesDownloaderService(t *testing.T) {
//
//		// make and configure a mocked service.DownloaderService
//		mockedDownloaderService := &DownloaderServiceMock{
//		}
//
//		// use mockedDownloaderService in code that requires service.DownloaderService
//		// and then make assertions.
//
//	}
type DownloaderServiceMock struct {
	// calls tracks calls to the methods.
	calls struct {
	}
}
