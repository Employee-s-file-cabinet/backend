-- start transaction
begin;

truncate table users restart identity cascade;
truncate table authorizations restart identity cascade;
truncate table roles restart identity cascade;
truncate table organization_structure restart identity cascade;
truncate table positions restart identity cascade;
truncate table departments restart identity cascade;

insert into public.roles (id, title, description)
values  (1, 'admin', 'Администратор системы. Имеет доступ к созданию аккаунтов для HR и рекрутёров, их настройке.'),
        (2, 'hr', 'HR компании. Имеет доступ к созданию и редактированию данных сотрудников (карточки)'),
        (3, 'recruiter', 'Специалист по подбору персонала. Может просматривать только данные кандидатов.'),
        (4, 'employee', 'Сотрудник компании. Может просматривать только свои данные.'),
        (5, 'candidate', 'Кандидат на должность в компании. Может создать и отправить свою анкету на вакансию.');

insert into public.departments (id, title, description)
values  (1, 'Управление', 'Самое главное подразделение'),
        (2, 'Отдел управления персоналом', 'Самый добрый отдел'),
        (3, 'Отдел подбора персонала', 'Самый общительный отдел'),
        (4, 'Бухгалтерия', 'Самый щедрый отдел');

insert into public.positions (id, title, description)
values  (1, 'Директор', 'Самый гуманный и солнцеликий управленец'),
        (2, 'Начальник отдела', 'Первый после солнцеликого'),
        (3, 'Главный специалист', 'Ведущий сотрудник любого отдела'),
        (4, 'Специалист', 'Сотрудник любого отдела');

insert into public.organization_structure (id, head_department_id, head_position_id, subordinate_department_id)
values  (1, 1, null, 2),
        (2, 1, null, 3),
        (3, 1, null, 4);

insert into public.users (id, lastname, firstname, middlename, gender, date_of_birth, place_of_birth, position_id, department_id, grade, phone_numbers, work_email, registration_address, residential_address, nationality, insurance_number, taxpayer_number)
values  (1, 'Корепанов', 'Роман', 'Даниилович', 'Мужской', '1988-12-14', 'г. Серпухов', 1, 1, '6', '{"mobile": "+79215511436"}', 'korepanov@company.com', 'Россия, г. Хасавюрт, Заречная ул., д. 24 кв.199', 'Россия, г. Серпухов, Зеленая ул., д. 2 кв.90', 'русский', '34665359207', '298885601004'),
        (2, 'Яппарова', 'Галина', 'Гермоновна', 'Женский', '1993-08-09', 'г. Дербент', 2, 2, '5', '{"mobile": "+79215511436"}', 'yapparova@company.com', 'Россия, г. Рязань, Ленина В.И.ул., д. 6 кв.145', 'Россия, г. Дербент, Комсомольская ул., д. 21 кв.50', 'русский', '52759362623', '154912198705'),
        (3, 'Кучеров', 'Герман', 'Антонович', 'Мужской', '1981-08-22', 'г. Красноярск', 2, 3, '5', '{"mobile": "+79215511436"}', 'kucherov@company.com', 'Россия, г. Реутов, Ленина В.И.ул., д. 23 кв.22', 'Россия, г. Красноярск, Почтовая ул., д. 3 кв.201', 'русский', '64190091782', '581551039113'),
        (4, 'Дарюшина', 'Ольга', 'Герасимовна', 'Женский', '1983-05-13', 'г. Железногорск', 2, 4, '5', '{"mobile": "+79215511436"}', 'daryushina@company.com', 'Россия, г. Казань, Чкалова ул., д. 7 кв.149', 'Россия, г. Железногорск, Речная ул., д. 22 кв.137', 'русский', '87293349859', '110483833670'),
        (5, 'Яловенко', 'Римма', 'Феодосьевна', 'Женский', '1979-04-17', 'г. Нижний Тагил', 3, 4, '4', '{"mobile": "+79215511436"}', 'yalovenko@company.com', 'Россия, г. Ковров, Колхозный пер., д. 11 кв.57', 'Россия, г. Нижний Тагил, Новоселов ул., д. 17 кв.105', 'русский', '48499927384', '660070348269'),
        (6, 'Гроссман', 'Семен', 'Сергеевич', 'Мужской', '1997-02-23', 'г. Новомосковск', 4, 4, '3', '{"mobile": "+79215511436"}', 'grossman@company.com', 'Россия, г. Одинцово, Коммунистическая ул., д. 7 кв.102', 'Россия, г. Новомосковск, Якуба Коласа ул., д. 15 кв.101', 'русский', '26899661063', '472328433371'),
        (7, 'Цельнер', 'Татьяна', 'Аркадивна', 'Женский', '1987-03-09', 'г. Брянск', 4, 2, '3', '{"mobile": "+79215511436"}', 'zelner@company.com', 'Россия, г. Абакан, Юбилейная ул., д. 15 кв.64', 'Россия, г. Брянск, Речная ул., д. 2 кв.3', 'русский', '30049154628', '273169571690'),
        (8, 'Мартюшева', 'Полина', 'Акимовна', 'Женский', '1988-07-21', 'г. Орехово-Зуево', 4, 2, '3', '{"mobile": "+79215511436"}', 'martyusheva@company.com', 'Россия, г. Арзамас, Якуба Коласа ул., д. 3 кв.35', 'Россия, г. Орехово-Зуево, Новый пер., д. 12 кв.43', 'русский', '79294540861', '646764546435'),
        (9, 'Сиянович', 'Мила', 'Павловна', 'Женский', '1998-11-14', 'г. Первоуральск', 4, 3, '3', '{"mobile": "+79215511436"}', 'siyanovich@company.com', 'Россия, г. Муром, Березовая ул., д. 5 кв.201', 'Россия, г. Первоуральск, Заводская ул., д. 5 кв.56', 'русский', '65167909125', '350651114375'),
        (10, 'Грибов', 'Захар', 'Вениаминович', 'Мужской', '1956-09-09', 'г. Норильск', 4, 3, '3', '{"mobile": "+79215511436"}', 'gribov@company.com', 'Россия, г. Салават, Молодежный пер., д. 12 кв.141', 'Россия, г. Норильск, Речная ул., д. 23 кв.82', 'русский', '27281788613', '993152634947');

-- hash for password (pa$$word_) are equal for everyone
insert into public.authorizations (id, user_id, password_hash, role_id)
values  (1, 1, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 1),
        (2, 2, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 2),
        (3, 3, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 3),
        (4, 4, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 4),
        (5, 5, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 4),
        (6, 6, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 4),
        (7, 7, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 2),
        (8, 8, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 2),
        (9, 9, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 3),
        (10, 10, '$2a$12$rrNnco5DWbaFxKMq457MouwkzL/R2XFoe1MrpoX9bN0ms09Zgk6ee', 3);

-- commit the change
commit;