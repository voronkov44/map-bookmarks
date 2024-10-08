# **Создание закладок в go через функцию map**

## **Описание**

Описание программы "Менеджер закладок"

Данная программа представляет собой простое консольное приложение для управления закладками. Пользователь может выполнять следующие действия в рамках приложения:

1. Посмотреть закладки - эта опция позволяет пользователю просматривать все сохраненные закладки. Если закладок нет, будет выведено соответствующее сообщение.

2. Добавить закладку - пользователь может ввести название и адрес новой закладки. Закладка будет сохранена в памяти приложения.

3. Удалить закладку - если пользователь желает удалить закладку, он может ввести название закладки, которую хочет удалить. Программа удалит указанную закладку из списка.

4. Выход - завершение работы приложения.

### Как работает приложение:

- Программа использует цикл, который отображает меню и позволяет пользователю выбирать опции.
- Закладки хранятся в виде карты (map) с ключами (название закладки) и значениями (адреса закладок).
- Все функции, такие как просмотр, добавление и удаление закладок, реализованы как отдельные функции для удобства и структуры кода.
  
Программа проста в использовании и подходит для пользователей, желающих легко управлять своими закладками. Чтобы начать работу, просто запустите программу и следуйте инструкциям меню.

## **Установка**

Для клонирования репозитория необходимо перейти в любую удобную директорию и выполнить команду в терминале:

```no-highlight
https://github.com/voronkov44/map-bookmarks.git
```


## **Использование**

Для запуска проекта необходимо выполнить следующие шаги:

**1.** Открываем проект в любом удобном IDE, я буду открывать проект в GoLand от компании JetBrains.

![image](https://github.com/user-attachments/assets/3ea65068-e9ef-44f7-9289-08e857e526a3)


Для установки IDE GoLand смотрите [зависимости.](https://github.com/voronkov44/map-bookmarks/tree/readme_branch?tab=readme-ov-file#%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8) 

**2.** Для запуска программы нужно ввести следующую команду в терминале вашего IDE.
```no-highlight
go run main.go
```

**3.** Пример работы программы.

![image](https://github.com/user-attachments/assets/86a095d8-dff5-4d0f-a418-bb078d49eae2)



## **Зависимости**

Установка IDE [GoLang.](https://www.jetbrains.com/go/)
