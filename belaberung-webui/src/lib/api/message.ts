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
  data: unknown;

  constructor(
    type: MessageAttachmentType = MessageAttachmentType.None,
    data: unknown = null
  ) {
    this.type = type;
    this.data = data;
  }
}

export class Message {
  content: string;
  attachment?: MessageAttachment;
  timestamp: Date;
  username: string;
  userId: number;
  roomId: number;

  constructor(params: {
    content: string;
    timestamp: Date | string;
    username: string;
    userID: number;
    roomID: number;
    attachment?: MessageAttachment;
  }) {
    this.username = params.username
    this.content = params.content;
    this.attachment = params.attachment;
    this.timestamp =
      params.timestamp instanceof Date
        ? params.timestamp
        : new Date(params.timestamp);
    this.userId = params.userID;
    this.roomId = params.roomID;
  }

  static fromJson(json: any): Message {
    return new Message({
      content: json.content,
      timestamp: json.timestamp,
      username: json.username,
      userID: json.user_id,
      roomID: json.room_id,
      attachment: json.attachment
        ? new MessageAttachment(
            json.attachment.type,
            json.attachment.data
          )
        : undefined,
    });
  }

  toJson() {
    return {
      content: this.content,
      attachment: this.attachment,
      timestamp: this.timestamp.toISOString(),
      username: this.username,
      user_id: this.userId,
      room_id: this.roomId,
    };
  }

  toString() {
    const json = this.toJson()
    return JSON.stringify(json)
  }
}