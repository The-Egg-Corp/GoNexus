# GoNexus
Unofficial Golang client for [NexusMods](https://www.nexusmods.com).\
Provides a convenient way to interface with the API to avoid sending HTTP requests manually.

## Install and Import
```console
go get github.com/the-egg-corp/gonexus
```

```go 
import (
    "github.com/the-egg-corp/gonexus/v1"
)
```

## Setup
> [!NOTE]
> A personal API key is required to return any data!\
> You can get generate one in your [account settings](https://next.nexusmods.com/settings/api-keys). 
>
> It is recommended to store your key safely using environment variables.

```go
var NexusClient, ClientError = v1.NewNexusClient(os.Getenv("NEXUS_KEY"))
```

## Usage
Until there is proper documentation, you can look at the [tests](./v1/tests/) for some usage examples.

## Contact
Feel free to join my [discord](https://discord.gg/BwfzZpytjf) for support or suggestions.
