# Things a shell does

    _Execute scripts in PATH in foreground or background
    _Give error message if input isn't intelligible
    _Execute different functions based on special characters
        _$   - command substitution
        _' ' - literal evaluation
        _\   - escape character
        _>   - pipe output to something (create/overwrite file)
        _>>  - pipe output to something (append/overwrite file)
        _<   - pipe input to something
        _|   - pipe output to command's input
        _&   - execute in background
        _&&  - execute command and command
        _!!  - last command
        _*,? - globs
    _Interrupts
        _ctrl+c - kill
        _ctrl+z - suspend

# Things a shell has

    _History
    _Environment Variables (PATH
    _Prompt
    _Input buffer
    _Output buffer
    _Builtin functions
        _cd
        _exit
    _directory location

# Design

    _History file
    _Possible customization via .shellrc
    _Builtin functions
    _Input parsing

# Feature Ranking

    _Builtin functions cd + exit
    _History logging
    _Input Parsing
    _Foreground and Background execution (thread safety warning?)

