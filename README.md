# ccfw

> [!IMPORTANT]
> This project is in early development. Expect breaking changes.

`ccfw` is a **C**laude **C**ode **F**rame**W**ork for reducing maintenance costs.

The `ccfw` CLI manages Claude Code configuration files using `.ccfw/settings.json` as the single source of truth! ✨

<div style="display: flex; align-items: center; gap: 32px">

```json
{
  "version": 0,
  "agents": {
    "designer": {},
    "implementer": {},
    "reviewer": {}
  },
  "options": {
    "agents": {
      "commandPrefix": "x-",
      "invocationMode": "command"
    }
  }
}
```

➡️

```text
.
└── .claude/
    ├── agents/
    │   ├── designer.md
    │   ├── implementer.md
    │   └── reviewer.md
    └── commands/
        ├── x-designer.md
        ├── x-implementer.md
        └── x-reviewer.md
```

</div>

## Policy

**"Don't train, define rules."**

Humans only maintain "AI-assisted workflows" and "Rules for the AI", while the AI itself maintains "slash-commands", "sub-agents", and "documents" based on these inputs.
