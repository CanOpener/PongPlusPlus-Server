package games

import (
	"encoding/binary"
	"github.com/canopener/PongPlusPlus-Server/server/connection"
	"github.com/canopener/serverlog"
	"github.com/satori/go.uuid"
	"time"
)

// Game is a structure which handles a game
type Game struct {
	// ID is the id of the game
	ID string
	// Name is the name of the game given by the initiator
	Name string
	// Initiator is the connection of the player who started the game
	Initiator *connection.Conn
	// Player2 is the connection of player 2
	Player2 *connection.Conn
	// InitTime is the time the game object was initiated
	InitTime time.Time
	// StartTime is the time the game was started
	BeginTime time.Time
	// Ready is true if both players are in the game
	Ready bool
}

// NewGame returns a pointer to a game instance given two connections
func NewGame(initiator *connection.Conn, name string) *Game {
	id := uuid.NewV4().String()
	serverlog.General("Initiating game:", id, "with", initiator.Identification())
	return &Game{
		ID:        id,
		Name:      name,
		Initiator: initiator,
		InitTime:  time.Now(),
		Ready:     false,
	}
}

// Start will start the game
func (g *Game) Start(player2 *connection.Conn) {
	serverlog.General("Starting game:", g.ID, "with player 2", player2.Identification())
	g.Player2 = player2
	g.Ready = true
	// TODO: start unix domain server
	// TODO: start player message listeners
}

// Kill destroys all game related gorutines
func (g *Game) Kill() {
	// TODO: kill unix domain socket server
	// TODO: kill player message listeners
	g.Initiator.InGame = false
	if g.Ready {
		g.Player2.InGame = false
	}
}

// Bytes returns an API friendly binary representation of the game object
// which can be sent to clients.
func (g *Game) Bytes() []byte {
	serverlog.General("Getting byte version of game:", g.Name)
	ubts := make([]byte, 4)
	unix := uint32(g.InitTime.Unix())
	binary.LittleEndian.PutUint32(ubts, unix)
	gid := append([]byte(g.ID), byte('\000'))
	gname := append([]byte(g.Name), byte('\000'))
	pname := append([]byte(g.Initiator.Alias), byte('\000'))

	ret := append(ubts, gid...)
	ret = append(ret, gname...)
	ret = append(ret, pname...)
	return ret
}

// Identification returns a human readable way of differenciating
// between games
func (g *Game) Identification() string {
	return "Game Named: " + g.Name
}
