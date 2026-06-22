import { Room } from "./rooms.js";

export class User {
  id: number = 0;
  username: string;
  biography: string = "";
  profilePicture: string = "default";
  pronouns: string = "";
  domain: string;
  password: string;
  suspended: boolean = false;
  globalRole: GlobalUserRoles = GlobalUserRoles.Member;
  rooms: Room[] = [];

  constructor(username: string, domain: string, password: string) {
    this.username = username;
    this.domain = domain;
    this.password = password;
  }
}

export enum GlobalUserRoles {
  Member = "Member",
  InstanceModerator = "InstanceModerator",
  InstanceAdministrator = "InstanceAdministrator",
}
