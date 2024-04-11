# Lecture 20 - Postgres Database

## Postgres with Docker

* docker run --name learnpostgresdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:latest

* docker ps

* docker exec -it {container id} bash (บ่งบอกว่าเราจะใช้ shell ประเภท bash ใน container นี้

* psql -U postgres

## Inside the Postgres 

* CREATE DATABASE testdb;

* \\l -> เพื่อดู tables

* DROP DATABASE testdb; -> ลบ table

* \\! clear -> เคลียร์หน้าจอ CMD

* \\c testdb -> เข้าใช้งาน table testdb

## Open DBeaver

New connection to Postgres. Which I run the docker will be 
Host: localhost 
Port: 5432
Database: testdb
Username: postgres # default
Password: 123456

F3 to run script in DBeaver and create alt+X

Properties will see the data models and ER diagram.

## Postgresql
Insensitive case query :

* SELECT * FROM items
WHERE name ilike '%blade%';

* ORDER BY level ASC: -> Ascending น้อยไปมาก
ORDER BY level DESC;

เปลี่ยนแปลงข้อมูลที่เคย INSERT

* UPDATE players SET name = 'Noobza007' WHERE id = 3;

### Caution! PostgreSQL use only '' not "" it is not the same!

* DELETE FROM players WHERE id = 3;

* select *
from players as p
full outer join items as i
on i.player_id = p.id;

(join -> inner join, left join , right join)

* Database Transaction
BEGIN;

INSERT INTO players (
  name, level)
VALUES ('Rayato159', 1),
('JOSH', Error); -> Error! Database will rollback now. but auto_increment will remove the id that has created

COMMIT;

Transaction มันจะย้อนหลังให้ทั้งหมดใน BEGIN-COMMIT ทั้งหมดนั่นเอง

## Semantic Versioning 2.0.0

Major -> Version when you make incompatible API changes การที่เราเปลี่ยนแอพครั้งใหญ่
Minor -> Feature การทำงานใหม่ๆเข้ามา
Patch -> Feature ที่เราทำมาเราแก้ไปแล้วกี่ครั้ง

## Git Merge and Git Rebase (Github Workflow)

* Git rebase คือการทำให้ Commit ของอีก Branch มาเป็นฐานของเราใหม่ ก็คือต้อง Rebase ทั้ง 2 branch ด้วยกัน
