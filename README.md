# mt

`mt` is a minimalistic todo CLI. It helps you manage tasks directly from your
code or Markdown files. The project focuses on simplicity and a lean, no-frills
approach to task management.

It supports the same basic task management features as [Obsidian Task Management](https://publish.obsidian.md/tasks/Getting+Started/Introduction),
but it is a standalone tool that can be used in the terminal. The goal is be interoperable with the Task Management plugin for Obsidian on defining
dates, priority and recurring tasks. We also want to be able to use the same query concept as in Task Management plugin.

The goal is that you can configure the global note directory to be the same as the Obsidian vault and use the same tasks in both tools.

This is currently being refactored into Zig, the original (simple proof of concept) version was written in Golang.
I like Golang, but since I use Go for work, I want to learn a new language.

Discuss and follow the development in [the issue #5](https://github.com/simenandre/mt/issues/5).

## Usage

```sh
mt [command]
```

If no command is provided, `mt` will list all tasks in the current directory or the
global note directory if one is configured. If you have configured a global note
directory, you can also specify a different directory to list tasks from:

```sh
mt -d ~/notes
```

### Add a task

```sh
mt add "Buy milk"
```

Assuming you have configured a global note directory, this command will add a
new task to the daily note for today. If not, it will add it to `TODO.md` in the
current directory. You can also specify a different file to add the task to:

```sh
mt add "Buy milk" -f ~/notes/TODO.md
```

You can also add a due date to the task:

```sh
mt add "Buy milk" -d 2021-12-31
```

The output of this command will be in the format of:

```markdown
- [ ] Buy milk 📅 2021-04-09
```

You can also add a priority to the task:

```sh
mt add "Buy milk" -p 1
```

The priorites can be from 1 to 6, with 1 being the highest priority. You can also
write lowest (6), low (5), none (4), medium (3), high (2) or highest (1) instead.

The output of this command will be in the format of:

```markdown
- [ ] Buy milk 📅 2021-04-09 🔺
```

### Editing a task (i.e. completing it)

I want an interactive mode for editing, but also something that works in the terminal (and can be scripted).

It might be something like:

```sh
mt complete 1
```

### List tasks

```sh
mt list
```

Or, just `mt` with no command. You can specify a different directly to list tasks from:

```sh
mt list -d ~/notes
```

Or `mt -d ~/notes` with no command. We also have a short version of the command:

```sh
mt l
```

To list all tasks in the current directory, you can use:

```sh
mt lc
```

## Contributing

Contributions are welcome! Please open an issue or a pull request if you have any
ideas or suggestions. If you are interested in writing some code, we would love
if you could take a look at the open issues and see if there is anything you would
like to work on.

`mt` tries to embrace simplicity and minimalism, so please keep that in mind when
contributing. We want to keep the codebase as small, simple and accessible as possible.
This means we focus on direct, readable code. We aim to keep the scope of the project
to a minimum, so we can keep the codebase small and maintainable. Each part of the
code should work correctly, be well-tested (end to end is prefered), with no
technical debt.

We try to avoid introducing unnecessary abstractions and keep things straightforward,
we focus on clear, understandable logic over clever tricks. We embrace the _negative
space programming_ concept, so we try to use `assert` liberally to document and
enforce assumptions about the code.

`mt` is built to be a small, sharp tool that does its job well and stays out of your way.
We hope you want to help us with that and have fun doing it!

## License

Apache 2.0
