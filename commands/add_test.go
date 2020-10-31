package commands_test

import (
	"errors"
	"testing"

	"bakku.dev/dotf"
	"bakku.dev/dotf/commands"
	"bakku.dev/dotf/mocks"
	"github.com/golang/mock/gomock"
)

func TestAdd_ShouldFailIfNoHomeVarExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("")

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err to not be nil")
	}
}

func TestAdd_ShouldAbortIfDotfileDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(false)

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfFirstReadLineReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return("", errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfExpandPathReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return(".vimrc", nil)
	m.EXPECT().ExpandPath(".vimrc").Return("", errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfSecondReadLineReturnsAnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return(".vimrc", nil)
	m.EXPECT().ExpandPath(".vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().Log("Please insert the path of the file inside the repo (default .vimrc): ")
	m.EXPECT().ReadLine().Return("", errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfConfigCannotBeRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().Log("Please insert the path of the file inside the repo (default .vimrc): ")
	m.EXPECT().ReadLine().Return("", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return(nil, errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfConfigCannotBeDeserialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().Log("Please insert the path of the file inside the repo (default .vimrc): ")
	m.EXPECT().ReadLine().Return("", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfConfigCannotBeSerialized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().Log("Please insert the path of the file inside the repo (default .vimrc): ")
	m.EXPECT().ReadLine().Return("", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(nil)
	m.EXPECT().SerializeConfig(gomock.Eq(dotf.Config{"", false, []dotf.TrackedFile{{".vimrc", "/home/.vimrc"}}})).Return(nil, errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldFailIfConfigCannotBeWritten(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().Log("Please insert the path of the file inside the repo (default .vimrc): ")
	m.EXPECT().ReadLine().Return("", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(nil)
	m.EXPECT().SerializeConfig(gomock.Eq(dotf.Config{"", false, []dotf.TrackedFile{{".vimrc", "/home/.vimrc"}}})).Return([]byte("ABC"), nil)
	m.EXPECT().WriteFile("/home/.dotf", []byte("ABC")).Return(errors.New("error"))

	err := commands.Add(m)

	if err == nil {
		t.Fatalf("Expected err not to be nil")
	}
}

func TestAdd_ShouldWorkSuccessfullyIfNoErrorsOccur(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSysOpsProvider(ctrl)

	m.EXPECT().GetEnvVar(gomock.Eq("HOME")).Return("/home/")
	m.EXPECT().GetPathSep().Return("/")
	m.EXPECT().CleanPath("/home//.dotf").Return("/home/.dotf")
	m.EXPECT().PathExists(gomock.Eq("/home/.dotf")).Return(true)
	m.EXPECT().Log("Please insert the path of the file you want to track: ")
	m.EXPECT().ReadLine().Return("/home//.vimrc", nil)
	m.EXPECT().ExpandPath("/home//.vimrc").Return("/home/.vimrc", nil)
	m.EXPECT().Log("Please insert the path of the file inside the repo (default .vimrc): ")
	m.EXPECT().ReadLine().Return("", nil)
	m.EXPECT().ReadFile(gomock.Eq("/home/.dotf")).Return([]byte("ABC"), nil)
	m.EXPECT().DeserializeConfig(gomock.Eq([]byte("ABC")), gomock.AssignableToTypeOf(&dotf.Config{})).Return(nil)
	m.EXPECT().SerializeConfig(gomock.Eq(dotf.Config{"", false, []dotf.TrackedFile{{".vimrc", "/home/.vimrc"}}})).Return([]byte("ABC"), nil)
	m.EXPECT().WriteFile("/home/.dotf", []byte("ABC")).Return(nil)

	err := commands.Add(m)

	if err != nil {
		t.Fatalf("Expected err to be nil")
	}
}
