# genbots.controller

genbots.controller is the central coordination point that connects all
the other objects.

It is the main execution loop that ties everything together:
 - consumes Actions from each Bot's output channel and implements the necessary
   outcome (e.g. updating the Arena for position change, or Cloning/Mutating the bot)
   as well as decrementing the appropriate amount of energy.
 - Manages the energy levels of bots and kills them as necessary
   (each bot will have it's energy levels slowly deplete regardless of the actions
   they take (except if they are sleeping))
The controller 