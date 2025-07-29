---
title: 货币自定义字段
description: 创建货币字段以跟踪带有适当格式和验证的货币值
---

货币自定义字段允许您存储和管理带有相关货币代码的货币值。该字段支持72种不同的货币，包括主要法定货币和加密货币，并具有自动格式化和可选的最小/最大约束。

## 基本示例

创建一个简单的货币字段：

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

## 高级示例

创建一个带有验证约束的货币字段：

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

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 货币字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `CURRENCY` |
| `currency` | String | 否 | 默认货币代码（3字母ISO代码） |
| `min` | Float | 否 | 允许的最小值（存储但在更新时不强制执行） |
| `max` | Float | 否 | 允许的最大值（存储但在更新时不强制执行） |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：项目上下文会根据您的身份验证自动确定。您必须有权访问您正在创建字段的项目。

## 设置货币值

要在记录上设置或更新货币值：

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

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的ID |
| `customFieldId` | String! | ✅ 是 | 货币自定义字段的ID |
| `number` | Float! | ✅ 是 | 货币金额 |
| `currency` | String! | ✅ 是 | 3字母货币代码 |

## 创建带有货币值的记录

在创建带有货币值的新记录时：

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

### 创建的输入格式

在创建记录时，货币值的传递方式不同：

| 参数 | 类型 | 描述 |
|------|------|------|
| `customFieldId` | String! | 货币字段的ID |
| `value` | String! | 作为字符串的金额（例如，“1500.50”） |
| `currency` | String! | 3字母货币代码 |

## 支持的货币

Blue支持72种货币，包括70种法定货币和2种加密货币：

### 法定货币

#### 美洲
| 货币 | 代码 | 名称 |
|------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | 委内瑞拉主权玻利瓦尔 |
| 玻利维亚玻利维亚诺 | `BOB` | 玻利维亚玻利维亚诺 |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### 欧洲
| 货币 | 代码 | 名称 |
|------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| 挪威克朗 | `NOK` | 挪威克朗 |
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

#### 亚太地区
| 货币 | 代码 | 名称 |
|------|------|------|
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
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### 中东和非洲
| 货币 | 代码 | 名称 |
|------|------|------|
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

### 加密货币
| 货币 | 代码 |
|------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `number` | Float | 货币金额 |
| `currency` | String | 3字母货币代码 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## 货币格式化

系统会根据区域设置自动格式化货币值：

- **符号位置**：正确放置货币符号（前/后）
- **小数分隔符**：使用区域特定的分隔符（. 或 ,）
- **千位分隔符**：应用适当的分组
- **小数位数**：根据金额显示0-2位小数
- **特殊处理**：美元/加元显示货币代码前缀以提高清晰度

### 格式化示例

| 值 | 货币 | 显示 |
|----|------|------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## 验证

### 金额验证
- 必须是有效数字
- 最小/最大约束与字段定义一起存储，但在值更新期间不强制执行
- 支持最多2位小数的显示（内部存储完整精度）

### 货币代码验证
- 必须是72个支持的货币代码之一
- 区分大小写（使用大写字母）
- 无效代码会返回错误

## 集成功能

### 公式
货币字段可以在公式自定义字段中用于计算：
- 求和多个货币字段
- 计算百分比
- 执行算术运算

### 货币转换
使用货币转换字段自动在货币之间转换（请参见 [货币转换字段](/api/custom-fields/currency-conversion)）

### 自动化
货币值可以根据以下条件触发自动化：
- 金额阈值
- 货币类型
- 值变化

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**注意**：虽然任何项目成员都可以创建自定义字段，但设置值的能力取决于为每个字段配置的基于角色的权限。

## 错误响应

### 无效的货币值
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

当发生以下情况时，会出现此错误：
- 货币代码不是72个支持的代码之一
- 数字格式无效
- 值无法正确解析

### 找不到自定义字段
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

## 最佳实践

### 货币选择
- 设置与您的主要市场相匹配的默认货币
- 一致使用ISO 4217货币代码
- 在选择默认值时考虑用户位置

### 值约束
- 设置合理的最小/最大值以防止数据输入错误
- 对于不应为负数的字段，将0作为最小值
- 在设置最大值时考虑您的用例

### 多货币项目
- 使用一致的基础货币进行报告
- 实施货币转换字段以实现自动转换
- 文档说明每个字段应使用的货币

## 常见用例

1. **项目预算**
   - 项目预算跟踪
   - 成本估算
   - 费用跟踪

2. **销售与交易**
   - 交易金额
   - 合同金额
   - 收入跟踪

3. **财务规划**
   - 投资金额
   - 融资轮次
   - 财务目标

4. **国际业务**
   - 多货币定价
   - 外汇跟踪
   - 跨境交易

## 限制

- 显示最多2位小数（尽管存储更多精度）
- 标准货币字段中没有内置货币转换
- 不能在单个字段值中混合货币
- 没有自动汇率更新（请使用货币转换）
- 货币符号不可自定义

## 相关资源

- [货币转换字段](/api/custom-fields/currency-conversion) - 自动货币转换
- [数字字段](/api/custom-fields/number) - 用于非货币数值
- [公式字段](/api/custom-fields/formula) - 使用货币值进行计算
- [列表自定义字段](/api/custom-fields/list-custom-fields) - 查询项目中的所有自定义字段