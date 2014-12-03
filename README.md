
To Build Spnee Run the following commands:

```
./run/buildonce
./run/buildspnee
```

The **buildonce** script only needs to be run once :)

And any time you add new go packages or delete old ones.

Depending on the APIs you are testing different scripts and simulators
need to be run...

The rules API can be tested via the rules simulator.

All other APIs can be tested by running

```
./run/buildtest
```
