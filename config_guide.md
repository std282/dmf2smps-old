# DMF2SMPS configuration files guide

What is _DMF2SMPS configuration file_? It's a file that appears after you run a
command like this:

```
> dmf2smps config song_name.dmf
```

After this command, you'll find a file `song_name.dmf2smps-config`. It contains
a set of _DMF-to-SMPS_ conversion rules in JSON format. One have to configure 
them in order to get appropriate conversion, since this file contains such 
parameters as DAC sample mapping.

## Description of structure

Here you can see an example of dmf2smps configuration.

```json
{
    "file": "song_name.dmf",
    "params": {
        "preferFM6": false,
        "preferPSG3": false,
        "vibratoDecay": false,
        "restartAfterEnd": true,
        "extendedPSG": false,
        "nonDMFArpPorta": false
    },
    "dac": [
        {
            "note": "C",
            "bank": 1,
            "name": "Kick",
            "dacSample": "81"
        },
        {
            "note": "C#",
            "bank": 1,
            "name": "Snare",
            "dacSample": "82"
        },
        {
            "note": "D",
            "bank": 1,
            "name": "FX",
            "dacSample": null
        }
    ],
    "psg": [
        {
            "number": 5,
            "name": "Inst 5",
            "psgEnvelope": 2
        },
        {
            "number": 8,
            "name": "Inst 8",
            "psgEnvelope": null
        }
    ]
}
```

### Global object fields

Field name | Field type | Meaning
-----------|------------|---------
file | string | A path to DMF file to be converted
params | object | Global parameters of conversion
dac | array | Array of DAC sample mapping objects
psg | array | Array of PSG mapping objects

### Parameters of conversion

Description of _params_ object

Field name | Field type | Meaning
-----------|------------|---------
preferFM6 | bool | If _true_, prefer FM6 over DAC in SMPS
preferPSG3 | bool | If _true_, prefer PSG3 tone over noise in SMPS
vibratoDecay | bool | If _true_, decay all vibrato to frequency alteration
restartAfterEnd | bool | If _true_ song will loop at start if it's finished without looping
extendedPSG | bool | If _true_, extra PSG notes will be used for low pitches <br> Look at table in "Extended PSG Notes" chapter.
nonDMFArpPorta | bool | If _true_, arpeggio (`00xy`) and continuous portamento (`01xx`, `02xx`) effects will be able to combine (by default portamento overrides arpeggio)

### DAC mapping object

Objects of this type are contained in _dac_ array.

Field name | Field type | Meaning
-----------|------------|---------
note | string | Name of note in DMF that plays the sample
bank | number | Number of bank in DMF that contains the sample
name | string | Sample name
dacSample | string <br> number <br> null | DAC sample number to map to, either as string (hex representation) or as number (decimal representation) <br> If it's _null_, then the sample is ignored

### PSG mapping object

Objects of this type are contained in _psg_ array.

Field name | Field type | Meaning
-----------|------------|---------
number | number | Number of instrument
name | string | Name of instrument
psg | string <br> number <br> null | PSG envelope number to map to, either as string (hex representation) or as number (decimal representation) <br> If it's _null_, then STD volume envelope is decayed into PSG volume alterations

## Extended PSG notes

By default, Sonic 1 SMPS doesn't cover full range of SN76489 (PSG) frequencies,
making it unable to play lowest B, A# and A notes. By setting field _extendedPSG_
you make DMF2SMPS map those 3 notes to unused SMPS slots.

DMF STD note | SMPS PSG extended note
-------------|------------------------
`A-0` | `$DD`
`A#0` | `$DE`
`B-0` | `$DF`