# Cheatsheet Navigator (CSN)

It's just a fancy find and grep, implemented in Go. It's mostly a personal tool
also use to learn more about Go plus the following learnings goals.

## Learning Goals

1. Implementation of a simple CLI app without 3party packages.
2. Cross platform release using goreleaser.
3. Publish on homebrew and maybe other linux package manager.
4. Create zsh autocomplete for a modern CLI experience.

## Why?

Well I tried other tools like navi, cheat and I wanted something simpler
I just want to very quickly access my notes or parts of them from the terminal.

For example I can use with my Obsidian Vault and every quickly read my note in
markdown or even filter just parts that I want.

## Dependencies

## How to

## Stuff I learned

### ZSH Command Completion on Tab

Create a function that generates the possible completion options

```bash
function _myapp_completion {
    # Use ls to get the list of files/directories in the current directory
    local files=($(ls))
    _describe 'values' files
}
```

After that on `.zshrc` add the completion

```bash
compdef _myapp_completion myapp
```

### Homebrew Tap Repository

1. Create a repository in this case `fabiodcorreia/homebrew-repo`
2. `brew tap fabiodcorreia/package`
