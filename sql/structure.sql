CREATE TABLE IF NOT EXISTS appointments (
    id bigint PRIMARY KEY,
    trainer_id bigint NOT NULL,
    user_id bigint,
    starts_at timestamptz NOT NULL,
    ends_at timestamptz NOT NULL,
    status text NOT NULL
);

INSERT INTO appointments (
    id,
    trainer_id,
    user_id,
    starts_at,
    ends_at,
    status)
VALUES (
    1,
    1,
    1,
    '2022-03-24T09:00:00-08:00',
    '2022-03-24T09:30:00-08:00',
    'BOOKED'),
(
    2,
    1,
    2,
    '2022-03-24T10:00:00-08:00',
    '2022-03-24T10:30:00-08:00',
    'BOOKED'),
(
    3,
    1,
    3,
    '2022-03-25T10:00:00-08:00',
    '2022-03-25T10:30:00-08:00',
    'BOOKED'),
(
    4,
    1,
    4,
    '2022-03-25T10:30:00-08:00',
    '2022-03-24T11:00:00-08:00',
    'BOOKED'),
(
    5,
    1,
    5,
    '2022-03-26T10:00:00-08:00',
    '2022-03-26T10:30:00-08:00',
    'BOOKED'),
(
    6,
    2,
    6,
    '2022-03-24T09:00:00-08:00',
    '2022-03-24T09:30:00-08:00',
    'BOOKED'),
(
    7,
    2,
    7,
    '2022-03-26T10:00:00-08:00',
    '2022-03-26T10:30:00-08:00',
    'BOOKED'),
(
    8,
    3,
    8,
    '2022-03-26T12:00:00-08:00',
    '2022-03-26T12:30:00-08:00',
    'BOOKED'),
(
    9,
    3,
    9,
    '2022-03-26T13:00:00-08:00',
    '2022-03-26T13:30:00-08:00',
    'BOOKED'),
(
    10,
    3,
    10,
    '2022-03-26T14:00:00-08:00',
    '2022-03-26T14:30:00-08:00',
    'BOOKED'),
(
    11,
    3,
    NULL,
    '2022-03-26T14:30:00-08:00',
    '2022-03-26T15:00:00-08:00',
    'OPEN'),
(
    12,
    1,
    NULL,
    '2022-03-26T14:30:00-08:00',
    '2022-03-26T15:00:00-08:00',
    'OPEN'),
(
    13,
    1,
    NULL,
    '2022-03-26T13:30:00-08:00',
    '2022-03-26T14:00:00-08:00',
    'OPEN'),
(
    14,
    2,
    NULL,
    '2022-03-26T13:30:00-08:00',
    '2022-03-26T14:00:00-08:00',
    'OPEN');