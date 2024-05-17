# System Design - Социальная сеть для путешественников

## Функциональные требования
- Публикация постов из путешествий с фотографиями, небольшим описанием и привязкой к конкретному месту путешествия.

- Оценка и комментарии постов других путешественников.

- Подписка на других путешественников, чтобы следить за их активностью.

- Поиск популярных мест для путешествий и просмотр постов с этих мест в виде ТОПа мест по странам и городам.

- Общение с другими путешественниками в личных сообщениях.

- Просмотр ленты других путешественников.

## Нефункциональные требования
- DAU 10 000 000.

- Доступность 99,95%.

- Запуск только на территории СНГ.

- Данные храним всегда.

- Ограничения:
    - Аккаунт
        - Максимальное количество фотографий в аккаунте - 1.
        - Максимальное количество символов в поле ФИО - 80.
        - Максимальное количество символов в описании аккаунта - 200.
    - Пост
        - Максимальное количество фотографий в посте - 5.
        - Максимальное количество символов в описании поста - 2000.
        - Максимальный размер фотографии - 1МB.
    - Комментарии
        - Максимальное количество символов в комментарии - 500.
        - Количество комментариев под постом не ограничено.
    - Личные сообщения
        - Максимальное количество символов в одном сообщении - 500.
    - Подписки
        - Количество подписок не ограничено.

- Ожидается, что пользователь будет:
    - В день
        - Просматривать 100 постов.
        - Просматривать информацию о 5-ти аккаунтах.
        - Оставлять 30 оценок.
        - Оставлять 10 комментариев.
        - Отправлять 80 личных текстовых сообщений и 20 фото.
        - Осущетсвлять поиск 5-ти мест.
    - В неделю
        - Публиковать 2 поста.
        - Подписываться на 1 человека.

- Тайминги:
    - Максимальная задержка на публикацию поста - 2сек.
    - Максимальная задержка на загрузку ленты - 1сек.
    - Допускается, что опубликованные посты отображаются в лентах подписчиков с небольшой задержкой.

## Расчёты
### Посты
1. Фотографии на 1 пост: 1MB * 5 = 5MB
```
    RPS = 10 000 000 * 0,29 / 86400 = 34.
    Трафик в секунду = 34 * 5MB = 170MB/s.
    Трафик в день = 170MB/s * 86400 = 14TB.
```

Общий объём данных в год: 14TB * 365 = 5PB. Будем складывать в S3 хранилище.

2. Пост: ID(8B) + описание(2000\*2B=4КB) + дата(8B) + локация(140B) + ссылки на фото(100\*2B*5=1КБ) = 5KB.
```
    RPS = 10 000 000 * 0,29 / 86400 = 34.
    Трафик в секунду = 34 * 5KB = 170KB/s.
    Трафик в день = 170KB/s * 86400 = 14GB.
```

3. Комментарий: ID поста(8B) + ID пользователя(8B) + текст(500*2B=1KB) + дата(8B) = 1КB.
```
    RPS = 10 000 000 * 10 / 86400 = 1158.
    Трафик в секунду = 1158 * 1KB = 1,2KB/s.
    Трафик в день = 1,2KB/s * 86400 = 104GB.
```

Общий объём данных в год: (17GB + 104GB) * 365 = 44TB.

Необходимое количество дисков (HDD по 6TB) на 1 год: по размеру (44TB/6TB=8) и пропускной способности (100MB/s) достаточно 8-и дисков, но для поддержания рассчитанного RPS (1159RPS/100IOPS=12) необходимы ещё 4 диска. Итого 12 HDD по 6TB.

## Дизайн
Дизайн системы представлен в нотации С4. Построены диаграммы для двух уровней абстракции: context и container.

#### Уровень 1. System context diagram
![context-diagram](architecture/diagram.png?raw=true)

#### Уровень 2. Media system container diagram
![media-diagram](architecture/media_system/diagram.png?raw=true)

#### Уровень 2. Messenger system container diagram
![messenger-diagram](architecture/messenger_system/diagram.png?raw=true)

#### Уровень 2. Post system container diagram
![post-diagram](architecture/post_system/diagram.png?raw=true)

#### Уровень 2. Feed system container diagram
![feed-diagram](architecture/feed_system/diagram.png?raw=true)