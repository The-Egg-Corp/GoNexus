package v1

import (
	"testing"

	v1 "github.com/the-egg-corp/gonexus/v1"
)

const testModGame = "HogwartsLegacy"
const testModId = 464
const testModFileId = 1444

var ModService = v1.NewModService(NexusClient)

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
}

func TestGetModFiles(t *testing.T) {
	files, err := ModService.GetModFiles(testModGame, testModId)
	if err != nil {
		t.Fatalf("failed to get files for mod: %d\n%v", testModId, err)
	}

	if files == nil {
		t.Fatalf("failed to get files for mod: %d. response was nil but no error was provided", testModId)
	}
}

func TestGetModFileByID(t *testing.T) {
	file, err := ModService.GetModFileByID(testModGame, testModId, testModFileId)
	if err != nil {
		t.Fatalf("failed to get file %d for mod: %d\n%v", testModFileId, testModId, err)
	}

	if file == nil {
		t.Fatalf("failed to get file %d for mod: %d. response was nil but no error was provided", testModFileId, testModId)
	}
}

func TestGetTrending(t *testing.T) {
	trending, err := ModService.GetTrending(testModGame)
	if err != nil {
		t.Fatalf("failed to get trending mods for game: %s\n%v", testModGame, err)
	}

	if trending == nil {
		t.Fatalf("failed to get trending mods for game: %s. response was nil but no error was provided", testModGame)
	}
}

func TestGetLatestAdded(t *testing.T) {
	latestAdded, err := ModService.GetLatestAdded(testModGame)
	if err != nil {
		t.Fatalf("failed to get latest added mods for game: %s\n%v", testModGame, err)
	}

	if latestAdded == nil {
		t.Fatalf("failed to get latest added mods for game: %s. response was nil but no error was provided", testModGame)
	}
}

func TestGetLatestUpdated(t *testing.T) {
	latestUpdated, err := ModService.GetLatestUpdated(testModGame)
	if err != nil {
		t.Fatalf("failed to get latest updated mods for game: %s\n%v", testModGame, err)
	}

	if latestUpdated == nil {
		t.Fatalf("failed to get latest updated mods for game: %s. response was nil but no error was provided", testModGame)
	}
}

// If this works, 1w and 1m are likely to also.
func TestGetUpdatedWithinPeriod1d(t *testing.T) {
	updatedWithinPeriod, err := ModService.GetUpdatedWithinPeriod(testModGame, v1.MOD_UPDATE_PERIOD_ONE_DAY)
	if err != nil {
		t.Fatalf("failed to get updated mods within 1d period for game: %s\n%v", testModGame, err)
	}

	if updatedWithinPeriod == nil {
		t.Fatalf("failed to get mods updated within 1d period for game: %s. response was nil but no error was provided", testModGame)
	}
}
