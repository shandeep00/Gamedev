Problem statement: In online multiplayer games, there are four kind of use cases based on client connectivity to the server.

•	Client is always connected to the server.

•	Client is disconnected but reconnects to the server in some time.

•	Client is completely disconnected from the server.

•	Client is always connected to the server but has low bandwidth.

We need to design a solution at server level which sends appropriate information to the client based on the use case. This solution should be a common one which can be used for different kinds of games. 

The solution should handle live streaming of game and should be highly scalable (Game should not be affected by the number of viewers watching the game).

Case 1: Client is always connected to the server

•	This is an ideal scenario and server can send all the information about the current state of the game to the client. 

•	Server shows the client connectivity information (connected or disconnected) to all other clients.

Case 2: Client is disconnected but reconnects to the server in some time.

•	In this case, server can send previous state information from when client disconnected to the point client reconnected, along with current state of the game. 

Case 3: Client is completely disconnected from the server.

•	No information need to be passed to the client.

Case 4: Client is always connected to the server but has low bandwidth.

•	Here, data is not updated frequently or data is not frequently sent from the server to the client

•	So, it is not required to show every state of the game after every action to the client, instead we can skip showing some states.

•	States like where other user is performing action can be skipped.

•	Questions
1.	How we check status of the clients (connectivity, network strength (low bandwidth or not)?

Entities 

•	Gameserver

•	Client or player

•	server

Streaming live game:

•	If we are streaming synchronously, where there is no delay, then game server needs to be sequentially sending game data to the live stream server. This may affect the performance of the game server when viewers increase.

•	Instead we can stream the game with delay depending on the hardware of the server. In this case, we send data to the live stream server after certain event count/after some time limit is reached.

•	We can use a queue to maintain the order of the events.