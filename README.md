# genbots

# design notes

## Bot Evaluation
- bots brains are evaluated concurrently
- one of the Controller's jobs is to check for an Action from each of the bots
- bots will wait until the Controller has collected the Action before continuing with the next evaluation.

### Managing over-thinking
- Occasionally, a bot may try to over think a problem.

## Food and Smells

### Option 1
- Food gives off a smell. The strength of the smell it gives off is proportional to the distance.
- The total amount of 'smell' a bot recieves at a particular location is the sum of the (proportional) smells it has access to
- Smells fade away when the source of the smell is no longer present.
- Bots can drop smells

### Option 1b
- There are different types of smells.
- There is no crossover between smell types.
- A bot has to sniff for a particular smell.

### Option 2
- Food is directly 'pollable' - i.e. bots can check if there is food in their vicinity