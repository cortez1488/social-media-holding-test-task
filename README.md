# social-media-holding-test-task
social media holding junior test task

Схема бд: ![Снимок экрана от 2022-04-28 21-02-32](https://user-images.githubusercontent.com/93537782/165795483-a078fc6d-b571-4b4d-9500-1cd1eaed3194.png)


Для запуска:
В директории проекта

docker-compose build && docker-compose up


При ошибках:
 - can't stat .... Прописать sudo chown -R $USER <путь к папке>
 - could not open file "global/pg_filenode.map": Permission denied  .... Запусить уже построенный docker-compose еще раз
