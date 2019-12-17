# Let's Breakup a Mono Repo

This is where we actually break up this repo into smaller segments.

## Before we begin

We'll be using go module functionality alongside the GOPATH so ensure that you have GO111MODULE set to on.

This can be set in your current shell by running the following:

    export GO111MODULE=on

The following bash function will be used throughout this tutorial:

    # shellcheck shell=bash

    function rsed() {
        if [ "$#" -ne 3 ]; then
            echo "usage: rsed <path> <search string> <replacement>"
            return 1
        fi

        case $(uname) in
            Linux)
                find "$1" -name .git -prune -o -type f -exec sed -i "s|$2|$3|g" {} +
                ;;
            Darwin)
                find "$1" -name .git -prune -o -type f -exec sed -i '' "s|$2|$3|g" {} +
                ;;
            *)
                echo 'unsupported operating system'
                return 1
        esac
    }

## Break out library_a

Start by copying your mono repo to a new version of itself before proceeding (note we do a checkout because otherwise Darwin machines may fail).

    cp -r letstalkaboutmonorepos letstalkaboutmonorepos_library_a
    cd letstalkaboutmonorepos_library_a
    git checkout master

Run a subdirectory-filter to rewrite the repo history so that it includes only our library code with no nested directories.

    git filter-branch --prune-empty --subdirectory-filter src/library_a @

Rewrite our test imports (we use the rsed func from above because, if we had more than one file this would be faster than manually doing a find and replace).

    rsed . github.com/ingcr3at1on/letstalkaboutmonorepos/src/library_a github.com/ingcr3at1on/letstalkaboutmonorepos_library_a

initialize a go module.

    go mod init github.com/ingcr3at1on/letstalkaboutmonorepos_library_a

And ensure the tests run.

    go test

You should see:

    PASS
    ok  	github.com/ingcr3at1on/letstalkaboutmonorepos_library_a	0.003s

Finally commit your changes:

    git add go.mod
    git commit -a -m 'Extract librarya from mono repo.'

## Break out the service.

Start with the cmd directory.

    cd ../
    cp -r letstalkaboutmonorepos service_cmd
    cd service_cmd
    git checkout master

We start by running a subdirectory-filter on the cmd directory (we do this in phases for scale, if we try to use a tree filter to do this all at once it would work fine in this case but could take several days on larger repos depending on the machine resources).

    git filter-branch --prune-empty --subdirectory-filter cmd/service @

This will rewrite the git history removing all but the files in cmd/service (in our example this will leave nothing but a main.go in the directory).

Now that we have reduced the size of our tree we can use a tree-filter to rewrite it without risking having to wait a super large amount of time for it to run.

    git filter-branch --prune-empty --force --tree-filter \
    'mkdir cmd && git mv -k * cmd' @

The `--force` flag is required because we already ran the git filter-branch command on this repo and it will fail otherwise.

We use the `-k` flag for our mv command to ensure that if something doesn't exist in a specific commit it's simply ignored while still being rewritten in other commits that it does exist in.

Next we want our service code.

    cd ../
    cp -r letstalkaboutmonorepos service_src
    cd service_src
    git checkout master

Like before we're going to collapse this down to just the files we care about:

    git filter-branch --prune-empty --subdirectory-filter src/service @

Finally put it together in a new final repo:

    cd ../
    mkdir letstalkaboutmonorepos_service
    cd letstalkaboutmonorepos_service
    git init .

You need _something_ in the repo history before you can proceed so we'll just commit the empty tree.

    touch .gitkeep
    git add .
    git commit -a -m 'Initial empty repo'

Use local remotes to merge in our 2 temporary repos (to my knowledge this only works in *nix systems):

    git remote add cmd ../service_cmd
    git remote add src ../service_src
    git fetch cmd
    git merge cmd/master --allow-unrelated-histories
    git fetch src
    git merge src/master --allow-unrelated-histories
    git remote rm cmd
    git remote rm src

Now fix the build and tests:

    rsed . github.com/ingcr3at1on/letstalkaboutmonorepos/src/library_a github.com/ingcr3at1on/letstalkaboutmonorepos_library_a
    rsed . github.com/ingcr3at1on/letstalkaboutmonorepos/src/service github.com/ingcr3at1on/letstalkaboutmonorepos_service
    go mod init github.com/ingcr3at1on/letstalkaboutmonorepos_service

We haven't pushed library_a anywhere (this is why we are working in the GOPATH), run our tests without modules:

    GO111MODULE=off go test ./...

Assuming everything passes, edit the README.md however you like, consider pushing library_a to a location and finally commit your changes (I won't cover these last few steps here, this is part of why I suggested knowing git basics), lastly don't forget to clean up your extra repos.

## Final thoughts

This was a process I worked out months ago and never took the time to fully write down in this manner, it has been done entirely from memory to avoid referencing work materials so it's possible I have missed a tactic or two with regards to ensuring the final repos are as minimal in size as possible. If you have anything to add please open an issue on this repo.
