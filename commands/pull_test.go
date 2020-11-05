package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf"
	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestPull_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.Pull(m)

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestPull_ShouldAbortIfDotfileDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)

	err := commands.Pull(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPull_ShouldFailIfConfigCannotBeRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return(nil, errors.New("error"))

	err := commands.Pull(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPull_ShouldFailIfConfigCannotBeDeserialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(errors.New("error"))

	err := commands.Pull(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPull_ShouldFailIfRepositoryUpdateFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		false,
		[]dotf.TrackedFile{},
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
	m.EXPECT().UpdateRepo("/home/repo").Return(errors.New("error"))

	err := commands.Pull(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPull_ShouldFailIfCopyFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		false,
		[]dotf.TrackedFile{
			{".vimrc", "/home/.vimrc"},
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
	m.EXPECT().UpdateRepo("/home/repo").Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo/.vimrc").Return("/home/repo/.vimrc")
	m.EXPECT().CopyFile("/home/repo/.vimrc", "/home/.vimrc").Return(errors.New("error"))

	err := commands.Pull(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestPull_ShouldBehaveCorrectlyForNoBackups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		false,
		[]dotf.TrackedFile{
			{".vimrc", "/home/.vimrc"},
			{".emacs.d/init.el", "/home/.emacs.d/init.el"},
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
	m.EXPECT().UpdateRepo("/home/repo").Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo/.vimrc").Return("/home/repo/.vimrc")
	m.EXPECT().CopyFile("/home/repo/.vimrc", "/home/.vimrc").Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo/.emacs.d/init.el").Return("/home/repo/.emacs.d/init.el")
	m.EXPECT().CopyFile("/home/repo/.emacs.d/init.el", "/home/.emacs.d/init.el").Return(nil)

	err := commands.Pull(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}

func TestPull_ShouldBehaveCorrectlyForBackups(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	cfg := dotf.Config{
		"/home/repo",
		true,
		[]dotf.TrackedFile{
			{".vimrc", "/home/.vimrc"},
			{".emacs.d/init.el", "/home/.emacs.d/init.el"},
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
	m.EXPECT().UpdateRepo("/home/repo").Return(nil)
	m.EXPECT().CopyFile("/home/.vimrc", "/home/.vimrc.bk").Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo/.vimrc").Return("/home/repo/.vimrc")
	m.EXPECT().CopyFile("/home/repo/.vimrc", "/home/.vimrc").Return(nil)
	m.EXPECT().CopyFile("/home/.emacs.d/init.el", "/home/.emacs.d/init.el.bk").Return(nil)
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home/repo/.emacs.d/init.el").Return("/home/repo/.emacs.d/init.el")
	m.EXPECT().CopyFile("/home/repo/.emacs.d/init.el", "/home/.emacs.d/init.el").Return(nil)

	err := commands.Pull(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
