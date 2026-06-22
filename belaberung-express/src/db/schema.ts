import {
  integer,
  pgTable,
  varchar,
  boolean,
  pgEnum,
  timestamp,
  text
} from "drizzle-orm/pg-core";

export const globalUserRoles = pgEnum("global_user_roles", [
  "Member",
  "InstanceModerator",
  "InstanceAdministrator",
]);

export const roomUserRoles = pgEnum("room_user_roles", [
    "Suspended",
  "Member",
  "Moderator",
  "Administrator",
]);

export const messageAttachmentType= pgEnum("message_attachment_type", [
  "None",
  "Mention",
  "Reply",
  "BinaryFile",
  "TextFile",
  "Document",
  "Hyperlink",
  "Image",
  "Audio",
  "Video",
]);

export const usersTable = pgTable("users", {
  id: integer().primaryKey().generatedAlwaysAsIdentity(),
  username: varchar({ length: 255 }).notNull().unique(),
  password: varchar({ length: 255 }).notNull(),
  biography: varchar({ length: 1024 }),
  profilePicture: varchar("profile_picture", { length: 255 })
    .notNull()
    .default("default"),
  pronouns: varchar({ length: 255 }),
  domain: varchar({ length: 255 }).notNull(),
  suspended: boolean().notNull().default(false),
  globalRole: globalUserRoles("global_role").notNull().default("Member"),
});

export const roomsTable = pgTable("rooms", {
  id: integer().primaryKey().generatedAlwaysAsIdentity(),
  name: varchar({ length: 255 }).notNull().unique(),
  description: varchar({ length: 1024 }).notNull(),
  domain: varchar({ length: 255 }).notNull(),
});

export const roomUsersTable = pgTable("room_users", {
  id: integer().primaryKey().generatedAlwaysAsIdentity(),
  userId: integer("user_id").references(()=> usersTable.id),
  role: roomUserRoles().notNull().default("Member")
});

export const usersRoomsTable = pgTable("users_rooms", {
    roomId: integer("room_id").references(() => roomsTable.id),
    userId: integer("user_id").references(() => roomUsersTable.id)
})

export const messagesTable = pgTable("messages", {
    id: integer().primaryKey().generatedAlwaysAsIdentity(),
    content: varchar({ length: 1024 }).notNull(),
    timestamp: timestamp({mode: "date"}).notNull().defaultNow(),
    roomId: integer("room_id").references(() => roomsTable.id),
    userId: integer("user_id").references(() => usersTable.id),
    attachmentId: integer("attachment_id").references(()=> messageAttachmentTable.id)
})

export const messageAttachmentTable = pgTable("message_attachment", {
    id: integer().primaryKey().generatedAlwaysAsIdentity(),
    type: messageAttachmentType().notNull().default("None"),
    data: text(),
})