# genbots.bot

A Bot stores the following attributes:

 - energy level
 - brain
 - memory
 - colour

 Notice that the bot doesn't store its current location - the arena is
 responsible for tracking this.

 A bot is clonable and mutatable. This clones/mutates the bot's brain as well.

 Once a bot has produced an Action, it will wait until the Controller has consumed
 and implemented that action before it begins its next evaluation