import { Room } from "./rooms.js";
import { User } from "./users.js";

export class Message {
  id: number = 0;
  content: string;
  attachment: MessageAttachment;
  timestamp: Date;
  user: User;
  room: Room;
  constructor(
    content: string,
    attachment: MessageAttachment,
    user: User,
    room: Room,
  ) {
    this.content = content;
    this.attachment = attachment;
    this.user = user;
    this.room = room;
    this.timestamp = new Date();
  }
}

export enum MessageAttachmentType {
  None = "None",
  Mention = "Mention",
  Reply = "Reply",
  BinaryFile = "BinaryFile",
  TextFile = "TextFile",
  Document = "Document",
  Hyperlink = "Hyperlink",
  Image = "Image",
  Audio = "Audio",
  Video = "Video",
}

export class MessageAttachment {
  type: MessageAttachmentType;
  data: any;
  constructor(type: MessageAttachmentType, data: any) {
    this.type = type;
    this.data = data;
  }
}
