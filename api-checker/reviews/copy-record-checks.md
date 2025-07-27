# Verification for: Copy Record
Path: /content/en/api/3.records/8.copy-record.md
Status: [🔄] In Progress

## 1. GraphQL Schema Verification

### Mutation Operation
- [✅] Verify `copyTodo` mutation exists
  - Operation: `copyTodo`
  - Location in schema: `/bloo-api/src/schema.graphql`
  - Actual vs Documented: **MOSTLY MATCHES** - `copyTodo(input: CopyTodoInput!): Boolean!`

## 2. Input Parameter Verification

### CopyTodo Input Fields
- [❌] `title` field
  - Documented type: `String!` (Required: Yes)
  - Actual type: **String** (OPTIONAL)
  - Required status matches: **NO** - Documentation wrong

- [✅] `todoId` field
  - Documented type: `String!`
  - Actual type: **String!**
  - Required status matches: **YES**

- [✅] `todoListId` field
  - Documented type: `String!`
  - Actual type: **String!**
  - Required status matches: **YES**

- [✅] `options` field
  - Documented type: `Array!`
  - Actual type: **[CopyTodoOption!]!**
  - Required status matches: **YES**

## 3. Copy Options Enum Verification

### CopyTodoOptions Enum
- [✅] `DESCRIPTION` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `DUE_DATE` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `CHECKLISTS` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `ASSIGNEES` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `TAGS` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `CUSTOM_FIELDS` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [⚠️] **Missing from docs**: `COMMENTS` option exists in schema but not documented

## 4. Response Field Verification

### CopyTodo Response
- [❌] Response structure mismatch
  - Documented: `{ success: Boolean }`
  - Actual type: **Boolean!** (direct boolean, not object)
  - Documentation shows wrong response format

## 5. Implementation Verification

### Resolver Check
- [✅] Resolver exists for `copyTodo`
  - Location: `/bloo-api/src/resolvers/Mutation/copyTodo.ts`
  - Handler function: `copyTodo`

### Business Logic Verification
- [✅] Copy functionality implemented correctly
  - Options handling: **EXCELLENT** - all options work as documented
  - Positioning logic: **CORRECT** - places at bottom of list
  - Implementation: `/bloo-api/src/datasources/TodoDatasource.ts`

## 6. Permission Verification

### Required Permissions
- [✅] Edit permissions on source list
  - Permission check exists: **YES** - OWNER/ADMIN/MEMBER required
  - Location: `/bloo-api/src/permissions/permissions.ts`

- [✅] Edit permissions on target list
  - Permission check exists: **YES** - cross-project restrictions for MEMBER
  - Location: `copyTodo.ts` resolver

- [✅] **Additional restrictions found**:
  - MEMBER role can only copy within same project
  - Assignees filtered by target project membership

## 7. Error Code Verification

### Documented Error Codes
- [❌] `BAD_USER_INPUT`
  - Exists in codebase: **YES** but not used for invalid IDs
  - Used for invalid IDs: **NO** - uses specific errors instead
  - Actual errors: `TODO_NOT_FOUND`, `TODO_LIST_NOT_FOUND`

- [✅] `FORBIDDEN`
  - Exists in codebase: **YES**
  - Used for permission errors: **YES** (UnauthorizedError)
  - Location: `/bloo-api/src/lib/errors.ts:139`

- [❌] `GRAPHQL_VALIDATION_FAILED`
  - Exists in codebase: **NOT FOUND**
  - Used for missing fields: **NO** - GraphQL handles validation
  - This error code doesn't exist in the codebase

## 8. Headers Verification

### Required Headers
- [🔄] `x-bloo-token-id` header requirement
  - Actually required: [checking...]
  - Part of authentication: [checking...]

- [🔄] `x-bloo-token-secret` header requirement
  - Actually required: [checking...]
  - Part of authentication: [checking...]

- [🔄] `x-bloo-project-id` header requirement
  - Actually required: [checking...]
  - Used for project context: [checking...]

- [🔄] `x-bloo-company-id` header requirement
  - Actually required: [checking...]
  - Used for company context: [checking...]

## 9. Link Verification

### Internal API Links
- [🔄] `/api/records/move-record-list` - [checking if file exists...]
- [🔄] `/api/error-codes` - [checking if file exists...]

## 10. Code Example Verification

### Basic Example
- [🔄] GraphQL syntax is valid
- [🔄] All fields in mutation exist
- [🔄] Required fields are included
- [🔄] Response structure matches actual response

## 11. Special Considerations

### Business Rules
- [🔄] "Copied record placed at bottom of target list" - verify default positioning
- [🔄] Copy options actually copy the specified data elements
- [🔄] Cross-list copying is supported

## Summary

### Critical Issues (Must Fix)
[To be populated after verification]

### Minor Issues (Should Fix)
[To be populated after verification]

### Suggestions
[To be populated after verification]