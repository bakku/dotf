package commands_test

import (
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

func TestInit_ShouldTryToCreateDotfFileIfNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath(gomock.Eq("/home/.dotf")).Return("/home/.dotf")
	m.EXPECT().FileExists(gomock.Eq("/home/.dotf")).Return(false)
	
	err := commands.Init(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}

func TestInit_ShouldAbortIfDotfFileAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath(gomock.Eq("/home/.dotf")).Return("/home/.dotf")
	m.EXPECT().FileExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log(gomock.Eq("/home/.dotf already exists"))

	err := commands.Init(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
