
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
  var isVersion = /^v(\d+.\d+)$/.test(filename)
  if (!isVersion) {
    targetUrl += '/' + filename;
  }

  window.location.href = targetUrl;
}

function tooltipHandling() {
  $('[data-toggle="tooltip"]').tooltip();
}

function anchorTagHandling() {
  anchors.options = {
    placement: 'right',
    class: 'aah-color-imp',
    icon: '¶'
  };
  anchors.add().remove('.no-anchor');

  $('a[href^="#"').click(function() {
    if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'')
      && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
      if (target.length) {
        $('html,body').animate({
          scrollTop: target.offset().top - 95 //offsets for fixed header
        }, 1000);
        window.location.hash = this.hash
        return false;
      }
    }
  });

  // scroll to anchor tag
  var anchorVal = location.href.split("#")[1];
  if($(anchorVal)) {
    var target = $('#'+anchorVal);
    if (target.length) {
      $('html,body').animate({
        scrollTop: target.offset().top - 95 //offset height of header here too.
      }, 1000);
      return false;
    }
  }
}
