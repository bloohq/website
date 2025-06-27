- This is a new website for Blue, a b2b saas process management system 
- The idea is to make a site with a high degree of polish to be on a world class level like Stripe, OpenAI, and Linear. 
- I am a single developer, so this has to be simple and easy maintanble. 

## Rules

- Write and setup idiomatic golang
- Tailwind CLI is alwasy in watch mode, no need for you to rebuild.


## Plan

1. Setup initial folder structure ✅
2. Setup Golang ✅
3. Install Goldmark, Air, godotev, https://github.com/go-yaml/yaml ✅
4. Setup file based route navigation for /pages with a dedicated router inside a web package ✅
5. Setup main.html with left/right sidebar and navbar components ✅
6. Add favicon.ico and logo.svg ✅
7. Added data/nav.json for dynamic rendering of navigation acros components. ✅
8. Made topbar buttons hotkeys work ✅
9. Make left/right sidebar toggle open/close ✅
10. Setup TailwindCSS Properly with CLI and Watcher ✅
11. Dynamic left sidebar ✅
12. Metadata.json for dynmaic titles and so on ✅
13. Markdown generation ✅
14. Dynamic sidebar for docs and api docs ✅
15. Centralized redirets in redirects.json ✅
16. Prose styling
17. Right sidebar dynamic both for html pages and markdown files with up and down arrows ✅
18. Replace font-awesome with svgs for arrows and light/dark mode?
19. dark mode implementation
20. Make legal navigation dynamic and page rendering working ✅
21. Sales modal?
22. Changelog Implementation ✅
23. Pricing Page
24. Roadmap   ✅
25. Review what components to "borrow" from other sites.
26. Realtime search ✅
27. YouTube widget for markdown ✅
28. Improve search
29. Take all old redirects from previous website ✅
30. Align all elements in respective footers perfectly. 
31. Customers page (testimonial videos essentially) ✅
32. Individual insight 
33. Brand page
34. AlpineJS components setup ✅
35. Home page structure / Messaging
36. Add a global scope for customer numbers, just need to update in one place and updates everywhere. ✅ 
37. Make the website multi linguagal?  Really like that to be honest.
38. For api docs, put company or project id and all the mutations update in the examples. Like Stripe.
39. Make my own status page (frontend) ✅
40. Find all "big" brand customers via AI and add them to logos on customers page.
41. Add SOP to website?
42. Add Git commit history to website?
43. Fix ana upside down video lol
44. Fix white label add on button
45. Status page with backend
46. Make page titles a component ✅
47. Customer stories (eventually)
48. Create a api endpoint on blue to count companies, and then use that to power website customer count.




## Tech Stack

- Golang
- AlpineJS
- @alpinejs/collapse@3 plugin
- https://github.com/markmead/alpinejs-component (for components)
- TailwindCSS v4


## Folder Structure

- Layouts
- Pages (main pages and subfolder with pages, we will have route based navigation)
- Public (all static assets, images, etc)
- SEO (files for SEO like metadata and redirects, schema, robots, etc)

## Full Website Structure

```
.
├── components/
│   ├── client-logos.html
│   ├── head.html
│   ├── left-sidebar.html
│   ├── page-heading.html
│   ├── right-sidebar.html
│   ├── testimonial-videos.html
│   └── topbar.html
├── content/
│   ├── agency-success-guide/
│   ├── alternatives/
│   ├── api-docs/
│   │   ├── 1.start-guide/
│   │   ├── 11.libraries/
│   │   ├── 2.projects/
│   │   ├── 3.records/
│   │   ├── 5.custom fields/
│   │   ├── 6.automations/
│   │   ├── 7.user management/
│   │   ├── 8.company-management/
│   │   └── 9.dashboards/
│   ├── company-news/
│   ├── docs/
│   │   ├── 1.start-guide/
│   │   ├── 10.use cases/
│   │   ├── 2.projects/
│   │   ├── 3.records/
│   │   ├── 4.views/
│   │   ├── 5.custom fields/
│   │   ├── 6.automations/
│   │   │   └── 4.actions/
│   │   ├── 7.user management/
│   │   │   └── 8.roles/
│   │   ├── 8.dashboards/
│   │   └── 9.integrations/
│   ├── engineering-blog/
│   ├── frequently-asked-questions/
│   ├── insights/
│   ├── legal/
│   ├── modern-work-practices/
│   ├── product-updates/
│   ├── project-management-dictionary/
│   └── tips-&-tricks/
├── data/
│   ├── nav.json
│   ├── metadata.json
│   └── (other data files)
├── layouts/
│   └── main.html
├── pages/
│   ├── company/
│   ├── platform/
│   └── (main page HTML files)
├── public/
│   ├── customers/
│   ├── logo/
│   ├── testimonials/
│   └── (static assets)
├── seo/
│   ├── redirects.json
│   └── (SEO-related files)
├── web/
│   └── (Go web package files)
├── go.mod
├── go.sum
├── main.go
├── tailwind.config.js
└── CLAUDE.md
```

## Development Workflow

### Starting Development
```bash
# Start the Go server with hot reload
air

# Tailwind CSS is already in watch mode
# Access site at http://localhost:8080
```

### Common Commands
```bash
# Run Go server
go run main.go

# Build for production
go build -o blue-website

# Format Go code
go fmt ./...

# Run tests
go test ./...
```

## Key Architecture Decisions

### Routing System
- File-based routing from `/pages` directory
- Dynamic route handling in `web/router.go`
- HTML pages map directly to URLs (e.g., `/pages/pricing.html` → `/pricing`)

### Component System
- Server-side HTML components in `/components`
- AlpineJS for client-side interactivity
- Components are included via Go templates

### Data Management
- `data/nav.json` - Navigation structure for all menus
- `data/metadata.json` - Page titles, descriptions, and SEO data
- `seo/redirects.json` - URL redirect mapping

### Content Structure
- Markdown files in `/content` are rendered with Goldmark
- YAML frontmatter support for metadata
- Automatic table of contents generation for docs

## Component Patterns

### Creating New Components
1. Add HTML file to `/components`
2. Use AlpineJS directives for interactivity
3. Include in layouts via `{{template "component-name" .}}`

### Dynamic Data in Components
- Pass data through Go template context
- Use `x-data` for AlpineJS state
- Global data available via `window.blueData`

## Important Files

### Core Configuration
- `main.go` - Entry point and server setup
- `web/router.go` - Route handling logic
- `web/markdown.go` - Markdown processing
- `tailwind.config.js` - Tailwind configuration

### Key Components
- `layouts/main.html` - Main layout wrapper
- `components/topbar.html` - Navigation bar
- `components/left-sidebar.html` - Documentation navigation
- `components/right-sidebar.html` - Page table of contents

## Debugging Tips

### Common Issues
1. **404 Errors** - Check file exists in `/pages` or `/content`
2. **Styling Issues** - Ensure Tailwind CLI is running
3. **Component Not Rendering** - Verify template name matches
4. **Markdown Not Parsing** - Check frontmatter format

### Debug Mode
- Air provides hot reload and error messages
- Check browser console for AlpineJS errors
- Go template errors appear in terminal

## Future Enhancements Needed

### High Priority
- [ ] Internationalization (i18n) system
- [ ] API endpoint for live customer count
- [ ] Backend for status page
- [ ] Search improvements (fuzzy matching, filters)

### Medium Priority
- [ ] Dark mode implementation
- [ ] Component library documentation
- [ ] Performance optimization (caching, CDN)
- [ ] Analytics integration

### Nice to Have
- [ ] Git commit history display
- [ ] Multi-language support
- [ ] Advanced search with Algolia/Meilisearch
- [ ] A/B testing framework
