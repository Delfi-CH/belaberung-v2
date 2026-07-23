package ws

type Hub struct {
	// roomID -> clients
	rooms map[int]map[*client]bool

	register chan roomSubscription

	unregister chan roomSubscription

	broadcast chan Message
}


type roomSubscription struct {
	client *client
	roomID int
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[int]map[*client]bool),

		register: make(chan roomSubscription),

		unregister: make(chan roomSubscription),

		broadcast: make(chan Message),
	}
}


// Run starts the hub event loop.
//
// Run should typically be started in its own goroutine:
//
//	hub := ws.NewHub()
//	go hub.Run()
func (h *Hub) Run() {
	for {
		select {

		case subscription := <-h.register:

			clients, exists := h.rooms[subscription.roomID]

			if !exists {
				clients = make(map[*client]bool)
				h.rooms[subscription.roomID] = clients
			}

			clients[subscription.client] = true


		case subscription := <-h.unregister:

			if clients, exists := h.rooms[subscription.roomID]; exists {

				delete(
					clients,
					subscription.client,
				)

				if len(clients) == 0 {
					delete(
						h.rooms,
						subscription.roomID,
					)
				}
			}


		case message := <-h.broadcast:

			clients, exists := h.rooms[message.RoomID]

			if !exists {
				continue
			}

			for client := range clients {

				select {

				case client.send <- message:

				default:
					close(client.send)
					delete(
						clients,
						client,
					)
				}
			}
		}
	}
}

func (h *Hub) Subscribe(
	client *client,
	roomID int,
) {
	h.register <- roomSubscription{
		client: client,
		roomID: roomID,
	}
}


func (h *Hub) Unsubscribe(
	client *client,
	roomID int,
) {
	h.unregister <- roomSubscription{
		client: client,
		roomID: roomID,
	}
}


func (h *Hub) Broadcast(message Message) {
	h.broadcast <- message
}
