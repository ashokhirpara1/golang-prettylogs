package prettylogs

import (
	"errors"
	"fmt"
	"testing"
	"time"

	mocklogger "github.com/ashokhirpara1/golang-prettylogs/logger/mock"
	"github.com/golang/mock/gomock"
)

func TestInfo(t *testing.T) {
	logMessage := "testing info method message"

	ctlr := gomock.NewController(t)
	mockLogger := mocklogger.NewMockStorage(ctlr)
	logs = mockLogger

	mockLogger.EXPECT().Info(gomock.Any(), gomock.Eq(logMessage)).Times(1)

	Info(logMessage)
}

func TestError(t *testing.T) {
	logMessage := "error"
	logError := errors.New("failed to do something")
	logFullMessage := fmt.Sprintf("%s: %s", logMessage, logError.Error())

	ctlr := gomock.NewController(t)
	mockLogger := mocklogger.NewMockStorage(ctlr)
	logs = mockLogger

	mockLogger.EXPECT().Error(gomock.Any(), gomock.Eq(logFullMessage)).Times(1)

	Error(logMessage, logError)
}

func TestFatal(t *testing.T) {
	logMessage := "error"
	logError := errors.New("breaking changes")
	logFullMessage := fmt.Sprintf("%s: %s", logMessage, logError.Error())

	ctlr := gomock.NewController(t)
	mockLogger := mocklogger.NewMockStorage(ctlr)
	logs = mockLogger

	mockLogger.EXPECT().Fatal(gomock.Any(), gomock.Eq(logFullMessage)).Times(1)

	Fatal("TestFatal", logMessage, logError)
}

func TestEnter(t *testing.T) {
	logMethod := "TestEnter"
	logMessage := "Started"
	logFullMessage := fmt.Sprintf("%s %s", logMessage, logMethod)

	ctlr := gomock.NewController(t)
	mockLogger := mocklogger.NewMockStorage(ctlr)
	logs = mockLogger

	mockLogger.EXPECT().Info(gomock.Any(), gomock.Eq(logFullMessage)).Times(1)

	if startTime, method := Enter(); method != logMethod {
		t.Errorf("Enter() = %v, %s, didn't return %s", startTime, method, logMethod)
	}
}

func TestExit(t *testing.T) {
	logMethod := "TestExit"
	startTime := time.Now()

	ctlr := gomock.NewController(t)
	mockLogger := mocklogger.NewMockStorage(ctlr)
	logs = mockLogger

	mockLogger.EXPECT().Info(gomock.Any(), gomock.Any()).Times(1)

	Exit(startTime, logMethod)
}
