# Event sourcing (DRAFT)

An event is something that happened in the **PAST**.
Do not confuse events with commands.
A command is an order you want to be executed in the **FUTURE**.

It should be named with a past tensed verb.

## Naming

* pictureScanned => is a good name for an event
* pictureScan => wrong, that is a command
* picturePut => could be both

## Within a single service, leveraging DB

Sample

## Between services, leveraging an event stream

NATS jetstream, Kafka
