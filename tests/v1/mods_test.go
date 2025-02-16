package v1

import (
	"testing"

	v1 "github.com/the-egg-corp/gonexus/v1"
)

var ModService = v1.NewModService(NexusClient)

const testModGame = "HogwartsLegacy"
const testModId = 464
const testModFileId = 1444

func TestGetModByID(t *testing.T) {
	mod, err := ModService.GetModByID(testModGame, testModId)
	if err != nil {
		t.Fatalf("failed to get mod by id: %d\n%v", testModId, err)
	}

	if mod == nil {
		t.Fatalf("failed to get mod by id: %d. mod is nil but no error was provided", testModId)
	}

	//util.PrettyPrint(mod)
}

func TestGetModChangelogs(t *testing.T) {
	changelogs, err := ModService.GetModChangelogs(testModGame, testModId)
	if err != nil {
		t.Fatalf("failed to get changelogs for mod: %d\n%v", testModId, err)
	}

	if changelogs == nil {
		t.Fatalf("failed to get changelogs for mod: %d. response was nil but no error was provided", testModId)
	}

	//util.PrettyPrint(mod)
}

func TestGetModFiles(t *testing.T) {
	files, err := ModService.GetModFiles(testModGame, testModId)
	if err != nil {
		t.Fatalf("failed to get files for mod: %d\n%v", testModId, err)
	}

	if files == nil {
		t.Fatalf("failed to get files for mod: %d. response was nil but no error was provided", testModId)
	}

	//util.PrettyPrint(mod)
}

func TestGetModFile(t *testing.T) {
	file, err := ModService.GetModFileByID(testModGame, testModId, testModFileId)
	if err != nil {
		t.Fatalf("failed to get file %d for mod: %d\n%v", testModFileId, testModId, err)
	}

	if file == nil {
		t.Fatalf("failed to get file %d for mod: %d. response was nil but no error was provided", testModFileId, testModId)
	}

	//util.PrettyPrint(mod)
}
