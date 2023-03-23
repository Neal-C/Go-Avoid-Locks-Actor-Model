//lint:file-ignore ST1006 because

package server

import "fmt"

type Server struct {
	gameState *GameState
}

func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

type GameState struct {
	// playersLock sync.RWMutex
	players []*Player
	messageChannel chan any
}

func (self *GameState) Receive(message any){
	self.messageChannel <- message;
}

func NewGameState() *GameState {
	gameState := &GameState{
		players: make([]*Player, 0),
		messageChannel: make(chan any, 10),
	}

	go gameState.loop();

	return gameState;
}

func (self *GameState) loop(){
	for msg := range self.messageChannel{
		self.handleMessage(msg);
	}
}

func (self *GameState) handleMessage(message any){
	switch msg := message.(type){
		case *Player:
			self.addPlayer(msg);
		default:
			panic("oh, you went here? Hire me 🥱")
	}
}

func (self *GameState) addPlayer(player *Player){
	// self.playersLock.Lock()
	self.players = append(self.players, player);
	//self.playersLock.Unlock()
	fmt.Println(player.Name);
}

type Player struct {
	Name string
}

func (self *Server) handleNewPlayer(player *Player) error {
	self.gameState.Receive(player);
	return nil;
}
