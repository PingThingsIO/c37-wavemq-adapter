# C37 to WAVEMQ protocol adapter

This application will connect to a C37 device or stream splitter and emit channels from that device on WAVEMQ uris.

To start it, run:

```
./c37wavemq <config.yaml>
```

Here is an example config:


```
# The C37 details
c37TargetAddress: "192.168.50.51"
c37id: 1

# the WAVEMQ details
siteRouter: "127.0.0.1:4516"
namespace: "GyCetklhSNcgsCKVKXxSuCUZP4M80z9NRxU1pwfb2XwGhg=="
entityFile: "entity.ent"

# What to publish
outputs:
  - uri: "upmu/L1"
    channels:
    - "L1MagAng"
    - "FREQ"
  - uri: "upmu/L2"
    channels:
    - "L2MagAng"
    - "DFREQ"
  - uri: "upmu/L3"
    channels:
    - "L3MagAng"
```
