# CRC-64 Check Sum Calculator.


## Short Description.

The 64-Bit Cyclic Redundancy Check Sum Calculator.

## Full Description.

This is a 64-Bit Cyclic Redundancy Check Sum Calculator.

The 64-Bit Cyclic Redundancy Check Sum Calculation is done using the
Polynomial Table provided by the European Computer Manufacturers Association 
(ECMA).

Receives a space-separated List of File Paths as an Input.
Writes Results to the Standard Output Stream ('stdout').

This Program can process only small Files which can be placed in the RAM.


## Installation.

Import Commands:
```
go get -u "github.com/legacy-vault/tool/crc64ecma/code"
```

## Usage.

Run with '-h' Command Line Argument to list all possible Command Line Arguments.

Command Line Arguments must be provided before any File Path. This is a 
Requirement of the current Version of Golang's built-in 'flag' Package.

Examples:<br />
crc64ecma -uc ./data/a.txt data/b.txt<br />
./crc64ecma data/a.txt ./data/b.txt<br />
crc64ecma /tmp/my_file.dat<br />
