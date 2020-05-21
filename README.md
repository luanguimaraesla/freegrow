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

Freegrow supports some different board configurations. At this moment, we have two different backends `fakeboard` and `raspberry`.

##### FakeBoard

In order to run development tests and some proof of concepts, we created the FakeBoard controller, which mocks GPIO enabling you to run the whole system inside your machine.

```bash
# running freegrow outside a board
freegrow start --board fakeboard --log debug
```

##### Raspberry Pi 3

We don't have a image docker or something similar to facilitate the deployment of freegrow within your board. In order to ship this code to your raspberry you need to configure golang and install freegrow using make. After doing this, you might run:

```bash
# running freegrow inside a raspberry
freegrow start --board raspberry
```
