# go-learner

🐹 Learning Go one exercise at a time — practical mini-projects and challenges covering fundamentals, concurrency, and idiomatic Go patterns.

This repo is a structured curriculum for learning Go through hands-on projects, ordered from basics to advanced topics and design patterns.

## How this repo is organized

Each project lives in its own top-level folder, numbered roughly in learning order:

```
01-<project-name>/
  README.md       # topics covered, setup steps, goals for this project
  go.mod
  *.go
02-<project-name>/
  ...
```

Every project folder has its own `README.md` describing:
- **Topics covered** — the specific Go concepts or patterns the project teaches
- **Setup** — how to init/run/test that project (`go mod init`, `go run`, `go test`, etc.)
- **Goals** — what you should be able to do/explain by the end of it

Start at `01-` and work through them in order — later projects assume you're comfortable with earlier concepts.

## Ground rules for the intern

**Turn off AI autocomplete (Copilot, Cursor AI, etc.) while working through these projects.** The point is to build muscle memory for Go syntax, the standard library, and how to debug your own mistakes. Autocomplete short-circuits that — you can turn it back on for real work afterward.

Other things to actually do:
- **Write the code yourself first, then compare** — don't look at a solution/reference before attempting the exercise.
- **Run `gofmt -l .` (or just save with gofmt-on-save) and `go vet ./...` before considering anything done.** Idiomatic formatting isn't optional in Go — it's the convention, not a style choice.
- **Read the compiler/vet errors fully.** Go's errors are usually precise about what's wrong and where.
- **Use `go doc <package>` or pkg.go.dev instead of guessing** at standard library behavior.
- **Write a test for anything with a branch or a loop**, even a small one — `go test ./...` should pass before you move to the next project.
- **Ask "why" questions, not just "how"** — e.g. why a channel deadlocks, why a slice re-slice shares backing arrays, not just how to fix the symptom.
- **Commit each finished project separately** so progress and mistakes are visible in history — don't squash your learning process away.

## Getting started

```bash
cd 01-<first-project>
go mod init <module-name>   # if not already done
go run .
go test ./...
```
