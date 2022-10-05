# Practise_1-Big_Numbers

Обрання варіантів у додатку здійснюється за рахунок натискання потрібної кнопки (цифра, або ‘*’) та натисканні клавіші Enter

Для завдань 1, 2, 3 використовується згенерована хеш-таблиця з кількістю можливих ключів. Генерується вона один раз при обранні якогось з варіантів завдання (задля збільшення ефективності програми). 

Короткий опис функцій:<br />
```keyHandler()``` – функція для обробки кнопок у головному меню.<br />
```ClearConsole()``` – очистка консолі (працює для ОС типу windows, linux, Darwin, в іншому випадку просто відступає 3 рядка від попереднього виводу)<br />
```fillKeyCount()``` – заповнює хеш-таблицю з кількістю варіантів ключів по n-бітної послідовності<br />
```task1()``` – виконання першого завдання<br />
```task2()``` – виконання другого завдання<br />
```bruteForce()``` – шукання значення, згенерованого випадково<br />
```task3()``` – виконання третього завдання<br />

Для запуску програми запустіть main.exe файл. 

Значення 0х00...0 і 0хFF...F - мінімальні і максимальні числа, представлені у 16-річній системі. Кількість знаків 0 та F залежать від бітової послідовності<br/>
0хFF - 255, 8 біт<br/>
0xFFFF - 65535, 16 біт<br/>
і т.д.
