# goShell (Better name soon)

goShell is, believe it or not, a shell implemented in golang. Currently very under development, this was created as my class project for CS321 Operating Systems.

### Current limitations

* goShell does not support autocompletions
* goShell only supports the following special characters, and only one type of argument per command:
    * |  - pipe output to command's input (only up to 4 pipes supported, arbitrary number soon)
    * && - execute command one and command two (arbitrary number supported)
    * &  - execute command in background (must be last character in command)

