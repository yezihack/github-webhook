package internal

type PayloadPong struct {
	Ok    bool   `json:"ok"`
	Event string `json:"event"`
	Error string `json:"error,omitempty"`
}

// repo info
type GitHubRepo struct {
	Name        string // repository name
	FullName    string // repository full name
	CloneURL    string // repository url
	CommitID    string // push commit id
	CommitName  string // push repo name
	CommitEmail string // push repo email
	CommitAt    string // push time
	BranchName  string // branch name
}
