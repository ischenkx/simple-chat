package postgres

// INPUT: username, password_hash
//
// OUTPUT: id, username, password_hash
const createUserSql = `
	insert into Users
		(username, password_hash)
		values ($1, $2)
	returning Users.id, Users.username, Users.password_hash
`

// INPUT: id
//
// OUTPUT: nil
const deleteUserSql = `
	delete from Users
		where id = $1
`

// INPUT: id
//
// OUTPUT: id, username, password_hash
const getUserSql = `
	select id, username, password_hash from Users
		where id = $1
`

// INPUT: username
//
// OUTPUT: id, username, password_hash
const getUserByUsernameSql = `
	select id, username, password_hash from Users
		where username = $1
`

// INPUT: id, new_username, new_password_hash
//
// OUTPUT: id, username, password_hash
const updateUserSql = `
	update Users
	set username = $2,
		password_hash = $3
	where id = $1
	returning Users.id, Users.username, Users.password_hash
`

// INPUT: user1_id, user2_id
//
// OUTPUT: nil
const createFriendConnectionSql = `
	insert into FriendConnections
		(user1_id, user2_id)
		values ($1, $2)
`

// INPUT: id1, id2
//
// OUTPUT: nil
const deleteFriendConnectionSql = `
	delete from FriendConnections
		where user1_id::text = $1 and user2_id::text = $2
`

// INPUT: id, offset, limit
//
// OUTPUT: id, username, password_hash
const getUserFriendsSql = `
	select u.id, u.username, u.password_hash from FriendConnections
	join Users u on ((u.id = user1_id or u.id = user2_id) and u.id != $1)
	where user1_id = $1 or user2_id = $1
	order by u.id
	offset $2
	limit $3
`

// INPUT: id
//
// OUTPUT: amount_of_friends
const countUserFriendsSql = `
	select count(*) from FriendConnections
	where user1_id = $1 or user2_id = $1
`

// INPUT: user1_id, user2_id
//
// OUTPUT: connection_exists
const friendConnectionCheckerSql = `
	select exists(
		select * from FriendConnections
			where user1_id = $1 and user2_id = $2
				or user1_id = $2 and user2_id = $1
	)
`

// INPUT: from, to, time
//
// OUTPUT: id, from_id, to_id, time
const createFriendRequestSql = `
	insert into FriendRequests as req
	(from_id, to_id, time)
	values ($1, $2, $3)
	returning req.id, req.from_id, req.to_id, req.time
`

// INPUT: id
//
// OUTPUT: nil
const deleteFriendRequestSql = `
	delete from FriendRequests
		where id = $1
`

// INPUT: id, offset, limit
//
// OUTPUT: id, from_id, to_id, time
const getUserIncomingFriendRequestsSql = `
	select id, from_id, to_id, time from FriendRequests
	where to_id = $1
	order by to_id
	offset $2
	limit $3
`

// INPUT: id, offset, limit
//
// OUTPUT: id, from_id, to_id, time
const getUserOutgoingFriendRequestsSql = `
	select id, from_id, to_id, time from FriendRequests
	where from_id = $1
	order by to_id
	offset $2
	limit $3
`

// INPUT: id, from
//
// OUTPUT: id, from_id, to_id, time
const getUserIncomingFriendRequestSql = `
	select id, from_id, to_id, time from FriendRequests
	where from_id = $2 and to_id = $1
`

// INPUT: id, to
//
// OUTPUT: id, from_id, to_id, time
const getUserOutgoingFriendRequestSql = `
	select id, from_id, to_id, time from FriendRequests
	where to_id = $2 and from_id = $1
`

// INPUT: id
//
// OUTPUT: count
const countUserIncomingFriendRequestsSql = `
	select count(*) from FriendRequests
		where to_id = $1
`

// INPUT: id
//
// OUTPUT: count
const countUserOutgoingFriendRequestsSql = `
	select count(*) from FriendRequests
		where from_id = $1
`

// INPUT: from, to
//
// OUTPUT: id, from_id, to_id, time
const getFriendRequestSql = `
	select id, from_id, to_id, time from FriendRequests
		where from_id = $1 and to_id = $2
`

// INPUT: id
//
// OUTPUT: id, from_id, to_id, time
const getFriendRequestByIDSql = `
	select id, from_id, to_id, time from FriendRequests
		where id = $1
`

// INPUT: name, description, owner_id
//
// OUTPUT: id, name, description, owner_id
const createChatSql = `
	insert into Chats as chat
	(chat_name, description, owner_id)
	values ($1, $2, $3)
	returning chat.id, chat.chat_name, chat.description, chat.owner_id
`

// INPUT: id
//
// OUTPUT: nil
const deleteChatSql = `
	delete from Chats
		where id = $1
`

// INPUT: id
//
// OUTPUT: id, name, description, owner_id
const getChatSql = `
	select id, chat_name, description, owner_id from Chats
		where id = $1
`

// INPUT: id, name, description
//
// OUTPUT: id, name, description, owner_id
const updateChatSql = `
	update Chats
	set chat_name = $2,
		description = $3
	where id = $1
	returning Chats.id, Chats.chat_name, Chats.owner_id
`

// INPUT: user_id, chat_id, status
//
// OUTPUT: user_id, chat_id, status
const createChatMemberSql = `
	insert into ChatMembers as mem
		(user_id, chat_id, status)
		values ($1, $2, $3)
		returning mem.user_id, mem.chat_id, mem.status
`

// INPUT: user_id, chat_id
//
// OUTPUT: nil
const deleteChatMemberSql = `
	delete from ChatMembers
		where user_id = $1 and chat_id = $2
`

// INPUT: user_id, chat_id, status
//
// OUTPUT: user_id, chat_id, status
const updateChatMemberSql = `
	update ChatMembers
	set status = $2
	where user_id = $1 and chat_id = $2
	returning ChatMembers.user_id, ChatMembers.chat_id, ChatMembers.status
`

// INPUT: userId, chatId
//
// OUTPUT: user_id, chat_id, status
const getChatMemberSql = `
	select user_id, chat_id, status from ChatMembers
		where user_id = $1 and chat_id = $2
`

// INPUT: user_id, chat_id, payload, timestamp, last_update
//
// Output: id, user_id, chat_id, payload, timestamp, last_update
const createMessageSql = `
	insert into Messages as mes
	(user_id, chat_id, payload, time, last_update)
	values ($1, $2, $3, $4, $5)
	returning mes.id, mes.user_id, mes.chat_id, mes.payload, mes.time, mes.last_update
`

// INPUT: id
//
// OUTPUT: nil
const deleteMessageSql = `
	delete from Messages
		where id = $1
`

// INPUT: id, payload, last_update
//
// OUTPUT: nil
const updateMessageSql = `
	update Messages as mes
	set payload = $2,
		last_update = $3
	where id  = $1
	returning mes.id, mes.user_id, mes.chat_id, mes.payload, mes.time, mes.last_update
`

// INPUT: id
//
// OUTPUT: id, user_id, chat_id, payload, time, last_update
const getMessageSql = `
	select id, user_id, chat_id, payload, time, last_update from Messages
		where id = $1
`

// INPUT: user_id, offset, count
//
// OUTPUT: user_id, chat_id, status
const getUserChatsSql = `
	select user_id, chat_id, status from ChatMembers
		where user_id = $1
		order by user_id
		offset $2
		limit $3
`

// INPUT: chat_id, offset, count
//
// OUTPUT: user_id, chat_id, status
const getChatMembersSql = `
	select user_id, chat_id, status from ChatMembers
		where chat_id = $1
		order by user_id
		offset $2
		limit $3
`

// INPUT: chat_id, offset, count
//
// OUTPUT: id, user_id, chat_id, payload, time, last_update
const getChatMessagesSql = `
	select id, user_id, chat_id, payload, time, last_update from Messages
		where chat_id = $1
		order by time desc
		offset $2
		limit $3
`

// INPUT: chat_id
//
// OUTPUT: count
const countChatMembersSql = `
	select count(*) from ChatMembers
		where chat_id = $1
`

// INPUT: chat_id
//
// OUTPUT: count
const countChatMessagesSql = `
	select count(*) from Messages
		where chat_id = $1
`

// INPUT: user_id
//
// OUTPUT: count
const countUserChatsSql = `
	select count(*) from ChatMembers
		where user_id = $1
`

//

const initializeTablesSql = `
-- Extensions
create extension if not exists "uuid-ossp";

-- Tables
create table if not exists Users (
	id uuid default uuid_generate_v1() primary key,
	username varchar (40) unique not null,
	password_hash varchar (200) not null
);

create table if not exists Chats (
	id uuid default uuid_generate_v1() primary key,
	chat_name varchar (40) not null,
	description varchar (500),
	owner_id uuid not null,
	
	foreign key (owner_id)
		references Users (id)
);

create table if not exists ChatMembers (
	user_id uuid not null,
	chat_id uuid not null,
	status int not null,
	
	foreign key (user_id)
		references Users (id)
			on delete cascade,
	foreign key (chat_id)
		references Chats (id)
			on delete cascade,
	primary key (user_id, chat_id)
);

create table if not exists FriendConnections (
	user1_id uuid not null,
	user2_id uuid not null,
	
	foreign key (user1_id)
		references Users (id),
	foreign key (user2_id)
		references Users (id),
	primary key (user1_id, user2_id)
);

create table if not exists FriendRequests (
	id uuid default uuid_generate_v1() primary key,
	from_id uuid not null,
	time timestamp default now(),
	to_id uuid not null,
	
	foreign key (from_id)
		references Users (id),
	foreign key (to_id)
		references Users (id)
);

create table if not exists Messages (
	id uuid  default uuid_generate_v1() primary key,
	user_id uuid not null,
	chat_id uuid not null,
	payload varchar (400) not null,
	time timestamp default now(),
	last_update timestamp,
	
	foreign key (user_id, chat_id)
		references ChatMembers (user_id, chat_id) on delete cascade
);

-- Indices

create index if not exists "index_message_time"
on Messages using btree (time);
`
