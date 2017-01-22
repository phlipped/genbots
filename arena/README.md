# genbots.arena

The arena represents the physical space where the action plays out.

The Arena's job is to track spatial information about entities, and allow
this information to be queried.

Queries can be made against particular entites, or against areas of the Arena.

The Arena provides a mechanism for other processes to 'subscribe' to events.
This is anticipated as teh primary mechanism used by a GUI to 
