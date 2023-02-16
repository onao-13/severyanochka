## api/v1/page/preview-main

<blockquote>
<h3>Метод: POST</h3>
</blockquote>

## Описание

Возвращает продукты, которые отображаются на главной странице. Содержит категории: <br>
<ul>
<li>Акции (stock)</li>
<li>Новинки (newProducts)</li>
<li>Покупали раньше (buyingBefore)</li>
<li>Специальные предложения (specialOffers)</li>
<li>Статьи (articles)</li>
</ul>

<blockquote>
<h4>❗ Примечание: <br>
Категория покупали раньше отображается только если пользователь в системе</h4>
</blockquote>

## Формат ответа

```json
{
  "data": {
    "newProducts": [
      {
        "name": string,
        "imageUrl": string,
        "price": float,
        "sale": int,
        "rating": int,
        "county": string
      }
    ],
    "stock": [
      {
        "name": string,
        "imageUrl": string,
        "price": float,
        "sale": int,
        "rating": int,
        "county": string
      }
    ]
  },
  "result": string,
  "status": int
}
```

### 1. Data
**newProducts**: Список новых продуктов из категории "Новинки" на главной. Размер: 20 штук.
- *name: название продукта
- *imageUrl*: ссылка на картинку продукта
- *price*: цена продукта
- *sale*: скидка на продукт
- rating*: количество звезд у продукта
- *county*: страна производитель продукта

**stock**: Акционные продукты из категории "Акции" на главной. Размер: 20 штук.
- *name*: название продукта
- *imageUrl*: ссылка на картинку продукта
- *price*: цена продукта
- *sale*: скидка на продукт
- *rating*: количество звезд у продукта
- *county*: страна производитель продукта

### 2. Result
Информация о том, как завершился запрос.

### 3. Status
Статус код запроса.

<blockquote>
<h4>❗ Примечание по статус-кодам:</h4>
<ul>
    <li>200: все прошло хорошо;</li>
    <li>201: создан новый объект;</li>
    <li>400: неправильный запрос;</li>
    <li>403: нет доступа;</li>
    <li>404: не найдено;</li>
    <li>500: ошибка сервера.</li>
</ul>
</blockquote>

