package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestInit_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.Init(m)

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestInit_ShouldFailIfReadLineReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().Log("Insert path to dotfile repo: ")
	m.EXPECT().ReadLine().Return("", errors.New("ERROR!!"))

	err := commands.Init(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestInit_ShouldFailIfRepoDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().Log("Insert path to dotfile repo: ")
	m.EXPECT().ReadLine().Return("invalid", nil)
	m.EXPECT().PathExists(gomock.Eq("invalid")).Return(false)
	
	err := commands.Init(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestInit_ShouldAbortIfDotfFileAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log(gomock.Eq("/home/.dotf already exists\n"))

	err := commands.Init(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
