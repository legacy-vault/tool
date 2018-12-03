# MD5 Check.


## Short Description.

An MD5 File Checksum Verification Tool.

## Full Description.

A Tool to verify Checksum Files created by 'md5sum' Utility.
Compares the Checksums in the File with the real Checksums.
Is used to verify the Integrity of the Files.
To get reliable Results, the Checksum File must be valid, 
i.e. it must contain good (not damaged) Check Sums.

This Tool uses a simple classic programming Approach: 
the Application is single-threaded and uses consecutive Read and Check.
Files are read into Memory (RAM), so be carefull with very large Files.

## Installation.

Import Commands:
```
go get -u "github.com/legacy-vault/tool/md5check"
```

## Usage.

Run with '-h' Command Line Argument to list all possible Command Line Arguments.
