# Learning & Observations

## Testing

What I like is that we can run tests using standard Go tooling.
No additional setup or configuration is required! This is a very good start
from developer/user experience that would like to use the project/library.  
All tests pass! Great.

```bash
➜  habitfield git:(learning) ✗ go test
Habit added!
Habit added!
Habit added!
PASS
ok   github.com/RyanRalphs/habitfield 0.451s
```

How about making successfull test run be silent (not printing anything to the terminal)? How to achive output like in the example below?

```bash
➜  habitfield git:(learning) ✗ go test
PASS
ok   github.com/RyanRalphs/habitfield 0.451s
```

## Tracking habits

I love indoor and outdoor rowing. I want to use the habit tracker as a logbook for my daily activities. How can I install it on my laptop?
Since the project is in active development, I will try to build a binary myslef and run it on my laptop.

- Building binary:

```bash
➜  habitfield git:(learning) ✗ go build -o habitfield ./cmd/habitfield/main.go
➜  habitfield git:(learning) ✗ ls -l
total 6368
-rw-r--r--  1 jakub  staff      688 25 Apr 08:39 LEARNING.md
-rw-r--r--  1 jakub  staff     1068 25 Apr 08:26 LISENCE
-rw-r--r--  1 jakub  staff     3192 25 Apr 08:26 README.md
drwxr-xr-x  3 jakub  staff       96 25 Apr 08:26 cmd
-rw-r--r--  1 jakub  staff      179 25 Apr 08:26 go.mod
-rw-r--r--  1 jakub  staff     3597 25 Apr 08:26 go.sum
-rw-r--r--  1 jakub  staff     2866 25 Apr 08:26 habit.go
-rw-r--r--  1 jakub  staff     2230 25 Apr 08:26 habit_test.go
-rwxr-xr-x  1 jakub  staff  3198144 25 Apr 08:48 habitfield
-rw-------  1 jakub  staff    65536 25 Apr 08:28 test.db
```

Great, it compiled successfully!

- Tracking habits
What happens if I run habitfield without specifying options or commands?
Let's try.

```bash
➜  habitfield git:(learning) ✗ ./habitfield
Welcome to your personal habit tracker!!

To add a habit, run `habitfield add <habit>`.
To update a habit, run `habitfield update <habit>`.
To list all habits, run `habitfield list`.

2023/04/25 08:50:25 Exiting Program. Please try again after reading the above help message!
```

I like the clear information of what I can do! I am not sure if the log message helps me to achieve my goal. I would like to see only available options.

Let's track  my daily ``rowing``:

```bash
➜  habitfield git:(learning) ✗ ./habitfield add rowing
2023/04/25 09:08:29 add is not a habit command
```

Hmm... am I doing something incorrectly? What would be the best place to get help, ``README`` file, ``--help`` option?

Let's try some options:

```bash
➜  habitfield git:(learning) ✗ ./habitfield --help
2023/04/25 09:10:40 --help is not a habit command
➜  habitfield git:(learning) ✗ ./habitfield -h
2023/04/25 09:10:47 -h is not a habit command
➜  habitfield git:(learning) ✗ ./habitfield help
2023/04/25 09:10:51 help is not a habit command
```

How about to guide users by hand and make them feel satisfied?

- What habits am I tracking?

```bash
➜  habitfield git:(learning) ✗ ./habitfield list
2023/04/25 09:14:56 list is not a habit command
```

- Am I doing something wrong?
- Did I build tracker correctly?
- All tests passed, yet I am confused a bit when I try to use the application?

How the user experience can be improved?

##
