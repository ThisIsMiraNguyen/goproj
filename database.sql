create database todo_app;
use todo_app;
create table todos (
id int auto_increment primary key,
title varchar(255) not null,
completed boolean not null default 0
);
use todo_app;
INSERT INTO todos (title, completed) VALUES
('Complete project report', 0),
('Grocery shopping', 1),
('Read a book', 0),
('Exercise for 30 minutes', 1),
('Schedule dentist appointment', 0),
('Send email update to manager', 1),
('Plan weekend trip', 0),
('Clean the house', 0),
('Watch a tutorial on SQL', 1),
('Buy birthday gift for friend', 0),
('Update resume', 1),
('Organize files on computer', 0),
('Prepare for team meeting', 1),
('Call family', 0),
('Review budget', 1),
('Check messages', 0),
('Refill water bottle', 1),
('Plan weekly meals', 0),
('Set up a new workout routine', 1),
('Write a journal entry', 0);

