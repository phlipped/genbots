# genbots.bot

A Bot stores the following attributes:

 - energy level
 - brain
 
Notice that the bot doesn't store its current location - the arena is
responsible for tracking this.

A bot is clonable and mutatable. Obviously this clones/mutates the bot's brain as well.

Once a bot has produced an Action, it will wait until the Action has finished
being executed before it continues evaluation.