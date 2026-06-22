import { Message } from "./messages.js";
import { User } from "./users.js";

export class Room {
  id: number = 0;
  name: string;
  description: string;
  domain: string;
  users: RoomUser[] = [];

  constructor(name: string, description: string, domain: string) {
    this.name = name;
    this.description = description;
    this.domain = domain;
  }
}

export class RoomUser {
  user: User;
  role: RoomRoles = RoomRoles.Member;
  messages: Message[] = [];
  constructor(user: User) {
    this.user = user;
  }
}

export enum RoomRoles {
  Suspended = "Suspended",
  Member = "Member",
  Moderator = "Moderator",
  Administrator = "Administrator",
}
