package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf"
	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestInit_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.Init(m, "")

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestInit_ShouldFailIfExpandPathReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().ExpandPath("invalid").Return("", errors.New("error"))

	err := commands.Init(m, "invalid")

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
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().ExpandPath("invalid").Return("invalid", nil)
	m.EXPECT().PathExists(gomock.Eq("invalid")).Return(false)

	err := commands.Init(m, "invalid")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestInit_ShouldFailIfReadLineReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().ExpandPath("repo").Return("/home/repo", nil)
	m.EXPECT().PathExists(gomock.Eq("/home/repo")).Return(true)
	m.EXPECT().Log("Do you want to create backups of your dotfiles when pulling? (y/n): ")
	m.EXPECT().ReadLine().Return("", errors.New("error"))

	err := commands.Init(m, "repo")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestInit_ShouldFailIfConfigCannotBeSerialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().ExpandPath("/home/repo").Return("/home/repo", nil)
	m.EXPECT().PathExists(gomock.Eq("/home/repo")).Return(true)
	m.EXPECT().Log("Do you want to create backups of your dotfiles when pulling? (y/n): ")
	m.EXPECT().ReadLine().Return("blah", nil)
	m.EXPECT().Log("Do you want to create backups of your dotfiles when pulling? (y/n): ")
	m.EXPECT().ReadLine().Return("y", nil)
	m.EXPECT().
		SerializeConfig(gomock.Eq(dotf.Config{"/home/repo", true, []dotf.TrackedFile{}})).
		Return([]byte{}, errors.New("error"))

	err := commands.Init(m, "/home/repo")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestInit_ShouldFailIfConfigCannotBeWrittenToFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().ExpandPath("/home/repo").Return("/home/repo", nil)
	m.EXPECT().PathExists(gomock.Eq("/home/repo")).Return(true)
	m.EXPECT().Log("Do you want to create backups of your dotfiles when pulling? (y/n): ")
	m.EXPECT().ReadLine().Return("n", nil)
	m.EXPECT().
		SerializeConfig(gomock.Eq(dotf.Config{"/home/repo", false, []dotf.TrackedFile{}})).
		Return([]byte("ABC"), nil)
	m.EXPECT().WriteFile("/home/.dotf", []byte("ABC")).Return(errors.New("error"))

	err := commands.Init(m, "/home/repo")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestInit_ShouldTerminateSuccessfullyIfNoErrorIsRaised(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)
	m.EXPECT().ExpandPath("/home/repo").Return("/home/repo", nil)
	m.EXPECT().PathExists(gomock.Eq("/home/repo")).Return(true)
	m.EXPECT().Log("Do you want to create backups of your dotfiles when pulling? (y/n): ")
	m.EXPECT().ReadLine().Return("y", nil)
	m.EXPECT().
		SerializeConfig(gomock.Eq(dotf.Config{"/home/repo", true, []dotf.TrackedFile{}})).
		Return([]byte("ABC"), nil)
	m.EXPECT().WriteFile("/home/.dotf", []byte("ABC")).Return(nil)
	m.EXPECT().Log("Successfully created file at /home/.dotf\n")

	err := commands.Init(m, "/home/repo")

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
	m.EXPECT().CleanPath("/home/.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log(gomock.Eq("/home/.dotf already exists\n"))

	err := commands.Init(m, "")

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
