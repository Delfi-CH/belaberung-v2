import axios from 'axios';
import { getUserID } from './auth';

const backendURL = import.meta.env.DEV ? 'http://localhost:8081' : '/api';

export const api = axios.create({
	baseURL: backendURL,
	withCredentials: true
});

export async function getPublicRooms() {
  const res = await api.get("/rooms");
  const uid = getUserID()
  const res2 = await api.get(`/users/${uid}/joined`)
  const joinedRoomIDs = res2.data.map((roomUser)=> roomUser.Room.id)
  const rooms = new Set()
  
  res.data.map((room)=>{
	if (joinedRoomIDs.includes(room.id)) {
		  return
	  } else {
      rooms.add(room)
      return
    }
  })
  return [...rooms]
}
export async function getJoinedRooms() {
	const uid = getUserID()
	const res = await api.get(`/users/${uid}/joined`)
	return res.data.map((roomUser)=> roomUser.Room)
}

export async function joinRoom(roomID: number|string) {
    try {
      await api.get(`/rooms/${roomID}/join`)
      return "joined"
    } catch (err) {
      return "error: " + err
    }
}