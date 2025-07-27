# Verification for: 4.copy-project.md
Path: /content/en/api/2.projects/4.copy-project.md
Status: [ ] In Progress / [✓] Completed - MAJOR FIXES APPLIED

**MAJOR DISCREPANCIES FOUND:**
- CopyProjectStatus response schema is completely wrong
- Missing `coverConfig` option from documentation
- No logical dependencies validation found in code
- Missing actual copy worker implementation

## 1. GraphQL Schema Verification

### Mutation Name
- [ ] Verify the GraphQL operation name exists in schema
  - Operation: `copyProject`
  - Location in schema: [file:line]
  - Actual vs Documented: [any differences]

### Input Type Verification
- [ ] Verify input type names are correct
  - Documented: `CopyProjectInput`
  - Actual in schema: [actual name or "NOT FOUND"]
  - Location: [file:line]

- [ ] Verify options input type
  - Documented: `CopyProjectOptionsInput`
  - Actual in schema: [actual name or "NOT FOUND"]
  - Location: [file:line]

### Input Parameters - CopyProjectInput
For each parameter:
- [ ] `projectId`
  - Documented type: `String!`
  - Actual type: [actual or "NOT FOUND"]
  - Required status matches: [Yes/No]

- [ ] `name`
  - Documented type: `String!`
  - Actual type: [actual or "NOT FOUND"]
  - Required status matches: [Yes/No]

- [ ] `description`
  - Documented type: `String`
  - Actual type: [actual or "NOT FOUND"]
  - Required status matches: [Yes/No]

- [ ] `imageURL`
  - Documented type: `String`
  - Actual type: [actual or "NOT FOUND"]
  - Required status matches: [Yes/No]

- [ ] `companyId`
  - Documented type: `String`
  - Actual type: [actual or "NOT FOUND"]
  - Required status matches: [Yes/No]

- [ ] `options`
  - Documented type: `CopyProjectOptionsInput!`
  - Actual type: [actual or "NOT FOUND"]
  - Required status matches: [Yes/No]

### Input Parameters - CopyProjectOptionsInput
For each documented option:
- [ ] `assignees` - Type: [actual type]
- [ ] `automations` - Type: [actual type]
- [ ] `checklists` - Type: [actual type]
- [ ] `customFields` - Type: [actual type]
- [ ] `discussions` - Type: [actual type]
- [ ] `discussionComments` - Type: [actual type]
- [ ] `dueDates` - Type: [actual type]
- [ ] `files` - Type: [actual type]
- [ ] `forms` - Type: [actual type]
- [ ] `people` - Type: [actual type]
- [ ] `projectUserRoles` - Type: [actual type]
- [ ] `statusUpdates` - Type: [actual type]
- [ ] `statusUpdateComments` - Type: [actual type]
- [ ] `tags` - Type: [actual type]
- [ ] `todoActions` - Type: [actual type]
- [ ] `todoComments` - Type: [actual type]
- [ ] `todoLists` - Type: [actual type]
- [ ] `todos` - Type: [actual type]

### Response Type
- [ ] Verify response type
  - Documented: `Boolean`
  - Actual type: [actual or "NOT FOUND"]

## 2. Implementation Verification

### Resolver Check
- [ ] Resolver exists for this operation
  - Location: [file:line]
  - Handler function: `[functionName]`

### Business Logic Verification
- [ ] Asynchronous operation claim
- [ ] Queue-based processing claim
- [ ] One copy per user limitation
- [ ] Size limit (250,000 tasks) enforcement
- [ ] Logical dependencies enforcement
- [ ] Name trimming and URL removal

### Validation Rules
- [ ] Required fields enforced in code
- [ ] Character limits (name max 50, description max 500)
- [ ] Archive status check for source project

## 3. Permission Verification

### Access Requirements
- [ ] Permission checks exist in resolver
- [ ] Source project access verification
- [ ] Target company membership verification
- [ ] Cross-company copy restrictions

## 4. Related Query Verification

### copyProjectStatus Query
- [ ] Query exists in schema
- [ ] Response fields match documentation:
  - [ ] `queuePosition`
  - [ ] `progressPercentage`
  - [ ] `project` object with `id` and `name`

## 5. Error Response Verification

### Error Codes
For each documented error:
- [ ] `PROJECT_NOT_FOUND`
  - Exists in codebase: [Yes/No]
  - Message matches: [Yes/No]

- [ ] `COMPANY_NOT_FOUND`
  - Exists in codebase: [Yes/No]
  - Message matches: [Yes/No]

- [ ] `CREATE_PROJECT_LIMIT`
  - Exists in codebase: [Yes/No]
  - Message matches: [Yes/No]

- [ ] Copy already in progress error
  - Error message matches: [Yes/No]

## 6. Code Example Verification

### Basic Example
- [ ] GraphQL syntax is valid
- [ ] All fields in mutation exist
- [ ] Required fields are included
- [ ] Options structure is correct

### Advanced Example
- [ ] All option parameters exist
- [ ] Optional parameters shown are valid
- [ ] Cross-company copy example is valid

## 7. Business Logic Claims Verification

### Asynchronous Processing
- [ ] Verify job queue implementation exists
- [ ] Verify background processing claim

### Size Limitations
- [ ] Verify 250,000 task limit is enforced
- [ ] Verify where this limit is checked

### Dependencies
- [ ] Verify `assignees` requires `people: true`
- [ ] Verify `discussionComments` requires `discussions: true`
- [ ] Verify `statusUpdateComments` requires `statusUpdates: true`
- [ ] Verify `projectUserRoles` requires `people: true`

### Queue Features
- [ ] Verify enterprise priority claim
- [ ] Verify 6-hour cache claim
- [ ] Verify one copy per user limitation

## 8. Documentation Completeness

### Missing from Docs
- [ ] List any options found in code but not documented
- [ ] List any parameters found but not documented
- [ ] List any error cases found but not documented

### Extra in Docs (Hallucinated)
- [ ] List any options documented but not in code
- [ ] List any parameters documented but not in code
- [ ] List any features documented but not implemented

## Summary

### Critical Issues (Must Fix)
1. **WRONG RESPONSE SCHEMA**: copyProjectStatus returns different fields than documented
   - Documented: `queuePosition`, `progressPercentage`, `project`
   - Actual: `oldCompany`, `oldProject`, `newCompany`, `newProjectName`, `isTemplate`, `isActive`, `queuePosition`, `totalQueues`
   - Missing: `progressPercentage` field doesn't exist
2. **MISSING OPTION**: `coverConfig: Boolean` option exists but not documented
3. **MISSING WORKER**: No actual copy implementation found in codebase
4. **NO DEPENDENCY VALIDATION**: Claims about option dependencies not enforced in code

### Minor Issues (Should Fix)
1. **Vague error message**: "Oops! Something went wrong" for copy in progress
2. **Character limits**: 50/500 char limits mentioned but not enforced in schema

### Verified Accurate ✅
1. All input types and parameters exist
2. 250,000 task limit is enforced
3. Enterprise priority system works
4. 6-hour cache system works
5. One copy per user limitation works
6. Permission requirements are correct
7. Error codes exist and match