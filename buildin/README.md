# buildin

## New Feature: User API

The buildin package now supports fetching user tasks, bugs, and stories.

### Quick Start

```go
import "github.com/easysoft/go-zentao/v21/buildin"

// Create client with username and password
client, err := buildin.NewBasicAuthClient(
    "your_username",
    "your_password", 
    buildin.WithBaseURL("https://your-zentao-url.com"),
)
if err != nil {
    log.Fatal(err)
}

// Get all user data (tasks + bugs + stories)
result, _, err := client.GetUserTasksBugsStories()
if err != nil {
    log.Fatal(err)
}

// Display user info
fmt.Printf("User: %s (%s)\n", result.Profile.Realname, result.Profile.Account)

// Get counts
taskCount, bugCount, storyCount := result.GetItemCounts()
fmt.Printf("Tasks: %d, Bugs: %d, Stories: %d\n", taskCount, bugCount, storyCount)

// Get only specific types
tasks, _, err := client.GetUserTasks()        // Tasks only
bugs, _, err := client.GetUserBugs()          // Bugs only  
stories, _, err := client.GetUserStories()    // Stories only
```

### Features

- ✅ Get user's tasks, bugs, and stories
- ✅ Support for fetching individual types
- ✅ Built-in filtering methods (by status)
- ✅ Statistics and analysis methods
- ✅ Today's completed items query

See example at `./example/buildin-user/main.go`

## Usage

```bash
go get github.com/easysoft/go-zentao/v21/buildin
```
