export function getRoomList() {
    return JSON.parse(localStorage.getItem("lastRoomList") ?? "[]")
}

export function appendToRoomList(room: JSON) {
    const oldList = JSON.parse(localStorage.getItem("lastRoomList") ?? "[]")
    let newList
    if (oldList.length >= 3) {
        newList = [room, oldList[0], oldList[1]]
    } else {
        newList = [room, ...oldList]
    }
    localStorage.setItem("lastRoomList", JSON.stringify(newList))
    return newList
}

export function resetRoomList() {
    localStorage.setItem("lastRoomList", JSON.stringify([]))
}