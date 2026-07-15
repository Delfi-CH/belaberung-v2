import axios from 'axios';
import { getUserID } from './auth';

const backendURL = import.meta.env.DEV ? 'http://localhost:8081' : '/api';

export const api = axios.create({
	baseURL: backendURL,
	withCredentials: true
});

export async function getPublicRooms() {
	const res = await api.get("/rooms")
	const uid = getUserID()
	const res2 = await api.get(`/users/${uid}/joined`)
	if (res2.data === null) {
		return res.data
	}
	const joinedRoomsId = res2.data.map((e)=> e.Room.id)
	const notJoinedRooms = res.data.map((e)=>{
		if (joinedRoomsId.includes(e.id)) {
			return {} 
		} else {
			return e
		}
	})
	return notJoinedRooms
}

export async function getJoinedRooms() {
	const uid = getUserID()
	const res = await api.get(`/users/${uid}/joined`)
	if (res.data === null) {
		return []
	}
	const joinedRooms = res.data.map((e)=> e.Room)
	return joinedRooms
}