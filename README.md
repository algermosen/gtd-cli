# GTD on your TERMINAL

There are so many options to organize your tasks. On top of that a bunch o methods to work with. This is just to past the time, but I always wanted to create a gtd-specific to-do appliacation. Despite I would rather to do it on mobile, I find cli quite exquisite to work, so this is it.

## Installation

TBD

## Commads

Idk if I'm gonna use this at the end, but I'll just take it as a way to barinstorm ideas of how someone should interact with the cli.

```gtd```: This is the core, the process command. I chose this one since is a simple and short way to call the action most performed (process your inbox). It need sto be interactive and simple to follow. The process consists in having an overview of the action list while looping through inbox and someday list, use specific keys (TBD) to perform the task movement among lists.

```gtd add "new task"```: This will add a task to the inbox.
```gtd add "new task" -t (action, someday, tickler)```: This will add a task to a specific type (just in case)

```gtd check (action, someday, tickler)```: This will show the list and allow to check tasks from it by using vim motions ('j' for going down, 'k' for going up) or the arrow keys to move through the tasks, and enter or space to check or uncheck the list. This command will show just the unchecked tasks.
```gtd check (action, someday, tickler) -a```: This command will show all tasks of a list.

```gtd clean (action, someday, tickler)```: This will delete the checked tasks of a list. Might ask you to confirm.
```gtd dump (action, someday, tickler)```: This will delete all the tasks of a list. Might ask you to confirm by writing a word.

## Interactive mode

My thoughts of how the interactive modes should look and operate.

First, I though of using bubbletea and lipglose, but for the sake of KISS and actually writting functional behavior and not just 'Good Looking Stuffs (GLS)' I'm gonna use nothing but the common CLI.

Either way, dealing with a list of tasks just by typing keys, specially when you want to select or check a task, might be overwhelming. So, I can think of using numeric codes.

Think of showing the list of items to check:

```
[ ] Do the laundry
[ ] Do the dishes
[ ] Do sport
```

Then, allow the user to select them by typing the number of tasks to change:

```
1- [ ] Do the laundry
2- [ ] Do the dishes
3- [ ] Do sport
Items to check (separate with spaces)
>>> 1 3
```

Maybe it's not the optimal and better outcome, but I do preffer learn to implement something simple like this rather than learning how bubbletea and lipglose works, at least for now. Gonna focus with cobre and having something finished for once.

## Storage

I always get dizzy thinking about storage. Should implement a simple In memory Storage for developing and after deal with database? In the projects I'm not planning to actually share or be used, I don't like to having a big codebase of queries, But neither I want to make the things too specific, I want to be able to change add a database if I want to without having to change all the code.

That's when I get dizzy, I know I need some kind of Dependency Injection .....

What I need to store?

- Name of tasks
- Date when done
- Date when created
- List type / Task type (inbox, action, someday, tickler)

## Tasks

- [ ] File Store implementation
- [ ] File Store Seed Data
- [ ] Add new tasks
- [ ] Check Command (List tasks)
- [ ] Check Command (Check tasks)