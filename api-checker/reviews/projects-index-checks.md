# Verification for: Projects API Index
Path: /content/en/api/2.projects/1.index.md
Status: [✅] Completed

## 1. GraphQL Schema Verification

### Query/Mutation Operations Mentioned
- [✅] Verify `createProject` mutation exists
  - Operation: `createProject`
  - Location in schema: `/bloo-api/src/schema.graphql:~2298`
  - Actual vs Documented: **MATCHES** - `createProject(input: CreateProjectInput!): Project`

- [✅] Verify `projects` query exists
  - Operation: `projects` (for listing)
  - Location in schema: `/bloo-api/src/schema.graphql`
  - Actual vs Documented: **PARTIALLY MATCHES** - Legacy `projects` query exists but `projectList` is preferred

- [✅] Verify `todoLists` query exists
  - Operation: `todoLists`
  - Location in schema: `/bloo-api/src/schema.graphql`
  - Actual vs Documented: **MATCHES** - `todoLists(projectId: String!): [TodoList!]!`

- [✅] Verify `createTodoList` mutation exists
  - Operation: `createTodoList`
  - Location in schema: `/bloo-api/src/schema.graphql`
  - Actual vs Documented: **MATCHES** - `createTodoList(input: CreateTodoListInput!): TodoList!`

## 2. Enum Verification

### ProjectCategory Enum
- [✅] `ProjectCategory` enum exists
  - Exists in schema: **YES**
  - Location: `/bloo-api/src/generated/types.ts`
  - Values match:
    - [✅] `CRM` - **EXISTS**
    - [✅] `CROSS_FUNCTIONAL` - **EXISTS**
    - [✅] `CUSTOMER_SUCCESS` - **EXISTS**
    - [✅] `DESIGN` - **EXISTS**
    - [✅] `ENGINEERING` - **EXISTS**
    - [✅] `GENERAL` - **EXISTS**
    - [✅] `HR` - **EXISTS**
    - [✅] `IT` - **EXISTS**
    - [✅] `MARKETING` - **EXISTS**
    - [✅] `OPERATIONS` - **EXISTS**
    - [✅] `PRODUCT` - **EXISTS**
    - [✅] `SALES` - **EXISTS**
  - **Missing values in docs**: `PERSONAL`, `PROCUREMENT` (exist in schema but not documented)
  - **Extra values in docs**: None

### Permission Roles
- [✅] Permission roles verification (UserAccessLevel enum)
  - [✅] `OWNER` - **EXISTS**
  - [✅] `ADMIN` - **EXISTS**
  - [✅] `MEMBER` - **EXISTS**
  - [✅] `CLIENT` - **EXISTS**
  - [✅] `VIEW_ONLY` - **EXISTS**
  - [✅] `COMMENT_ONLY` - **EXISTS**

## 3. Input/Response Field Verification

### CreateProject Input Fields
- [✅] `name` field
  - Documented type: `String!`
  - Actual type: **String!**
  - Required status matches: **YES**

- [✅] `companyId` field
  - Documented type: `String!`
  - Actual type: **String!**
  - Required status matches: **YES**

- [✅] `category` field
  - Documented type: `ProjectCategory`
  - Actual type: **ProjectCategory**
  - Required status matches: **YES** (optional in both)

### Projects Query Filter Fields
- [⚠️] `companyIds` field
  - Documented type: `[String!]`
  - Actual type: **Uses `companyId: String` in legacy query, `companyIds: [String!]` in ProjectListFilter**

- [⚠️] `isArchived` field
  - Documented type: `Boolean`
  - Actual type: **Uses `archived: Boolean` in legacy query, `isArchived: Boolean` in ProjectListFilter**

- [✅] `categories` field
  - Documented type: `[ProjectCategory!]`
  - Actual type: **[ProjectCategory!]** (in ProjectListFilter)

### Response Fields Verification
- [✅] Project response fields
  - [✅] `id` - **EXISTS** (ID!)
  - [✅] `name` - **EXISTS** (String!)
  - [✅] `slug` - **EXISTS** (String!)
  - [✅] `category` - **EXISTS** (ProjectCategory!)
  - [✅] `todosCount` - **EXISTS** (Int)
  - [✅] `todosDoneCount` - **EXISTS** (Int)

- [✅] PageInfo fields
  - [✅] `hasNextPage` - **EXISTS** (Boolean!)
  - [✅] `total` - **EXISTS** (Int!)

## 4. Link Verification

### Internal API Links
- [✅] `/api/projects/create-project` - **EXISTS** (2.create-project.md)
- [✅] `/api/projects/list-projects` - **EXISTS** (2.list-projects.md)
- [✅] `/api/projects/delete-project` - **EXISTS** (2.delete-project.md)
- [✅] `/api/projects/archive-project` - **EXISTS** (3.archive-project.md)
- [✅] `/api/projects/rename-project` - **EXISTS** (3.rename-project.md)
- [✅] `/api/projects/copy-project` - **EXISTS** (4.copy-project.md)
- [✅] `/api/projects/lists` - **EXISTS** (5.lists.md)
- [✅] `/api/projects/templates` - **EXISTS** (11.templates.md)
- [✅] `/api/projects/project-activity` - **EXISTS** (3.project-activity.md)

### Related Resource Links
- [✅] `/api/todos` - **REDIRECTS TO** `/api/records` (3.records directory exists)
- [✅] `/api/custom-fields` - **EXISTS** (5.custom fields directory)
- [✅] `/api/automations` - **EXISTS** (6.automations directory)
- [✅] `/api/users` - **EXISTS** (7.user management directory)

## 5. Code Example Verification

### Basic Project Creation Example
- [✅] GraphQL syntax is valid
- [✅] All fields in mutation exist (`name`, `companyId`, `category`)
- [✅] Required fields are included
- [✅] Response structure matches actual response

### Projects Query Example
- [⚠️] Filter parameters exist (`companyIds`, `isArchived`, `categories`) - **USES OLDER SCHEMA** format, should use `projectList` query
- [⚠️] Sort parameter exists (`sort: [updatedAt_DESC]`) - **SHOULD BE** `sort: [ProjectSort!]`
- [✅] Pagination parameters exist (`take`)
- [✅] Response fields exist (`items`, `pageInfo`)

### TodoLists Query Example
- [✅] `todoLists` query exists with `projectId` parameter
- [✅] Response fields exist (`id`, `title`, `position`, `todosCount`)

### CreateTodoList Example
- [✅] `createTodoList` mutation exists
- [✅] Input fields exist (`projectId`, `title`, `position`)

## 6. Error Code Verification

### Documented Error Codes
- [✅] `PROJECT_NOT_FOUND`
  - Exists in codebase: **YES**
  - Location: `/bloo-api/src/lib/errors.ts:145`
  - Message matches: **YES** - ProjectNotFoundError

- [✅] `COMPANY_NOT_FOUND`
  - Exists in codebase: **YES**
  - Location: `/bloo-api/src/lib/errors.ts:79`
  - Message matches: **YES** - CompanyNotFoundError

- [✅] `UNAUTHORIZED`
  - Exists in codebase: **YES** (as FORBIDDEN)
  - Location: `/bloo-api/src/lib/errors.ts:139`
  - Message matches: **CLOSE** - Uses UnauthorizedError (FORBIDDEN code)

- [❌] `PROJECT_NAME_TOO_LONG`
  - Exists in codebase: **NO**
  - Location: Not found as specific error
  - Message matches: **NO** - Uses generic BAD_USER_INPUT with max length validation

## 7. Best Practices Verification

### Documented Limitations
- [✅] "Keep names under 50 characters" - **VERIFIED** - Schema has max(50) validation
- [⚠️] "Limit the number of lists per project (max 50)" - **NEEDS VERIFICATION** - Not confirmed in schema search
- [✅] Pagination recommendations - **REASONABLE** - Default take: 20, customizable

## 8. Special Considerations

### Project Structure Claims
- [✅] "Projects belong to companies" - **VERIFIED** - Project type has `company: Company!` field
- [✅] "Each project can have multiple lists" - **VERIFIED** - `todoLists(projectId: String!)` query exists
- [✅] "Lists contain todos" - **VERIFIED** - TodoList type has todos relationship
- [✅] "Projects support custom fields, tags, and automations" - **VERIFIED** - Project type includes these fields

## Summary

### Critical Issues (Must Fix)
1. **Query naming inconsistency**: Documentation uses `projects` query but should recommend `projectList` for new implementations
2. **Missing ProjectCategory values**: Documentation missing `PERSONAL` and `PROCUREMENT` enum values that exist in schema
3. **Missing error code**: `PROJECT_NAME_TOO_LONG` doesn't exist as specific error code (uses generic validation error)

### Minor Issues (Should Fix)
1. **Parameter naming**: Documentation shows `isArchived` but legacy query uses `archived` 
2. **Sort syntax**: Documentation shows `updatedAt_DESC` but should use `ProjectSort` enum format
3. **Link references**: Documentation references `/api/todos` but should reference `/api/records`

### Suggestions
1. **Update examples**: Use `projectList` query instead of deprecated `projects` query
2. **Complete enum documentation**: Add missing `PERSONAL` and `PROCUREMENT` categories
3. **Error handling**: Update error documentation to reflect actual validation error codes
4. **Query migration**: Add note about legacy vs current query formats

### Overall Assessment
The documentation is **largely accurate** with most operations, types, and fields correctly documented. The main issues are around using older API patterns and missing some newer enum values. The core functionality is properly documented.