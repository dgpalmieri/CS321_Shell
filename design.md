# Things a shell does

* Execute scripts in PATH in foreground or background
* Give error message if input isn't intelligible
* Execute different functions based on special characters
    * $   - command substitution
    * ' ' - literal evaluation
    * \   - escape character
    * \>   - pipe output to something (create/overwrite file)
    * \>>  - pipe output to something (append/overwrite file)
    * <   - pipe input to something
    * |   - pipe output to command's input
    * &   - execute in background
    * &&  - execute command and command
    * !!  - last command
    * \*,? - globs
    * Interrupts
        * ctrl+c - kill
        * ctrl+z - suspend

# Things a shell has

* History
* Environment Variables (PATH
* Prompt
* Input buffer
* Output buffer
* Builtin functions
    * cd
    * exit
* directory location

# Design

* History file
* Possible customization via .shellrc
* Builtin functions
* Input parsing

# Feature Ranking

- [ ] Builtin functions
    - [ ] cd
    - [x] exit
- [x] History logging
- [ ] Input Parsing
- [x] Foreground and Background execution (thread safety warning?)

