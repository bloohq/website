# API Documentation Standards

Each API documentation page should follow this consistent structure to ensure clarity and completeness:

## Required Sections for API Pages

1. **Title and Description** (frontmatter)
   - Clear, action-oriented title (e.g., "Create a Project", "List all Projects")
   - Concise description of what the endpoint does

2. **Overview/Introduction**
   - Brief explanation of the endpoint's purpose
   - When and why to use it

3. **Basic Example**
   - Simple, minimal GraphQL query/mutation
   - Shows only required parameters
   - Easy to copy and test

4. **Advanced Example** (if applicable)
   - Complex example with all optional parameters
   - Real-world use case
   - Shows nested configurations

5. **Parameter Tables**
   - **Input Parameters**: All available input fields with:
     - Parameter name
     - Type (with GraphQL notation)
     - Required (✅ Yes / No)
     - Description
   - **Nested Types**: Separate tables for complex input types
   - **Enum Values**: Tables listing all possible values with descriptions

6. **Response Fields**
   - Table showing all fields in the response
   - Field types and descriptions
   - Note which fields are always present vs optional

7. **Permissions/Authorization**
   - Required company-level roles
   - Required project-level roles (if applicable)
   - Clear table showing which roles can/cannot perform the action

8. **Error Responses**
   - Common error scenarios with example JSON
   - Error codes and their meanings
   - How to handle each error type

9. **Important Notes/Considerations**
   - Performance implications
   - Limitations (e.g., max records, rate limits)
   - Best practices
   - Related endpoints

## Example Structure Template

```markdown
---
title: [Action] [Resource]
description: [What this endpoint does]
---

## [Action] a [Resource]

[Brief explanation of what this endpoint does and when to use it]

### Basic Example

```graphql
[Minimal working example]
```

### Advanced Example

```graphql
[Complex example with all options]
```

## Input Parameters

### [InputType]

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `field1` | String! | ✅ Yes | Description of field1 |
| `field2` | Boolean | No | Description of field2 |

### [EnumType] Values

| Value | Description |
|-------|-------------|
| `VALUE1` | What VALUE1 means |
| `VALUE2` | What VALUE2 means |

## Response Fields

| Field | Type | Description |
|-------|------|-------------|
| `id` | String! | Unique identifier |
| `success` | Boolean! | Operation success status |

## Required Permissions

[Explain permission requirements]

| Role | Can Perform Action |
|------|-------------------|
| `OWNER` | ✅ Yes |
| `ADMIN` | ✅ Yes |
| `MEMBER` | ❌ No |

## Error Responses

### [Error Type]
```json
{
  "errors": [{
    "message": "Error message",
    "extensions": {
      "code": "ERROR_CODE"
    }
  }]
}
```
``` 