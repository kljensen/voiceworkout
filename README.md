# Voice workout helper

This code creates a computer voice that walks you through
a workout, similar to my shell interval trainer. In this
case, there is some interaction with the program due to the
inclusion of exercises that are not timed, but rather a
certain number of reps, which requires the user to touch
a key to continue the program when the reps are done.

The program is run like

```
voiceworkout routine.yaml
```

where `routine.yaml` has contents something like the following:

```yaml
title: Example routine
exercises:
  - instruction: band dislocations
    number: 10
    istimed: true
  - instruction: pushups
    number: 50
    istimed: false
```

A more thorough routine is
[https://gist.github.com/kljensen/49ecc10a1440d3efc1043eabf57463b4](in this gist).
That will yield output similar to the following

```
voiceworkout v1.1
Now doing band dislocations. Do band dislocations for 10 seconds
Done with band dislocations
Now doing pushups. Do pushups for 50 repps and hit any key when finished
Done with pushups
```

## Installation 

You can build this from source (it is written in Go) or download the
[latest release of voiceworkout](https://github.com/kljensen/voiceworkout/releases).
It does not compile on windows right now because of using the
[https://github.com/pkg/term](term) package.

Or, if you are a [homebrew](https://brew.sh/) user, you can install
`voiceworkout` via the [kljensen/tap](https://github.com/kljensen/homebrew-tap)
tap using either

```
brew install kljensen/tap/voiceworkout
```

or

```
brew tap kljensen/tap
brew install voiceworkout
```

