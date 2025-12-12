# ccfw

> [!IMPORTANT]
> This project is in early development. Expect breaking changes.

`ccfw` is a **C**laude **C**ode **F**rame**W**ork for reducing maintenance costs.

The `ccfw` CLI manages Claude Code configuration files using `.ccfw/settings.json` as the single source of truth! ✨

```json
{
  "version": 0,
  "agents": {
    "designer": {},
    "implementer": {
      "commandPrefix": "y-"
    },
    "reviewer": {
      "invocationMode": "auto"
    }
  },
  "options": {
    "agents": {
      "commandPrefix": "x-",
      "invocationMode": "command"
    }
  }
}
```

⬇️

```text
.
└── .claude/
    ├── agents/
    │   ├── designer.md
    │   ├── implementer.md
    │   └── reviewer.md
    └── commands/
        ├── x-designer.md
        └── y-implementer.md
```

## Install

Go

```sh
go install github.com/2754github/ccfw/cmd/ccfw@latest
```

cURL

```sh
curl -sSL "https://github.com/2754github/ccfw/releases/latest/download/ccfw_Darwin_arm64.tar.gz" -o ccfw.tar.gz && tar -xzf ccfw.tar.gz && rm ccfw.tar.gz
```

> [!NOTE]
> This example is for macOS (Apple Silicon). For details, please refer to [Releases](https://github.com/2754github/ccfw/releases).

## Quick Start

To try it out:

```sh
ccfw
```

---

You can start by generating a sample `.ccfw/settings.json`:

```sh
ccfw init
```

After customizing `.ccfw/settings.json`, you can manage Claude Code configuration files:

```sh
ccfw sync
```

## Policy

**"Don't train, define rules."**

Humans only maintain "AI-assisted workflows" and "Rules for the AI", while the AI itself maintains "sub-agents", "slash-commands", and "documents" based on these inputs.
