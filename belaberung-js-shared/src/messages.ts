import { Room } from "./rooms.js"
import { User } from "./users.js"

export class Message {
    id: number = 0
    content: string
    attachment: MessageAttachment
    timestamp: Date 
    user: User
    room: Room
    constructor(content: string, attachment: MessageAttachment, user: User, room: Room) {
        this.content = content
        this.attachment = attachment
        this.user = user
        this.room = room
        this.timestamp = new Date()
    }
}

export enum MessageAttachmentType {
    None,
    BinaryFile,
    TextFile,
    Document,
    Hyperlink,
    Image,
    Audio,
    Video
}

export class MessageAttachment {
    type: MessageAttachmentType
    data: any
    constructor(type: MessageAttachmentType, data: any) {
        this.type = type
        this.data = data
    }
}