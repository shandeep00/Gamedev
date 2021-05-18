package Entities

type server struct {
	clients map[int]*Client //All clients connected to server, clientId->client
	games map[int]*Gameserver //All created games gameId->Gameserver
	client_game map[int]int //client game relation clientId->gameId
	disconnected map[int]int //disconnected and reconnected clients clientId->time
	gameId int //unique gameid
	playerId int //unique playerid
}

type Server struct {
	*server
}

func (s *server) AddClient(name string) *Client {
	c := NewClient(s.playerId, name, 1)
	s.clients[s.playerId] = c
	s.playerId++
	return c
}

func (s *server) UpdateClient(clientId int, status int, time int) {
	s.clients[clientId].status = status
	if status == 2 {
		s.disconnected[clientId] = time
	} else if status == 3 {
		delete(s.clients, clientId)
		gameId := s.client_game[clientId]
		s.games[gameId].removePlayer(clientId)
		delete(s.client_game, clientId)
	}
}

func (s *server) ConnectGame(clientId int, gameid int) {
	s.client_game[clientId] = gameid
	s.games[gameid].addPlayer(clientId)
}

func (s *server) AddState(gameId int, state string) {
	s.games[gameId].AddGamestate(state)
}

func (s *server) SendState(clientId int, time int) string {
	if s.clients[clientId].status == 2 {
		return "disconnected"
	}

	gameId := s.client_game[clientId]
	endgame, curr_state := s.games[gameId].GetGamestate(time)
	if s.clients[clientId].status == 4 && !s.games[gameId].isPriorityState(clientId, curr_state) {
		return ""
	}

	if index, ok := s.disconnected[clientId]; ok {
		for i := index; i < time; i++ {
			_, prev_state := s.games[s.client_game[clientId]].GetGamestate(i)
			curr_state += "-" + prev_state
		}
		delete(s.disconnected, clientId)
	}

	if endgame {
		delete(s.games, gameId)
	}
	return curr_state
}

func (s *server) CreateGame(name string, capacity int, endState int) *Gameserver {
	gs := NewGameserver(s.gameId, name, capacity, endState)
	s.games[s.gameId] = gs
	s.gameId++
	return gs
}

func (s *server) GetClients(gameId int) []int {
	return s.games[gameId].players
}

func New() *Server {
	s := &server{
		clients: make(map[int]*Client, 0),
		games: make(map[int]*Gameserver, 0),
		client_game: make(map[int]int, 0),
		disconnected: make(map[int]int, 0),
		gameId: 1,
		playerId: 1,
	}

	S := &Server{s}
	return S
}