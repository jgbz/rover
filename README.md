
# Rover Explore

## Problem
A squad of robotic rovers are to be landed by NASA on a plateau on Mars. This plateau, which is curiously rectangular, must be navigated by the rovers so that their on-board cameras can get a complete view of the surrounding terrain to send back to Earth. A rover's position and location is represented by a combination of x and y co-ordinates and a letter representing one of the four cardinal compass points. The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North. In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot. 'M' means move forward one grid point, and maintain the same heading. Assume that the square directly North from (x, y) is (x, y+1).


## Installation

This application requires Go to run/build it

visit https://go.dev/doc/install for more information.

## Environment Variables

To run this project with docker, you will need to add the following environment variables to your dockerfile/docker-compose.yaml

`INPUT_PATH`



## Run Locally

Clone the project

```bash
  git clone https://github.com/jgbz/rover
```

Go to the project directory

```bash
  cd rover
```

Install/Update dependencies

```bash
  go mod tidy
```

Run the application with Make

```bash
  make run
```

You can also run with docker compose

```bash
  docker compose up
```
## Running Tests

To run tests, run the following command

```bash
  make test
```

