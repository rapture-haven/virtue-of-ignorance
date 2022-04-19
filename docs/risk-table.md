# Таблица рисков

* [Изменение протоколов передачи аудио/видео](#изменение-протоколов-передачи-аудиовидео)
* [Утечки данных](#утечки-данных)
* [Изменение АПИ используемых библиотек](#изменение-апи-используемых-библиотек)
* [Блокировка доступа](#блокировка-доступа)
* [Вредоносный код в Open source решениях](#вредоносный-код-в-open-source-решениях)

## Изменение протоколов передачи аудио/видео
### Приоритет: 
- 2

### Риск:
- Изменение протоколов передачи аудио/видео, поддерживаемых браузерами,
используемый или поддерживаемый в приложении.

### Состояние:
- «Мониторинг»

### Вероятность:
- «Низкая»

### Урон:
- «Высокий»

### Воздействие: 
- «Среднее»

### Тип стратегии:
- «Cнижение» 

### Стратегия:
- В код ПО заложить возможность смены протоколов

### Ответственный:
- [Козлов Максим](https://t.me/kbsjvxvdcshj)


## Утечки данных
### Приоритет: 
- 5

### Риск: 
- Утечка данных пользователей продукта.

### Состояние: 
- «Анализ»

### Вероятность:
- «Низкая»

### Урон:
- «Средний»

### Воздействие: 
- «Среднее»

### Тип стратегии:
- «Принятие»

### Стратегия:
- на 09.04.2022 приложение не использует никакую конфиденциальную 
информацию. При их появлении необходимо пересмотреть стратегию. 

### Ответственный:
- [Козлов Максим](https://t.me/kbsjvxvdcshj)

## Изменение АПИ используемых библиотек
### Приоритет: 
- 5

### Риск: 
- Используемые в приложении библиотеки могут изменить свой АПИ, н-р,
библиотека webrtc.

### Состояние: 
- «Закрыт»

### Вероятность:
- «Средняя»

### Урон:
- «Средний»

### Воздействие: 
- «Среднее»

### Тип стратегии:
- «Мониторинг»

### Стратегия:
- Анализ изменений и обратной совместимости с текущей версией;
- Привязка к конкретной версии библиотеки при breaking change изменениях;
 
### Ответственный:
- [Хетагуров Павел](https://t.me/e1e233)


## Блокировка доступа
### Приоритет: 
- 2

### Риск: 
Блокировка доступа к ресурсам, напр.
- репозитории пакетов npm/go
- GitHub
- ssh

### Состояние: 
- «Мониторинг»

### Вероятность:
- «Низкая»

### Урон:
- «Высокий»

### Воздействие: 
- «Среднее»

### Тип стратегии:
- «Снижение» 

### Стратегия:
- Подготовка списка зеркал для используемых пакетов/библиотек
- Использование дублирующего репозитория в GitLab

### Ответственный:
- [Мышонков Максим](https://t.me/maxdbd)

## Вредоносный код в Open source решениях
### Приоритет: 
- 2

### Риск: 
- Вредоносный код в  Open source решениях (малвари, шифровальщик и т.п.)

### Состояние:
- «Мониторинг»

### Вероятность:
- «Средняя»

### Урон:
- «Высокий»

### Воздействие: 
- «Высокое»

### Тип стратегии:
- «Уклонение»

### Стратегия:
- Учёт OpenSource продуктов, в которых возможны малвари, шифровальщики и т.п. [Ссылка на список опасных продуктов](https://docs.google.com/spreadsheets/d/1H3xPB4PgWeFcHjZ7NOPtrcya_Ua4jUolWm-7z9-jSpQ/htmlview?usp=sharing&pru=AAABf7qkysY*4LAvszFh2bLcVv_Pnp4xuA#);
- Использовать версии до появления вредоносного ПО;
- Возможно переход на другой аналогичный продукт;

### Ответственный:
- [Хетагуров Павел](https://t.me/e1e233)