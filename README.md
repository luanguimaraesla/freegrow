Freegrow
========

Freegrow is an opensource controller for hand-made greenhouses.

The idea behind freegrow is to create an open project of a DIY greenhouse using low-cost materials and a Raspberry Pi.

### Physical structure

The greenhouse's physical structure project will be available as soon as we finalize its first version.

### Hardware

The Raspberry Pi schema, sensors, motors, and so on, will be available as soon as we finalize its first version.

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
