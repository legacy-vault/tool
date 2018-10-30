# Torrent File Renamer.


## Short Description.

This Application renames BitTorrent Files.

## Full Description.

This Application renames BitTorrent Files.
It searches the Directory (with or without Sub-Directories), 
decodes all the '.torrent' Files, calculates BTIH for each One of them,
and renames each File, setting a BTIH Hash Sum as a new Name.
All Output Files are put in a flat Manner (without Sub-Folders).<br />
<br />
Before doing any write Operations, the Tool checks Syntax of all Files (which 
it finds) and checks them for Hash Sum Collisions, or, better to say, 
Duplicates, as real Collisions are very rare Things to see.<br />
<br />
If this Tool finds two or more Files with same BTIH, it offers to view a 
detailed List of Duplicates. End User may approve or abort the Procession of 
Files. In most Cases, BTIH Sum Duplicates appear either when you have the same 
File with different Names or you have a modified BitTorrent File with the same 
'info' Section but with Differences in other Sections (some Torrent Trackers 
like to modify your Content). The Probability of real BTIH Collisions is very 
low.

## Installation.

Import Commands:
```
go get -u "github.com/legacy-vault/tool/torrent-renamer"
```

## Usage.

Run with '-h' Command Line Argument to list all possible Command Line Arguments.
