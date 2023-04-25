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

## Code

### Tests

- I like naming convention. It's short and clear. Tests are ``external`` and run against package's public API.
- Nice use of table tests to group multiple scenarios. How about to separate ``positive`` and ``negative`` cases?
- How about make table tests (var ``scenarios``) not a global variable but internal to a test func? What impact could it have on running tests in parallel and avoid possible race conditions?
- Is using two var names ``got`` and ``input`` to keep the same value necessary? How about to use only ``got`` and ``want``? Would it improve readability? Would it decrease or increase cognitive load?
- Do we really care what error string the func returns, or only about the fact that the error occurs? Do we have a business logic that cares about error types? Assuming that the func errors, do we want to fail the test because of business logic problem or because someone (potentally by mistake) changed error string?
- Great approach with using bytes.Buffer for handling fake outputs! This is a nice way of avoiding printing messages from tests.
- vars `ht` and `dbName` are global. What can happen if tests would want to simultanously access for reading / writing? Is data safe from being corrupted? If tests would be executed in parallel how we could avoid race conditions? How about each test having its own obility to setup and tear down pre-conditions?
- Creating pre-conditions in tests - If an error happens during the db setup, how would you inform test runner about it? How about to leverage t.Helper() func from the testing package? Where would you benefit from using ``t.Fatal`` method?
- I like the fact that you re-create db for each test. Does the syntax ``defer setupTest()()`` show clearly intent for a reader? How to approach cleaning resources after each test with functions / methods from the ``testing`` pkg?
- I really like your short error messages in tests. Using ``want`` and ``got`` make code clear and nicely shows the intent.

### Code

- NewTracker depends on a particular DB storage. How about to abstract store, for example if some user of your package would like to use MySQL?
- It shapes really nicely! All tests pass. I like the functionality, how can I make it work on my machine?
