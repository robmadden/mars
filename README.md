# PROBLEM ONE: MARS ROVERS

## Instructions

To run the Mars simulation, clone this repository and execute `make run` on the console. Ensure that you have created an "instructions.txt" file. Mars uses "instructions.txt" to run the simulation. Formatting for the instructions is as was given in the initial instruction set. For Example:

```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

If the instructions are not formatted properly by the user, Mars will have unexpected results. Mars is not built to handle input gracefully, it expects a certain format as specified and will exit if it receives instructions in an otherwise specified format.

## Behavioral Rules

1.  A Rover cannot "run into" another Rover.
2. A Rover should never fall off an edge of the plateau.
3. Bad instructions (telling a Rover to move off an edge or run into another Rover) are ignored and may affect expected behavior/output.
