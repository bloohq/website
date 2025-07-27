# Verification for: upload-files.md
Path: /content/en/api/1.start-guide/8.upload-files.md
Status: [🔧] Fixed - Updated REST API file size limit from 5GB to 4.8GB to match implementation

## 1. GraphQL Upload Verification

### uploadFile Mutation
- [✓] Mutation exists
  - Location: /Users/manny/Blue/bloo-api/src/resolvers/Mutation/uploadFile.ts
  - Input type: `UploadFileInput` with file, projectId, companyId
  - Returns: `File!` with all documented fields

### uploadFiles Mutation
- [✓] Mutation exists
  - Location: /Users/manny/Blue/bloo-api/src/resolvers/Mutation/uploadFiles.ts
  - Input type: `UploadFilesInput` with files array, projectId, companyId
  - Returns: `[File!]!` array

### File Size Limits - GraphQL
- [✓] 256MB limit verified
  - Location: /Users/manny/Blue/bloo-api/src/lib/file-upload-utils.ts:141
  - Code: `const maxSize = 256 * 1024 * 1024; // 256MB`

### Batch Upload Limits
- [✓] 10 files maximum verified
  - Location: /Users/manny/Blue/bloo-api/src/resolvers/Mutation/uploadFiles.ts:27-30
- [✓] 1GB total batch size limit verified
  - Location: /Users/manny/Blue/bloo-api/src/resolvers/Mutation/uploadFiles.ts:36

## 2. REST API Upload Verification

### /uploads Endpoint
- [✓] Endpoint exists
  - Location: /Users/manny/Blue/bloo-api/src/lib/uploads.ts
  - Mounted at: /uploads in app.ts:195-204

### S3 Credentials Response
- [✓] Returns S3 credentials as documented
  - Returns signed URL with method, headers, fields
  - Response format matches documentation

### File Size Limits - REST API
- [❌] **DISCREPANCY FOUND**
  - **Documented**: 5GB limit
  - **Actual**: 4.8GB limit
  - Location: /Users/manny/Blue/bloo-api/src/lib/app.ts
  - Code: `maxSize: 1000000000 * 4.8, // 4800 MB or 4.8 GB`

## 3. Additional Mutations Verification

### createFile Mutation
- [✓] Mutation exists
  - Location: /Users/manny/Blue/bloo-api/src/resolvers/Mutation/createFile.ts
  - Used for REST API file registration

### createTodoCustomFieldFile Mutation
- [✓] Mutation exists
  - Location: /Users/manny/Blue/bloo-api/src/resolvers/Mutation/createTodoCustomFieldFile.ts
  - Associates files with todo custom fields

## 4. File Type Schema Verification

### File Type Fields
- [✓] All documented fields exist in schema:
  - id, uid, name, size, type, extension, shared, createdAt
  - project, folder relationships
  - FileStatus enum (CONFIRMED, PENDING)

### Upload Scalar
- [✓] Upload scalar type defined in schema
- [✓] Multipart/form-data handling implemented

## 5. Authentication Verification

### GraphQL Upload Permissions
- [✓] uploadFile: requires authentication only
- [✓] uploadFiles: requires authentication only
- [✓] createFile: requires project access permissions

### REST API Authentication
- [✓] Requires Bearer token authentication
- [✓] Headers verified in uploads.ts

## 6. Code Examples Verification

### Apollo Client Examples
- [✓] GraphQL syntax is valid
- [✓] Variable structures match schema
- [✓] File handling approach is correct

### Vanilla JavaScript Examples
- [✓] multipart/form-data format is correct
- [✓] GraphQL multipart spec compliance verified
- [✓] Form data construction matches specification

### Python REST API Examples
- [✓] Three-step process is accurate
- [✓] S3 upload format matches actual implementation
- [✓] GraphQL mutations are correct

### cURL Examples
- [✓] Syntax is valid
- [✓] Headers and form data are correct

## 7. File Processing Verification

### S3 Key Generation
- [✓] Key format verified in file-upload-utils.ts
- [✓] Project/company path structure matches
- [✓] Filename sanitization implemented

### File Status Workflow
- [✓] PENDING → CONFIRMED workflow verified
- [✓] Upload confirmation process documented correctly

## Summary

### Critical Issues (Must Fix)
None found - all core functionality works as documented

### Minor Issues (Should Fix)
1. **REST API file size limit**: Documentation claims 5GB, actual implementation is 4.8GB

### Suggestions
1. Update REST API file size limit from 5GB to 4.8GB to match implementation
2. All other documentation is accurate and comprehensive

### Overall Assessment
Excellent documentation with comprehensive examples covering both GraphQL and REST approaches. Only one minor discrepancy in file size limits needs correction.