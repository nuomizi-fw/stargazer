// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/nuomizi-fw/stargazer/service"
)

// Ensure, that BittorrentServiceMock does implement service.BittorrentService.
// If this is not the case, regenerate this file with moq.
var _ service.BittorrentService = &BittorrentServiceMock{}

// BittorrentServiceMock is a mock implementation of service.BittorrentService.
//
//	func TestSomethingThatUsesBittorrentService(t *testing.T) {
//
//		// make and configure a mocked service.BittorrentService
//		mockedBittorrentService := &BittorrentServiceMock{
//		}
//
//		// use mockedBittorrentService in code that requires service.BittorrentService
//		// and then make assertions.
//
//	}
type BittorrentServiceMock struct {
	// calls tracks calls to the methods.
	calls struct {
	}
}