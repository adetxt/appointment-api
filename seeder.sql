insert into users (`name`, `role`, timezone, updated_at) values
(`Ade`, `USER`, `Asia/Jakarta`, now()),
(`Christy Schumm`, `COACH`, `America/North_Dakota/New_Salem`, now()),
(`Natalia Stanton Jr.`, `COACH`, `Central Time (US & Canada)`, now()),
(`Nola Murazik V`, `COACH`, `America/Yakutat`, now()),
(`Elyssa O\`Kon`, `COACH`, `Central Time (US & Canada)`, now()),
(`Dr. Geovany Keebler`, `COACH`, `Central Time (US & Canada)`, now()),
(`Fulan`, `USER`, `Asia/Jakarta`, now());

insert into working_hours(user_id, `day`, `start`, `end`, updated_at) values
(2, 0, `09:00:00`, `17:30:00`, now()),
(2, 1, `08:00:00`, `16:00:00`, now()),
(2, 3, `09:00:00`, `16:00:00`, now()),
(2, 4, `07:00:00`, `14:00:00`, now()),

(3, 1, `08:00:00`, `10:00:00`, now()),
(3, 2, `11:00:00`, `18:00:00`, now()),
(3, 5, `09:00:00`, `15:00:00`, now()),
(3, 6, `08:00:00`, `15:00:00`, now()),

(4, 0, `08:00:00`, `10:00:00`, now()),
(4, 1, `11:00:00`, `13:00:00`, now()),
(4, 2, `08:00:00`, `10:00:00`, now()),
(4, 5, `08:00:00`, `11:00:00`, now()),
(4, 6, `07:00:00`, `09:00:00`, now()),

(5, 0, `09:00:00`, `15:00:00`, now()),
(5, 1, `06:00:00`, `13:00:00`, now()),
(5, 2, `06:00:00`, `11:00:00`, now()),
(5, 4, `08:00:00`, `12:00:00`, now()),
(5, 5, `09:00:00`, `16:00:00`, now()),
(5, 6, `08:00:00`, `10:00:00`, now()),

(6, 3, `07:00:00`, `14:00:00`, now()),
(6, 3, `15:00:00`, `17:00:00`, now());