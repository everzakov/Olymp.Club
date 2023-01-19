CREATE TABLE UserModel(
  id serial primary key,
  email text not null,
  pass_hash text not null,
  token1 text not null,
  token2 text not null
);

CREATE TABLE UnConfirmedUsers (
  id serial primary key,
  email text not null,
  token1 text not null,
  token2 text not null,
  pass_hash text not null,
  confirmed boolean
);

CREATE TABLE SessionModel (
 id serial primary key,
 user_id int4 not null,
 token text not null,
 expiry timestamp default (now() + '01:00:00'::interval),
 foreign key (user_id) REFERENCES UserModel(id)
);

CREATE TABLE AdminModel (
 id serial primary key,
 user_id int4 not null,
 priority int4 not null,
 foreign key (user_id) REFERENCES UserModel(id)
);

CREATE TABLE BigOlympiadModel (
 big_olympiad_id serial primary key,
 name text not null,
 short text not null,
 logo text,
 description text,
 status text
);

CREATE TABLE HolderModel (
 holder_id serial primary key,
 name text not null,
 logo text
);

CREATE TABLE OlympiadModel (
 id serial primary key,
 name text not null,
 subject text not null,
 level text not null,
 img text not null,
 short text not null,
 big_olympiad_id int4 not null,
 status text not null,
 grade text not null,
 holder_id int4 not null,
 website text,
 foreign key (big_olympiad_id) REFERENCES BigOlympiadModel(big_olympiad_id),
 foreign key (holder_id) REFERENCES HolderModel(holder_id)
);

CREATE TABLE OlympiadUserModel (
 id serial primary key,
 olympiad_id int4,
 user_id int4,
 foreign key (olympiad_id) REFERENCES OlympiadModel(id),
 foreign key (user_id) REFERENCES UserModel(id)
);

CREATE TABLE EventModel (
 id serial primary key,
 name text not null,
 description text,
 short text,
 img text,
 status text not null,
 holder_id int4 not null,
 website text,
 foreign key (holder_id) REFERENCES HolderModel(holder_id)
);

CREATE TABLE EventUserModel (
 id serial primary key,
 event_id int4 not null,
 user_id int4 not null,
 foreign key (event_id) REFERENCES EventModel(id),
 foreign key (user_id) REFERENCES UserModel(id)
);

CREATE TABLE NewsModel (
 id serial primary key,
 title text not null,
 description text,
 tablestruct text not null,
 key int4 not null
);
