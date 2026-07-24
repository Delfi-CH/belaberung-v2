import axios from 'axios';
import { getUserID } from './auth';
import { MessageAttachment, Message } from './message';

const backendURL = import.meta.env.DEV ? 'http://localhost:8081' : '/api';

export const api = axios.create({
	baseURL: backendURL,
	withCredentials: true
});

export async function getPublicRooms() {
	const res = await api.get('/rooms');
	const uid = getUserID();
	const res2 = await api.get(`/users/${uid}/joined`);
	const joinedRoomIDs = res2.data.map((roomUser) => roomUser.Room.id);
	const rooms = new Set();

	res.data.map((room) => {
		if (joinedRoomIDs.includes(room.id)) {
			return;
		} else {
			rooms.add(room);
			return;
		}
	});
	return [...rooms];
}

export async function getJoinedRooms() {
	const uid = getUserID();
	const res = await api.get(`/users/${uid}/joined`);
	return res.data.map((roomUser) => roomUser.Room);
}

export async function joinRoom(roomID: number | string) {
	try {
		await api.get(`/rooms/${roomID}/join`);
		return 'joined';
	} catch (err) {
		return 'error: ' + err;
	}
}

export async function loadInitialMessages(roomID: number) {
	try {
		const res = await api.get(`/rooms/${roomID}/messages`)

		return res.data.map(element => {
			return Message.fromJson(element)
		});
	} catch (err) {
		console.log(err)
		return []
	}
} 

export function createWebsocket() {
	const ws = new WebSocket(backendURL + '/ws');
	ws.addEventListener("open", ()=>{
		return
	})
	return ws;
}

export function sendMessage(
	ws: WebSocket,
	content: string,
	username: string,
	userID: number,
	roomID: number,
	attachment?: MessageAttachment
) {
	const msg = new Message({
		content: content,
		timestamp: new Date(),
		username: username,
		userID: userID,
		roomID: roomID,
		attachment: attachment
	});
	ws.send(msg.toString());
}

export function streamMessages(ws: WebSocket, callback: (message: Message) => void): void {
	ws.addEventListener('message', (event) => {
			const message = Message.fromJson(JSON.parse(event.data));
			callback(message);
	});
}
