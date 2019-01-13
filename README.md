# dmf2smps

Converter from DefleMask Format to Simple Music Playback System (Sonic 1)

## TODO
- [x] implement DMF parsing (package *dmfparse*)
- [x] implement SMPS assembling (package *smpsbuild*)
- [x] tidy up *dmfparse* and *smpsbuild*
- [ ] implement conversion algorithm
- [x] implement configuration file creation and analysis
- [ ] implement shell commands analyzer
- [ ] write proper README.md

## Current state

Right now I am trying to figure out how to actually convert DMF events into SMPS
events. This is going to take a while, since there are a lot of different bits
and pieces that are needed to be taken into account.