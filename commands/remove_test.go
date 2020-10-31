package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf"
	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestRemove_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestRemove_ShouldAbortIfDotfileDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfReadLineReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("", errors.New("error"))

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfExpandPathReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return(".vimrc", nil)
	m.EXPECT().ExpandPath(".vimrc").Return("", errors.New("error"))

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfConfigCannotBeRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return(nil, errors.New("error"))

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfConfigCannotBeDeserialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(errors.New("error"))

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfFileIsNotTracked(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, dotf.Config{"", false, []dotf.TrackedFile{}}).
		Return(nil)

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfConfigCannotBeSerialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, dotf.Config{"", false, []dotf.TrackedFile{{".vimrc", "/home/.vimrc"}}}).
		Return(nil)

	m.EXPECT().SerializeConfig(gomock.Eq(dotf.Config{"", false, []dotf.TrackedFile{}})).Return(nil, errors.New("error"))

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldFailIfConfigCannotBeWritten(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, dotf.Config{"", false, []dotf.TrackedFile{{".vimrc", "/home/.vimrc"}}}).
		Return(nil)

	m.EXPECT().SerializeConfig(gomock.Eq(dotf.Config{"", false, []dotf.TrackedFile{}})).Return([]byte("ABC"), nil)
	m.EXPECT().WriteFile("/home/.dotf", []byte("ABC")).Return(errors.New("error"))

	err := commands.Remove(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestRemove_ShouldWorkSuccessfullyIfNoErrorsOccur(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want dotf to stop tracking: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().
		DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).
		SetArg(1, dotf.Config{"", false, []dotf.TrackedFile{{".vimrc", "/home/.vimrc"}}}).
		Return(nil)

	m.EXPECT().SerializeConfig(gomock.Eq(dotf.Config{"", false, []dotf.TrackedFile{}})).Return([]byte("ABC"), nil)
	m.EXPECT().WriteFile("/home/.dotf", []byte("ABC")).Return(nil)

	err := commands.Remove(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
