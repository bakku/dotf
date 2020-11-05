package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf"
	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestPush_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.Push(m, "")

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestPush_ShouldAbortIfDotfileDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)

	err := commands.Push(m, "")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPush_ShouldFailIfConfigCannotBeRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return(nil, errors.New("error"))

	err := commands.Push(m, "")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPush_ShouldFailIfConfigCannotBeDeserialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(errors.New("error"))

	err := commands.Push(m, "")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPush_ShouldFailIfCopyFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		false,
		[]dotf.TrackedFile{
			{"/.vimrc", "/home/.vimrc"},
		},
	}

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, cfg).
		Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo//.vimrc").Return("/home/repo/.vimrc")
	m.EXPECT().CopyFile("/home/.vimrc", "/home/repo/.vimrc").Return(errors.New("error"))

	err := commands.Push(m, "")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPush_ShouldFailIfPushFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		false,
		[]dotf.TrackedFile{
			{"/.vimrc", "/home/.vimrc"},
		},
	}

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, cfg).
		Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo//.vimrc").Return("/home/repo/.vimrc")
	m.EXPECT().CopyFile("/home/.vimrc", "/home/repo/.vimrc").Return(nil)
	m.EXPECT().CommitRepo("/home/repo", "Update .vimrc").Return(errors.New("error"))

	err := commands.Push(m, "Update .vimrc")

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPush_ShouldWorkSuccessfullyWithoutErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		false,
		[]dotf.TrackedFile{
			{"/.vimrc", "/home/.vimrc"},
		},
	}

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, cfg).
		Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo//.vimrc").Return("/home/repo/.vimrc")
	m.EXPECT().CopyFile("/home/.vimrc", "/home/repo/.vimrc").Return(nil)
	m.EXPECT().CommitRepo("/home/repo", "Update .vimrc").Return(nil)

	err := commands.Push(m, "Update .vimrc")

	if err != nil {
		t.Fatalf("Expected err not to be nil")
	}
}
