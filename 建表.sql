

create table VideoOrder(OrderId int, \
	StartTime DateTime, EndTime DateTime, \
	RequestBy int, FulfilledBy int, Duration Integer,
	Subject varchar(50), Level int, Keywords varchar(256), \
	Tutor_Student_R int, 
	Student_Tutor_R int, 
	Primary Key(OrderId));



create table QuestionOrder(
	QId int, \
	StartTime DateTime, EndTime DateTime, \
	RequestBy int, FulfilledBy int,
	Subject varchar(50), Level int, Keywords varchar(256), \
	Student_Tutor_R int, Primary Key(QId)
	);


create table Tutor(
	TutorId int,
	Level int, 
	FirstName varchar(50),
	LastName varchar(50),
	Account_Balance int,
	Rating int,
	Primary Key(TutorId)
);

create table Student(
	StudentId int, 
	FirstName varchar(50),
	LastName varchar(50),
	Level int,
	Account_Balance int,
	Rating int,
	Primary Key(StudentId)
);

create table Users(
	Id int,
	email varchar(256),
	password varchar(256),
	birthday date,
	StudentId,
	TutorId,
	Foreign Key(StudentId) References Student(StudentId) ON DELETE CASCADE ON Update CASCADE
	Foreign Key(TutorId) References Tutor(TutorId) ON DELETE CASCADE ON Update CASCADE
)


create table Student_Preference(
	id int, 
	Prefered_tutor int,
	Comments varchar(256),
	Foreign Key(id) References Student(StudentId) ON DELETE CASCADE ON Update CASCADE
);


create table Student_History(
	id int, 
	VideoOrder int,
	QuestionOrder int,
	Foreign Key(id) References Student(StudentId) ON DELETE CASCADE ON Update CASCADE
);


create table Tutor_History(
	id int, 
	VideoOrder int,
	QuestionOrder int,
	Foreign Key(id) References Tutor(TutorId) ON DELETE CASCADE ON Update CASCADE
);