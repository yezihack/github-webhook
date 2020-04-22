package internal

import (
	"github.com/tidwall/gjson"
	"testing"
)

func TestDeserialize(t *testing.T) {
	payload := _unmarshal(t)
	if payload.Name != "barry" {
		t.Errorf("name parsed as expected: %s", payload.Name)
	}

	if payload.CloneURL != "https://github.com/yezihack/barry.git" {
		t.Errorf("clone url not parsed as expected: %s", payload.CloneURL)
	}

	if payload.CommitEmail != "freeit@126.com" {
		t.Errorf("CommitEmail is not parsed as expected: %s", payload.CommitEmail)
	}

	if payload.BranchName != "refs/heads/www" {
		t.Errorf("FullName is not parsed as expected: %s", payload.BranchName)
	}
}

func _unmarshal(t *testing.T) GitHubRepo {
	repo := GitHubRepo{}
	result := gjson.Parse(jsonPayload)
	repo.Name = result.Get("repository.name").Str
	repo.FullName = result.Get("repository.full_name").Str
	repo.CloneURL = result.Get("repository.clone_url").Str
	repo.CommitID = result.Get("head_commit.id").Str
	repo.CommitName = result.Get("head_commit.committer.name").Str
	repo.CommitEmail = result.Get("head_commit.committer.email").Str
	repo.CommitAt = result.Get("head_commit.timestamp").Time().Format("2006-01-02 15:04:05")
	repo.BranchName = result.Get("ref").Str
	return repo
}

const jsonPayload string = `
{
    "ref":"refs/heads/www",
    "before":"faccedbf3fa1bbbcaf6270ed6e354a0c2260badd",
    "after":"f3a91f5c6a27a6239cd3950a0b7451c187907023",
    "repository":{
        "id":247632243,
        "node_id":"MDEwOlJlcG9zaXRvcnkyNDc2MzIyNDM=",
        "name":"barry",
        "full_name":"yezihack/barry",
        "private":false,
        "owner":{
            "name":"yezihack",
            "email":"freeit@126.com",
            "login":"yezihack",
            "id":11530866,
            "node_id":"MDQ6VXNlcjExNTMwODY2",
            "avatar_url":"https://avatars0.githubusercontent.com/u/11530866?v=4",
            "gravatar_id":"",
            "url":"https://api.github.com/users/yezihack",
            "html_url":"https://github.com/yezihack",
            "followers_url":"https://api.github.com/users/yezihack/followers",
            "following_url":"https://api.github.com/users/yezihack/following{/other_user}",
            "gists_url":"https://api.github.com/users/yezihack/gists{/gist_id}",
            "starred_url":"https://api.github.com/users/yezihack/starred{/owner}{/repo}",
            "subscriptions_url":"https://api.github.com/users/yezihack/subscriptions",
            "organizations_url":"https://api.github.com/users/yezihack/orgs",
            "repos_url":"https://api.github.com/users/yezihack/repos",
            "events_url":"https://api.github.com/users/yezihack/events{/privacy}",
            "received_events_url":"https://api.github.com/users/yezihack/received_events",
            "type":"User",
            "site_admin":false
        },
        "html_url":"https://github.com/yezihack/barry",
        "description":null,
        "fork":false,
        "url":"https://github.com/yezihack/barry",
        "forks_url":"https://api.github.com/repos/yezihack/barry/forks",
        "keys_url":"https://api.github.com/repos/yezihack/barry/keys{/key_id}",
        "collaborators_url":"https://api.github.com/repos/yezihack/barry/collaborators{/collaborator}",
        "teams_url":"https://api.github.com/repos/yezihack/barry/teams",
        "hooks_url":"https://api.github.com/repos/yezihack/barry/hooks",
        "issue_events_url":"https://api.github.com/repos/yezihack/barry/issues/events{/number}",
        "events_url":"https://api.github.com/repos/yezihack/barry/events",
        "assignees_url":"https://api.github.com/repos/yezihack/barry/assignees{/user}",
        "branches_url":"https://api.github.com/repos/yezihack/barry/branches{/branch}",
        "tags_url":"https://api.github.com/repos/yezihack/barry/tags",
        "blobs_url":"https://api.github.com/repos/yezihack/barry/git/blobs{/sha}",
        "git_tags_url":"https://api.github.com/repos/yezihack/barry/git/tags{/sha}",
        "git_refs_url":"https://api.github.com/repos/yezihack/barry/git/refs{/sha}",
        "trees_url":"https://api.github.com/repos/yezihack/barry/git/trees{/sha}",
        "statuses_url":"https://api.github.com/repos/yezihack/barry/statuses/{sha}",
        "languages_url":"https://api.github.com/repos/yezihack/barry/languages",
        "stargazers_url":"https://api.github.com/repos/yezihack/barry/stargazers",
        "contributors_url":"https://api.github.com/repos/yezihack/barry/contributors",
        "subscribers_url":"https://api.github.com/repos/yezihack/barry/subscribers",
        "subscription_url":"https://api.github.com/repos/yezihack/barry/subscription",
        "commits_url":"https://api.github.com/repos/yezihack/barry/commits{/sha}",
        "git_commits_url":"https://api.github.com/repos/yezihack/barry/git/commits{/sha}",
        "comments_url":"https://api.github.com/repos/yezihack/barry/comments{/number}",
        "issue_comment_url":"https://api.github.com/repos/yezihack/barry/issues/comments{/number}",
        "contents_url":"https://api.github.com/repos/yezihack/barry/contents/{+path}",
        "compare_url":"https://api.github.com/repos/yezihack/barry/compare/{base}...{head}",
        "merges_url":"https://api.github.com/repos/yezihack/barry/merges",
        "archive_url":"https://api.github.com/repos/yezihack/barry/{archive_format}{/ref}",
        "downloads_url":"https://api.github.com/repos/yezihack/barry/downloads",
        "issues_url":"https://api.github.com/repos/yezihack/barry/issues{/number}",
        "pulls_url":"https://api.github.com/repos/yezihack/barry/pulls{/number}",
        "milestones_url":"https://api.github.com/repos/yezihack/barry/milestones{/number}",
        "notifications_url":"https://api.github.com/repos/yezihack/barry/notifications{?since,all,participating}",
        "labels_url":"https://api.github.com/repos/yezihack/barry/labels{/name}",
        "releases_url":"https://api.github.com/repos/yezihack/barry/releases{/id}",
        "deployments_url":"https://api.github.com/repos/yezihack/barry/deployments",
        "created_at":1584341864,
        "updated_at":"2020-03-16T07:41:35Z",
        "pushed_at":1587375882,
        "git_url":"git://github.com/yezihack/barry.git",
        "ssh_url":"git@github.com:yezihack/barry.git",
        "clone_url":"https://github.com/yezihack/barry.git",
        "svn_url":"https://github.com/yezihack/barry",
        "homepage":null,
        "size":2544,
        "stargazers_count":0,
        "watchers_count":0,
        "language":null,
        "has_issues":true,
        "has_projects":true,
        "has_downloads":true,
        "has_wiki":true,
        "has_pages":false,
        "forks_count":0,
        "mirror_url":null,
        "archived":false,
        "disabled":false,
        "open_issues_count":0,
        "license":null,
        "forks":0,
        "open_issues":0,
        "watchers":0,
        "default_branch":"master",
        "stargazers":0,
        "master_branch":"master"
    },
    "pusher":{
        "name":"yezihack",
        "email":"freeit@126.com"
    },
    "sender":{
        "login":"yezihack",
        "id":11530866,
        "node_id":"MDQ6VXNlcjExNTMwODY2",
        "avatar_url":"https://avatars0.githubusercontent.com/u/11530866?v=4",
        "gravatar_id":"",
        "url":"https://api.github.com/users/yezihack",
        "html_url":"https://github.com/yezihack",
        "followers_url":"https://api.github.com/users/yezihack/followers",
        "following_url":"https://api.github.com/users/yezihack/following{/other_user}",
        "gists_url":"https://api.github.com/users/yezihack/gists{/gist_id}",
        "starred_url":"https://api.github.com/users/yezihack/starred{/owner}{/repo}",
        "subscriptions_url":"https://api.github.com/users/yezihack/subscriptions",
        "organizations_url":"https://api.github.com/users/yezihack/orgs",
        "repos_url":"https://api.github.com/users/yezihack/repos",
        "events_url":"https://api.github.com/users/yezihack/events{/privacy}",
        "received_events_url":"https://api.github.com/users/yezihack/received_events",
        "type":"User",
        "site_admin":false
    },
    "created":false,
    "deleted":false,
    "forced":false,
    "base_ref":null,
    "compare":"https://github.com/yezihack/barry/compare/faccedbf3fa1...f3a91f5c6a27",
    "commits":[
        {
            "id":"f3a91f5c6a27a6239cd3950a0b7451c187907023",
            "tree_id":"d64bd14683c278667dbba2f86ac62b196c2ebf44",
            "distinct":true,
            "message":"fix",
            "timestamp":"2020-04-20T17:44:31+08:00",
            "url":"https://github.com/yezihack/barry/commit/f3a91f5c6a27a6239cd3950a0b7451c187907023",
            "author":{
                "name":"barry",
                "email":"freeit@126.com",
                "username":"yezihack"
            },
            "committer":{
                "name":"barry",
                "email":"freeit@126.com",
                "username":"yezihack"
            },
            "added":[

            ],
            "removed":[

            ],
            "modified":[
                "README.md"
            ]
        }
    ],
    "head_commit":{
        "id":"f3a91f5c6a27a6239cd3950a0b7451c187907023",
        "tree_id":"d64bd14683c278667dbba2f86ac62b196c2ebf44",
        "distinct":true,
        "message":"fix",
        "timestamp":"2020-04-20T17:44:31+08:00",
        "url":"https://github.com/yezihack/barry/commit/f3a91f5c6a27a6239cd3950a0b7451c187907023",
        "author":{
            "name":"barry",
            "email":"freeit@126.com",
            "username":"yezihack"
        },
        "committer":{
            "name":"barry",
            "email":"freeit@126.com",
            "username":"yezihack"
        },
        "added":[

        ],
        "removed":[

        ],
        "modified":[
            "README.md"
        ]
    }
}
`
