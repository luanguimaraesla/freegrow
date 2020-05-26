Freegrow
========

Freegrow is an opensource controller for hand-made greenhouses.

The idea behind freegrow is to create an open project of a DIY greenhouse based on low-cost materials, motors and sensors controlled by a single microprocessor (e.g Raspberry Pi).

## DIY Greenhouse Manuals

The following links describe how to build your own physical machine. They have a list of materials, motors, sensors, lights, and the steps to make everything work together. Although this project is supposed to be high-adaptable, some controllers were created to work for specific models/structures, and may not work for some weird use-cases.

* [Physical structure](docs/PHYSICAL_STRUCTURE.md)
* [Hardware Schema](docs/HARDWARE_SCHEMA.md)

## System Overview

Freegrow is a distributed system based on three main parts, the `storage`, the `machine`, and one or more `nodes`. These three pieces together create a fault-tolerant and customizable system which is basically an events manager that controlls the whole process of planting something, activating and deactivating valves, lights, and motors based on sensors data, and some simple configuration manifests.

![System Overview](docs/images/overview.jpg)

### Storage

Freegrow uses an Etcd cluster to store all resources it needs to handle the...


### Install

#### Requirements

- Go >= 1.14
- Make

#### Building and Installing

To install freegrow you still need to build it from scratch. Just clone this repository and run make:

```bash
# clone this repository
git clone https://github.com/luanguimaraesla/freegrow.git

# enter the project directory
cd freegrow

# download dependencies, build and install the golang program
make

# test if freegrow is working for you
freegrow version
```

### Running

Freegrow needs a simple manifest to configure its internal resources. We call these resources `Gadgets`, for which we have to define some custom settings in order to specify their behavior. Irrigator is an example of a Gadget, we can control the time it should start and stop using a simple crontab notation. Also we need to define the board digital port where it is connected.

```yaml
board: fakeboard
gadgets:
- class: irrigator
  spec:
    name: default
    port: 14
    states:
    - name: "on"
      schedule: "5 9 * * *"     # starts at 9:05AM
    - name: "off"
      schedule: "10 9 * * *"    # finishes at 9:10AM
```

Freegrow supports some different board configurations. At this moment, we have two different backends `fakeboard` and `raspberry`.

##### FakeBoard

In order to run development tests and some proof of concepts, we created the FakeBoard controller, which mocks GPIO enabling you to run the whole system inside your machine.

```bash
# running freegrow outside a board
freegrow start -f examples/irrigator.fakeboard.yaml --log debug
```

##### Raspberry Pi 3

Connect your Irrigator device to the raspberry digital pin 14 and run the following command:

```bash
# running freegrow inside a raspberry
freegrow start -f examples/irrigator.raspberry.yaml --log debug
```
