insert into usuarios (nome, nick, email, senha)
values
("usuario 1", "usuario1", "usuario1@test.com", "$2a$10$TNbJpQ7XgveC03KMYcWLA.gqCtkzRnZNIeAFIRhfTjYjHEUGUl5ee"),
("usuario 2", "usuario2", "usuario2@test.com", "$2a$10$Qp6c2lqyP63ieJiGt7QnK.s5TCnCSWvkdD0Xn42tj363gLC785v3m"),
("usuario 3", "usuario3", "usuario3@test.com", "$2a$10$3rQYoYp0.z/fTSGbMNkql.i9.46x2nT9KF96R4IZ0Nw9oA8IVLSPq"),
("usuario 4", "usuario4", "usuario4@test.com", "$2a$10$C.8z0VxJkKnuv7BpOmxH.ukbK0w27UTzHqec6GgGT78EVABFejJK."),
("usuario 5", "usuario5", "usuario5@test.com", "$2a$10$2izqcIbDPm8FsLE1MJTr4eCESDfbO0DzP1B7h03qhFTTVAMSzLDDW"),
("usuario 6", "usuario6", "usuario6@test.com", "$2a$10$TZQmfnamRfcReb8nVExgUOe/MHam2lYwh3yqHXMeBOFGKU.9fk3nC"),
("usuario 7", "usuario7", "usuario7@test.com", "$2a$10$N3OaPVjbQ9IVfUpnECnEbezrrHQg7SUTg5l7XHHF8oXdIzD1GBwxq"),
("usuario 8", "usuario8", "usuario8@test.com", "$2a$10$IweE/kB08SfI7l94klhBqewBcTyxh6m2qj5ZVta0ihIvwoPwvz.KC");

insert into seguidores (usuario_id, seguidor_id)
values
(1,2),
(1,3),
(2,2),
(2,4),
(3,1),
(6,7),
(4,5),
(4,1),
(5,8),
(7,6);

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicacao 1", "Publicacao 1", 1),
("Publicacao 2", "Publicacao 1", 2),
("Publicacao 3", "Publicacao 1", 3),
("Publicacao 4", "Publicacao 1", 4);