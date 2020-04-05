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
