Title: Features & Capabilities
Desc: Listing of features and capabilities provided by aah framework
Keywords: features, feature, aah framework, capabilities, batteries included, fully loaded
---
### Server and Extension Points

  * HTTP
  * HTTPS [refer to config]({{aah_docs_domain_url}}/app-config.html#section-server-ssl)
      * HTTP/2, can be disabled via configuration.
      * Certificate File and Key File.
      * [Let's Encrypt CA]({{aah_docs_domain_url}}/app-config.html#section-server-ssl-lets-encrypt) - automatic install and serve certificates.
  * UNIX Socket
  * [Server Extension Points]({{aah_docs_domain_url}}/server-extension.html) and [Request Lifecycle]({{aah_docs_domain_url}}/request-life-cycle.html).
  * Custom Server TLS Config refer to [documentation]({{aah_docs_domain_url}}/server.html#custom-tls-config).
  * `go1.8` Graceful shutdown feature.
  * Automatic Gzip Compression if client supports it.
  * <span class="badge lb-sm">Since v0.7</span> Server Access Log, refer to [documentation]({{aah_docs_domain_url}}/server-access-log.html).
  * <span class="badge lb-sm">Since v0.9</span> HTTP => HTTPS redirects, just enable it in the [config]({{aah_docs_domain_url}}/app-config.html#section-server-ssl-redirect-http).
  * <span class="badge lb-sm">Since v0.9</span> Server Dump Request & Response, refer to [documentation]({{aah_docs_domain_url}}/server-dump-log.html).

### Configuration

aah framework uses an enhanced version of [forge syntax]({{aah_docs_domain_url}}/configuration.html) (very similar to HOCON syntax) for application, route, project, security and i18n config files, etc.

  * Environment profiles are supported. For e.g: dev, qa, prod, etc.
  * Organize your config files as you need, you can always add `include` references.  

### URL Routing and Reverse Route

  * Customized version of High performance [httprouter](https://github.com/julienschmidt/httprouter).
  * Flexible [routes configuration]({{aah_docs_domain_url}}/routes-config.html) for application, [static files]({{aah_docs_domain_url}}/static-files.html) and [namespace/group]({{aah_docs_domain_url}}/routes-config.html#namespace-group-routes).
  * Supports Domains and Sub-domains. <span class="badge lb-sm">Since v0.6</span> Wildcard subdomain supported. Refer to [tutorial]({{aah_docs_domain_url}}/tutorial/domain-subdomain-and-wildcard-subdomain.html).
  * Flexible reverse route URL by `route name`.
  * Access root domain and subdomain reverse routes easily from view templates and application codebase.
  * Adding Controllers with or without sub-package names for routes. So `v1`, `v2`, sub-packages are possible.
  * Redirect Trailing Slash, Auto Options, and Method Not Allowed.
  * Custom Not Found handling via [Centralized Error Handler]({{aah_docs_domain_url}}/centralized-error-handler.html).
  * Max Body Size configuration at route level and global level.


### Request Parameters - Auto Parse and Bind

aah provides very flexible way to [auto parse and bind]({{aah_docs_domain_url}}/request-parameters-auto-bind.html) request values into appropriate Go data types. It supports following:

  * Bind any `Path`, `Form`, `Query` into controller action parameters.
  * Bind `JSON` or `XML` request body into `struct`.
  * Bind any `Path`, `Form`, `Query` into controller action `struct` fields.
  * Bind any `Path`, `Form`, `Query` into nested `struct` following `.` notation convention.
  * Bind supports bind of pointer and non-pointer target.
  * And you can also do combinations of above options
  * You can added your own [custom Value Parser by Type]({{aah_docs_domain_url}}/request-parameters-auto-bind.html#adding-custom-value-parser-by-type)

### Security - Authentication and Authorization

  * aah framework focus on flexible and powerful security implementation, easy to use and understand, it was inspired by [Shiro security library](https://shiro.apache.org). You can design your application secure, stable with authentication, authorization and session management.
  * Exposes clean and intuitive API that simplifies the developer's effort to make their application secure.
  * Terminology - Security can be really confusing because of the terminology used. To make life easier by clarifying some core concepts, so you understand how they’re reflected in the aah framework. Refer to [security terminology]({{aah_docs_domain_url}}/security-terminology.html).
  * aah security design goals to simplify application security by being intuitive and easy to use. Refer to [security design details]({{aah_docs_domain_url}}/security-design.html).
  * Highlights - [Authentication]({{aah_docs_domain_url}}/authentication.html), [Authorization]({{aah_docs_domain_url}}/authorization.html), [Session Management](#security-session-management)
  * Very flexible, you can implement Role based or Permission based or Role and [Permissions]({{aah_docs_domain_url}}/security-permissions.html) based secured application.
  * Out-of-the-box aah framework supports following auth schemes - `Form Auth`, `Basic Auth` and `Generic Auth`.
  * You can define one or more `Auth Scheme` for your application and can be mapped per route basis.

### Security - Session Management

  * aah framework provides `stateful` and `stateless` HTTP state management. Default is `stateless`. It is perfect fit for Web and API application, refer to [security configuration]({{aah_docs_domain_url}}/security-config.html).
  * Session data is Signed using HMAC and Encrypted using AES.
  * Out-of-the-box `cookie` and `file` session store is supported.
  * You can easily add your [own session store]({{aah_docs_domain_url}}/session.html).  

### Security - Anti-CSRF

<span class="badge lb-sm">Since v0.9</span> aah provides automatic Anti-CSRF (Cross Site Request Forgery) protection for the aah web application. It protects all the HTML forms on the page. Anti-CSRF protection is enabled by default, refer to [documentation]({{aah_docs_domain_url}}/anti-csrf-protection.html).

### View Engine

  * Go view engine with [partial inheritance support]({{aah_docs_domain_url}}/views.html) (Default engine).
  * Multiple view layouts for your unique use case.
  * Framework provided [template funcs]({{aah_docs_domain_url}}/template-funcs.html), Plus you can add your own easily.
  * Custom template delimiter for templates.
  * You can add your own view engine into the framework, refer [here]({{aah_docs_domain_url}}/view.html#adding-user-defined-view-engine-into-aah).
  * <span class="badge lb-sm">Since v0.6</span> You can use without layout too and take full-control of directory structure `pages/*` along with view rendering via `HTML*` methods.

### i18n Internalization and Localization

  * [Message files]({{aah_docs_domain_url}}/i18n.html) (aka Translation) supported with `Language ID + Region ID` or `Language ID`.
      * Language ID follows the two-letter `ISO 639-1` standard.
      * Region ID follows the two-letter `ISO 3166-1` standard.
  * Default fallback `i18n.default` if requested `Locale` is not found.
  * Messages are accessible from View template files, controller as well as anywhere in the application codebase.
  * Organize your message files with sub-directories as you like.
  * Zero coding efforts on localizing your application via header `Accept-Language`, by [URL Query Parameter]({{aah_docs_domain_url}}/tutorial/i18n-url-query-param.html) or by [Path Parameter]({{aah_docs_domain_url}}/tutorial/i18n-path-param.html).

### Middleware

  * Flexible [Middleware]({{aah_docs_domain_url}}/middleware.html) with [Abort]({{aah_docs_domain_url}}/middleware.html#abort-the-middleware-flow) feature and taking control of [response writing]({{aah_docs_domain_url}}/reply.html#done) within the framework.
  * Bring your `http.Handler`, `http.HandlerFunc` into [aah]({{aah_docs_domain_url}}/middleware.html#bring-go-lang-native-middleware-into-aah).

### Interceptors

  * aah framework supports [per controller and per action level]({{aah_docs_domain_url}}/interceptors.html) interceptors (`Before`, `After`, `Finally` and `Panic`).

### Secure HTTP Headers

aah provides application secure headers with many safe defaults for Web and RESTful API application. [Know more about configuration]({{aah_docs_domain_url}}/security-config.html#section-http-header).

  * X-XSS-Protection
  * X-Content-Type-Options
  * X-Frame-Options
  * Referrer-Policy
  * Strict-Transport-Security  (STS, aka HSTS)
  * Content-Security-Policy (CSP)
  * Public-Key-Pins (PKP, aka HPKP)
  * X-Permitted-Cross-Domain-Policies

### Reply Builder

  * Simple, efficient and chained [Reply builder]({{aah_docs_domain_url}}/reply.html) to compose your response.
  * Supports rich reply types `HTML`, `JSON`, `JSONP`, `XML`, `Text`, `Binary`, `File`, `FileDownload`, `FileInline`, `Redirect`, Custom Render, etc.

### Event Publisher/Emitter

Simple and efficient [Event Publisher]({{aah_docs_domain_url}}/event-publisher.html) with Asynchronous and Synchronous publish.

  * aah Server extension points built around event publisher.
  * Supports Publish Once mode too.

### Content Negotiation

Content negotiation feature is used to place `MIME` type restriction on HTTP header `Content-Type` and `Accept` for your REST API application. Some cases useful for web application too.

How to configure one, refer to [documentation]({{aah_docs_domain_url}}/app-config.html#section-content-negotiation).

### Centralized Error Handling

<span class="badge lb-sm">Since v0.8</span> aah provides centralized error handling for your application.

  * Framework utilizes this error handler across for all the HTTP error responses. Refer to [documentation]({{aah_docs_domain_url}}/centralized-error-handler.html).
  * Framework propagates all Error responses to Centralized Error Handler, you can control and customized the response.


### Static File Delivery

aah framework supports flexible and easy to use and configure [static file]({{aah_docs_domain_url}}/static-files.html) delivery.

  * Serves directory and it's subtree files.
  * Serves individual files.
  * Directory listing.
  * <span class="badge lb-sm">Since v0.6</span> Static files `Cache-Control` by mime types and default one. It gets applied only to `prod` environment profile. Refer to [documentation]({{aah_docs_domain_url}}/static-files.html#cache-control).
  * <span class="badge lb-sm">Since v0.7</span> Cache Busting using file name. Refer to [documentation]({{aah_docs_domain_url}}/static-files.html#cache-busting).
  * All capabilities of `http.ServeContent`.

### Logger

  * Simple to use log library and it's [configuration]({{aah_docs_domain_url}}/log-config.html).
  * Supported Log `Level`'s are `ERROR`, `WARN`, `INFO`, `DEBUG` and `TRACE`.
  * You can create multiple log instances besides the default one.
  * Out-of-the-box `Console` and `File` receivers are supported, use `Hook` for exporting your log to systems like splunk, kibana, etc.
  * `File` receiver supports `daily` log rotation, etc.
  * Define your custom log message format (text, json) in the config.
  * <span class="badge lb-sm">Since v0.6</span> you can bind standard Go logger enabled libraries with aah logger (`log.ToGoLogger()`), unified log at one place.
  * <span class="badge lb-sm">Since v0.7</span> supports logger `Hook`.

### HTML Minify

* <span class="badge lb-sm">Since v0.6</span> Framework provides HTML minify feature, refer to [minify tutorial]({{aah_docs_domain_url}}/tutorial/html-minify.html).
* HTML minify gets applied only to `prod` environment profile.

### Easy to use Application Binary

  * Easy to build and deploy [aah application binary]({{aah_docs_domain_url}}/aah-application-binary.html).
  * Cross compile build is supported (aah is only recognizes cross compile build request and setting the appropriate values, Go lang does the rest for you).

### Essentials Library

aah helps to increase your productivity, the framework's [essentials](https://godoc.org/aahframework.org/essentials.v0) library provides a lot of useful helper/util methods in the following areas:

  * filepath
  * GUID (Globally Unique Identifier)
  * Crypto random string, Math random string, random byte generation at fixed length
  * go
  * io
  * os
  * reflect
  * string
  * archive (zip)

Refer to [godoc](https://godoc.org/aahframework.org/essentials.v0).

### Hot-Reload for Development

  * <span class="badge lb-sm">Since v0.7</span> aah provides Hot-Reload for Development purpose.
  * Fire the `aah run` and forget the terminal. You can focus on your code and refresh the browser to see your changes.
  * Still lot of improvements can come-in; in-terms of formatted error display, watch files optimization, etc. So keep me posted on your issues. Gradually I will bring improvements :)

<br>
<center>**Spread the word of `aah`, the web framework for Go. Thank you!**</center>
