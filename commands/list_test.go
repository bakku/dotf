package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf"
	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestList_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.List(m)

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestList_ShouldAbortIfDotfileDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)

	err := commands.List(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestList_ShouldFailIfConfigCannotBeRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return(nil, errors.New("error"))

	err := commands.List(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestList_ShouldFailIfConfigCannotBeDeserialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(errors.New("error"))

	err := commands.List(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestList_ShouldCorrectlyListAllTrackedFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := dotf.Config{
		"",
		false,
		[]dotf.TrackedFile{
			{".vimrc", "/home/.vimrc"},
			{".emacs.d/init.el", "/home/.emacs.d/init.el"},
		},
	}

	expectedTableString := "" +
		"+------------------------+------------------+\n" +
		"|          FILE          |   PATH IN REPO   |\n" +
		"+------------------------+------------------+\n" +
		"| /home/.vimrc           | .vimrc           |\n" +
		"| /home/.emacs.d/init.el | .emacs.d/init.el |\n" +
		"+------------------------+------------------+\n"

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, cfg).
		Return(nil)
	m.EXPECT().Log(expectedTableString)

	err := commands.List(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
