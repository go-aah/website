Title: Security
Desc: aah takes security very seriously. What aah framework automatically does, security considerations, etc.
Keywords: put security first, security features, aah automatically does, security considerations
---
# Security

aah takes security very seriously, you're welcome to do peer review of aah's 100% open source code to ensure nobody's aah application is ever compromised or hacked. As an application developer you're responsible for any security breaches. I do my best to make sure aah application is as secure as possible.

  * [How to report security issues?](/security/vulnerabilities.html)
  * [What aah framework automatically does?](#what-aah-framework-automatically-does)
  * [Security Considerations](#security-considerations)

## What aah framework automatically does?

aah framework automatically does following safety measures to protect aah application from various attacks.

  * [Auto Parse and Bind]({{aah_docs_domain_url}}/request-parameters-auto-bind.html) Request parameter does value sanitization to prevent XSS attacks, it's highly recommended to use.
      - Any parse errors results in `400 Bad Request`.
  * All view template functions (such as i18n, config, session, flash, etc.) provided aah does sanitization to prevent XSS attacks.
      - String values are appropriately escaped.
      - URLs are parsed and encoded appropriately.
  * Secured session cookies with AES encrypted and HMAC signed.
      - Automatically `secure` is set to true if application uses SSL/TLS and always `HttpOnly` is set to true.
  * [Anti-CSRF protection]({{aah_docs_domain_url}}/anti-csrf-protection.html) is enabled by default for web application. aah automatically protects all the HTML forms on the page.
      - It is recommended to use Logout with POST request. Refer to [form based auth tutorial]({{aah_docs_domain_url}}/tutorial/form-based-auth.html)
  * [Secure HTTP Headers]({{aah_docs_domain_url}}/security-config.html#section-http-header) with many safe defaults for Web application and RESTful API service.
      - Enforces XSS filter in the browser.
      - Sets HSTS if application uses SSL/TLS to prevent protocol downgrade attacks and cookie hijacking.
      - Sets Content-Type options to `nosniff` to prevents Content Sniffing or MIME sniffing.
      - Sets frame options to `sameorigin` to prevents Clickjacking.
      - Sets Referrer-Policy explicitly to `no-referrer-when-downgrade`.
  * To protect against DDoS attacks caused by large HTTP request bodies by enforcing a hard limit.
      - Once limit is hit, aah server stops reading the request body and immediately closes the client connection.
      - Default value is 5MB for all requests and 32MB Multipart request. You can set this size limit at individual route level and global level.
  * Static File Delivery - Prevents directory traversal vulnerability.
  * All the errors and traces from framework gets logged into log file, not exposed outside on `prod` environment profile.
  * Plus Go lang provided safety measures.

## Security Considerations

  * aah provides very powerful, easy to use [Authentication]({{aah_docs_domain_url}}/authentication.html) and [Authorization]({{aah_docs_domain_url}}/authorization.html) capabilities to secure your application. Inspired by apache Shiro library.
      - Use `Before` action interceptor to check user has appropriate roles and permissions. Then permit user to access the data.
      - **`Upcoming`** Automatic roles and permissions check for individual routes via configuration.
  * Don't use a weak sign key and encryption key for Session. The `aah new` generates strong one for you automatically.
  * Don't use a weak sign key and encryption key for Anti-CSRF. The `aah new` generates strong one for you automatically.
  * CORS protection - preflight.
