
function isBlank(value) {
  return typeof value == 'string' && !$.trim(value) || typeof value == 'undefined' || value === null;
}

function redirectToVersion(version) {
  var targetUrl = '';

  if (!isBlank(version)) {
    targetUrl += '/' + version;
  }

  var pathname = window.location.pathname;
  var filename = pathname.substring(pathname.lastIndexOf("/") + 1);
  targetUrl += '/' + filename;
  window.location.href = targetUrl;
}
