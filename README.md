# social-media-holding-test-task
social media holding junior test task

Схема бд: ![Снимок экрана от 2022-04-28 21-02-32](https://user-images.githubusercontent.com/93537782/165795483-a078fc6d-b571-4b4d-9500-1cd1eaed3194.png)


Endpoints:
 GET -/api/get_users Информация по всем пользователям
 GET -/api/get_user?id=[integer] Информация по 1 пользователю (по ID в БД)
 GET -/api/get_history_by_tg?chatid=[integer64] Информация запросов по 1 пользователю (по chatID в БД)
 DELETE -/api/delete_ip?ip=[string] Удаление всех записей с этим ip

Для запуска:
В директории проекта

docker-compose build && docker-compose up


При ошибках:
 - can't stat .... Прописать sudo chown -R $USER <путь к папке>
 - could not open file "global/pg_filenode.map": Permission denied  .... Запусить уже построенный docker-compose еще раз
