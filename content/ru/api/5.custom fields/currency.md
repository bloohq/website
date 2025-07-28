---
title: Пользовательское поле валюты
description: Создайте поля валюты для отслеживания денежных значений с правильным форматированием и валидацией
---

Пользовательские поля валюты позволяют вам хранить и управлять денежными значениями с соответствующими кодами валют. Поле поддерживает 72 разные валюты, включая основные фиатные валюты и криптовалюты, с автоматическим форматированием и необязательными ограничениями на минимальные/максимальные значения.

## Основной пример

Создайте простое поле валюты:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Расширенный пример

Создайте поле валюты с ограничениями валидации:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Параметры ввода

### CreateCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Да | Отображаемое имя поля валюты |
| `type` | CustomFieldType! | ✅ Да | Должно быть `CURRENCY` |
| `currency` | String | Нет | Код валюты по умолчанию (3-буквенный ISO код) |
| `min` | Float | Нет | Минимально допустимое значение (хранится, но не применяется при обновлениях) |
| `max` | Float | Нет | Максимально допустимое значение (хранится, но не применяется при обновлениях) |
| `description` | String | Нет | Текст помощи, отображаемый пользователям |

**Примечание**: Контекст проекта автоматически определяется по вашей аутентификации. Вы должны иметь доступ к проекту, в котором создаете поле.

## Установка значений валюты

Чтобы установить или обновить значение валюты в записи:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### SetTodoCustomFieldInput Параметры

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Да | ID записи для обновления |
| `customFieldId` | String! | ✅ Да | ID пользовательского поля валюты |
| `number` | Float! | ✅ Да | Денежная сумма |
| `currency` | String! | ✅ Да | 3-буквенный код валюты |

## Создание записей с валютными значениями

При создании новой записи с валютными значениями:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      number
      currency
    }
  }
}
```

### Формат ввода для создания

При создании записей валютные значения передаются по-другому:

| Параметр | Тип | Описание |
|-----------|------|-------------|
| `customFieldId` | String! | ID поля валюты |
| `value` | String! | Сумма в виде строки (например, "1500.50") |
| `currency` | String! | 3-буквенный код валюты |

## Поддерживаемые валюты

Blue поддерживает 72 валюты, включая 70 фиатных валют и 2 криптовалюты:

### Фиатные валюты

#### Америка
| Валюта | Код | Название |
|----------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Венесуэльский боливар суверенный |
| Боливийский боливиано | `BOB` | Боливийский боливиано |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Европа
| Валюта | Код | Название |
|----------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Норвежская крона | `NOK` | Норвежская крона |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Азиатско-Тихоокеанский регион
| Валюта | Код | Название |
|----------|------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |

#### Ближний Восток и Африка
| Валюта | Код | Название |
|----------|------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Криптовалюты
| Валюта | Код |
|----------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Поля ответа

### TodoCustomField Ответ

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | String! | Уникальный идентификатор для значения поля |
| `customField` | CustomField! | Определение пользовательского поля |
| `number` | Float | Денежная сумма |
| `currency` | String | 3-буквенный код валюты |
| `todo` | Todo! | Запись, к которой принадлежит это значение |
| `createdAt` | DateTime! | Когда значение было создано |
| `updatedAt` | DateTime! | Когда значение было в последний раз изменено |

## Форматирование валюты

Система автоматически форматирует валютные значения в зависимости от локали:

- **Расположение символа**: Правильное размещение символов валюты (до/после)
- **Десятичные разделители**: Использует специфические для локали разделители (. или ,)
- **Тысячные разделители**: Применяет соответствующую группировку
- **Десятичные знаки**: Показывает 0-2 десятичных знака в зависимости от суммы
- **Специальная обработка**: USD/CAD показывают префикс кода валюты для ясности

### Примеры форматирования

| Значение | Валюта | Отображение |
|-------|----------|---------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Валидация

### Валидация суммы
- Должна быть допустимым числом
- Ограничения на минимальные/максимальные значения хранятся с определением поля, но не применяются при обновлении значений
- Поддерживает до 2 десятичных знаков для отображения (полная точность хранится внутри)

### Валидация кода валюты
- Должен быть одним из 72 поддерживаемых кодов валют
- Чувствителен к регистру (используйте заглавные буквы)
- Неверные коды возвращают ошибку

## Функции интеграции

### Формулы
Поля валюты могут использоваться в пользовательских полях ФОРМУЛЫ для расчетов:
- Суммировать несколько полей валюты
- Рассчитывать проценты
- Выполнять арифметические операции

### Конверсия валюты
Используйте поля CURRENCY_CONVERSION для автоматической конверсии между валютами (см. [Поля конверсии валюты](/api/custom-fields/currency-conversion))

### Автоматизации
Значения валюты могут вызывать автоматизации на основе:
- Порогов суммы
- Типа валюты
- Изменений значений

## Необходимые разрешения

| Действие | Необходимое разрешение |
|--------|-------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Примечание**: Хотя любой член проекта может создавать пользовательские поля, возможность устанавливать значения зависит от разрешений на основе ролей, настроенных для каждого поля.

## Ошибки ответов

### Неверное значение валюты
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

Эта ошибка возникает, когда:
- Код валюты не является одним из 72 поддерживаемых кодов
- Формат числа недопустим
- Значение не может быть правильно разобрано

### Пользовательское поле не найдено
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Рекомендации

### Выбор валюты
- Установите валюту по умолчанию, которая соответствует вашему основному рынку
- Последовательно используйте коды валют ISO 4217
- Учитывайте местоположение пользователя при выборе значений по умолчанию

### Ограничения значений
- Установите разумные минимальные/максимальные значения, чтобы предотвратить ошибки ввода данных
- Используйте 0 в качестве минимума для полей, которые не должны быть отрицательными
- Учитывайте ваш случай использования при установке максимальных значений

### Многовалютные проекты
- Используйте последовательную базовую валюту для отчетности
- Реализуйте поля CURRENCY_CONVERSION для автоматической конверсии
- Документируйте, какая валюта должна использоваться для каждого поля

## Распространенные случаи использования

1. **Бюджетирование проекта**
   - Отслеживание бюджета проекта
   - Оценки затрат
   - Отслеживание расходов

2. **Продажи и сделки**
   - Значения сделок
   - Суммы контрактов
   - Отслеживание доходов

3. **Финансовое планирование**
   - Суммы инвестиций
   - Раунды финансирования
   - Финансовые цели

4. **Международный бизнес**
   - Многовалютное ценообразование
   - Отслеживание валютного обмена
   - Трансакции через границу

## Ограничения

- Максимум 2 десятичных знака для отображения (хотя более высокая точность хранится)
- Нет встроенной конверсии валюты в стандартных полях CURRENCY
- Нельзя смешивать валюты в одном значении поля
- Нет автоматических обновлений курсов обмена (используйте CURRENCY_CONVERSION для этого)
- Символы валюты не настраиваемые

## Связанные ресурсы

- [Поля конверсии валюты](/api/custom-fields/currency-conversion) - Автоматическая конверсия валют
- [Числовые поля](/api/custom-fields/number) - Для неналоговых числовых значений
- [Поля формул](/api/custom-fields/formula) - Расчеты с валютными значениями
- [Пользовательские поля списка](/api/custom-fields/list-custom-fields) - Запрос всех пользовательских полей в проекте