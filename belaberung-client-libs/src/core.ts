import axios from 'axios';
import { getUserID } from './auth';
import { MessageAttachment, Message } from './message';

let backendURL = '/api';

export function setBackendUrl(url: string) {
	backendURL = url
}

export function getBackendUrl() {
	return backendURL
}

export function createAPI(backendURL: string) {
	return axios.create({
	baseURL: backendURL,
	withCredentials: true
});
}

export async function getPublicRooms() {
	const api = createAPI(getBackendUrl())
	const res = await api.get('/rooms');
	const uid = getUserID();
	const res2 = await api.get(`/users/${uid}/joined`);
	const joinedRoomIDs = res2.data.map((roomUser: any) => roomUser.Room.id);
	const rooms = new Set();

	res.data.map((room: any) => {
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
	const api = createAPI(getBackendUrl())
	const uid = getUserID();
	const res = await api.get(`/users/${uid}/joined`);
	return res.data.map((roomUser: any) => roomUser.Room);
}

export async function joinRoom(roomID: number | string) {
	const api = createAPI(getBackendUrl())
	try {
		await api.get(`/rooms/${roomID}/join`);
		return 'joined';
	} catch (err) {
		return 'error: ' + err;
	}
}

export async function loadInitialMessages(roomID: number) {
	const api = createAPI(getBackendUrl())
	try {
		const res = await api.get(`/rooms/${roomID}/messages`);

		return res.data.map((element: any) => {
			return Message.fromJson(element);
		});
	} catch (err) {
		console.log(err);
		return [];
	}
}

export function createWebsocket() {
	const ws = new WebSocket(getBackendUrl() + '/ws');
	ws.addEventListener('open', () => {
		return;
	});
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
	const api = createAPI(getBackendUrl())
	api.get(`/rooms/${roomID}/me`).then((res) => {
		const role = res.data.role;
		console.log(res.data.role);
		const msg = new Message({
			content: content,
			timestamp: new Date(),
			username: username,
			userID: userID,
			role: role,
			roomID: roomID,
			attachment: attachment
		});
		ws.send(msg.toString());
	});
}

export function streamMessages(ws: WebSocket, callback: (message: Message) => void): void {
	ws.addEventListener('message', (event) => {
		const message = Message.fromJson(JSON.parse(event.data));
		callback(message);
	});
}

export async function getUserDetails(userID: number | string) {
	const api = createAPI(getBackendUrl())
	try {
		const res = await api.get(`/users/${userID}`);
		return res.data
	} catch (err) {
		return 'error: ' + err;
	}
}