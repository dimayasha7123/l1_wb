### Вопрос

Чем отличаются буферизированные и не буферизированные каналы?

### Ответ

Как ясно из названия, буферизированные каналы имеют ненулевой буфер, который задается при создании канала.
Таким образом, небуферизированные каналы блокируют отправку данных, пока их не готовы прочитать.
Буферизированные же не блокируют отправку, пока буфер не заполнен.