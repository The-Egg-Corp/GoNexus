# GoNexus
Unofficial Golang client for [NexusMods](https://www.nexusmods.com).\
Provides a convenient way to interface with the API to avoid sending HTTP requests manually.

> [!IMPORTANT]
> This client was made with the intention to be used with a *personal* API key for testing and development purposes. However, it is possible that it also can be used in production, but extra caution should be taken as it has not been tested in any real application thus far. 
> 
> If you have registered your application with **NexusMods** and it has been publicly released, the [Single Sign-On (SSO)](https://github.com/Nexus-Mods/sso-integration-demo) system should be used to retrieve the user's key instead of using your own personal one.

## Install and Import
```console
go get github.com/the-egg-corp/gonexus
```

```go 
import (
    v1 "github.com/the-egg-corp/gonexus/v1"
)
```

## Setup
> [!NOTE]
> An API key is required to return any data!\
> You can generate and use the one in your [account settings](https://next.nexusmods.com/settings/api-keys) if the [Acceptable Use Policy](https://help.nexusmods.com/article/114-api-acceptable-use-policy) allows you under your current circumstances.
>
> It is recommended to store your key safely using environment variables.

### Using system variables
```go
var NexusClient, ClientError = v1.NewNexusClient(os.Getenv("NEXUS_KEY"))
```

### Using an `.env` file
```go
var NexusClient, ClientError = NewNexusClient()

func NewNexusClient() (*v1.NexusClient, error) {
	vars, err := godotenv.Read(".env") // Or a path to the env file. Ex: "../.env"
	if err != nil {
		log.Fatalf("\nFailed to load required environment variables.\n%s", err)
	}

	key, found := vars["NEXUS_KEY"]
	if !found {
		log.Fatalf("\nCould not find required environment variable: NEXUS_KEY\n%s", err)
	}

	return v1.NewNexusClient(key)
}
```

## Usage
Until there is proper documentation, you can look at the [tests](./tests/) for some usage examples.

### Basic Example
```go
const GAME_NAME = "HogwartsLegacy"

var NexusClient, ClientError = v1.NewNexusClient(os.Getenv("NEXUS_KEY"))

func main() {
    if ClientError != nil {
        // Handle error during client creation. Likely an invalid API key.
    }

    gameService := v1.NewGameService(NexusClient)
    modService := v1.NewModService(NexusClient)

    game, err := gameService.GetGame(GAME_NAME)
    if err != nil {
        // Handle game request error, including 404 "Not Found".
    }

    mod, err := modService.GetModByID(GAME_NAME, 69)
    if err != nil {
        // Handle mod request error, including 404 "Not Found".
    }

    // Contains `Files` and `FileUpdates` (see ModFile struct).
    modFilesResponse, err := modService.GetModFiles(GAME_NAME, 69) 
    if modFiles != nil {
        latestFile := modFilesResponse.Files[0]
        fmt.Printf("Description of latest mod file:\n%s", latestFile.Description)
    }
}
```

## Contact
Feel free to join my [discord](https://discord.gg/BwfzZpytjf) for support or suggestions.
