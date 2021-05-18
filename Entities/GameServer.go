package Entities

import "strconv"

type gameserver struct {
	gameId int
	name string
	capacity int
	players []int //in game players
	states []string
	endState int //where game ends
}

type Gameserver struct {
	*gameserver
}

func (gs *gameserver) GetGamestate(time int) (bool, string) {
	if time == gs.endState {
		return true, gs.states[time-1]
	}
	return false, gs.states[time-1]
}

func (gs *gameserver) AddGamestate(input string) {
	gs.states = append(gs.states, gs.updateState(input))
}

func (gs *gameserver) updateState(input string) string {
	//update state based on client input
	return "STATE"+input
}

func (gs *gameserver) addPlayer(clientId int) {
	if gs.isGameFull() {
		//error
	}

	gs.players = append(gs.players, clientId)
	if gs.isGameFull() {
		gs.startGame()
	}
}

func (gs *gameserver) removePlayer(clientId int) {
	j := 0
	for _, n := range gs.players {
		if n != clientId {
			gs.players[j] = n
			j++
		}
	}
	gs.players = gs.players[:j]
}

func (gs *gameserver) isGameFull() bool {
	return len(gs.players) == gs.capacity
}

func (gs *gameserver) isPriorityState(clientId int, state string) bool {
	//process the state for this client and see if it is required to send this state
	sub := state[5:]
	i, _ :=  strconv.Atoi(sub)
	return i % 2 == 0
}

func (gs *gameserver) startGame() {
	//start game
}

func newGameserver(gameId int, name string, capacity int, endState int, players []int, states []string) *Gameserver {
	gs := &gameserver{
		gameId: gameId,
		name: name,
		capacity: capacity,
		players: players,
		states: states,
		endState: endState,
	}
	GS := &Gameserver{gs}
	return GS
}

func NewGameserver(gameId int, name string, capacity int, endState int) *Gameserver {
	players := make([]int, 0)
	states := make([]string, 0)
	return newGameserver(gameId, name, capacity, endState, players, states)
}