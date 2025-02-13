-- Inserindo categorias
INSERT INTO categories(description) VALUES ('Curso');
INSERT INTO categories(description) VALUES ('Oficina');

INSERT INTO attendees(name, email) VALUES ('José Silva', 'jose@gmail.com');
INSERT INTO attendees(name, email) VALUES ('Tiago Faria', 'tiago@gmail.com');
INSERT INTO attendees(name, email) VALUES ('Maria do Rosario', 'maria@gmail.com');
INSERT INTO attendees(name, email) VALUES ('Teresa Silva', 'teresa@gmail.com');

INSERT INTO activities(name, description, price, category_id) 
VALUES ('Curso de HTML', 'Aprenda HTML de forma prática', 80.00, 1);

INSERT INTO activities(name, description, price, category_id) 
VALUES ('Oficina de Github', 'Controle versões de seus projetos', 50.00, 2);

INSERT INTO time_blocks(start_time, end_time, activity_id) 
VALUES ('2017-09-25T08:00:00Z', '2017-09-25T11:00:00Z', 1);

INSERT INTO time_blocks(start_time, end_time, activity_id) 
VALUES ('2017-09-25T14:00:00Z', '2017-09-25T18:00:00Z', 2);

INSERT INTO time_blocks(start_time, end_time, activity_id) 
VALUES ('2017-09-26T08:00:00Z', '2017-09-26T11:00:00Z', 2);

INSERT INTO attendee_activity(attendee_id, activity_id) VALUES (1, 1);
INSERT INTO attendee_activity(attendee_id, activity_id) VALUES (1, 2);
INSERT INTO attendee_activity(attendee_id, activity_id) VALUES (2, 1);
INSERT INTO attendee_activity(attendee_id, activity_id) VALUES (3, 1);
INSERT INTO attendee_activity(attendee_id, activity_id) VALUES (3, 2);
INSERT INTO attendee_activity(attendee_id, activity_id) VALUES (4, 2);

