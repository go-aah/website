
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
  var isVersion = /^v(\d+.\d+.*)$/.test(filename)
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

  $('a[href^="#"]').click(function() {
    if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'')
      && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
      if (target.length) {
        $('html,body').animate({
          scrollTop: target.offset().top - 160 //offsets for fixed header
        }, 1200);
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
        scrollTop: target.offset().top - 160 //offset height of header here too.
      }, 1000);
      return false;
    }
  }
}

function resizeAahBanner() {
  var aahBanner = $('#aahBanner')
  if (aahBanner) {
    var height = window.innerHeight;
    if (navigator.userAgent.match(/iPad/i) != null) {
      height = 700;
    }
    if (height >= 400) {
      height-=155;
    }
    aahBanner.css('height', height);
    aahBanner.css('margin-top', height/6);
  }
}

// Back 2 Top
$(window).scroll(function() {
  var height = $(window).scrollTop();
  if (height > 400) {
    $('#back2Top').fadeIn();
  } else {
    $('#back2Top').fadeOut();
  }
});

$(document).ready(function() {
  $('#back2Top').click(function(event) {
    event.preventDefault();
    $('html,body').animate({ scrollTop: 0 }, 1500);
    return false;
  });

  $('#heyImHereBtn').click(function() {
    var aahBanner = $('#aahBanner');
    var height = parseInt(aahBanner.css('height'));
    if (height > 484) {
      aahBanner.css('height', height-65);
    }
  });
});

(function($) {
    var element = $('.follow-scroll'), elementOffset = element.offset();
    if (elementOffset) {
      var originalY = elementOffset.top;

      // Space between element and top of screen (when scrolling)
      var topMargin = 70;

      // Should probably be set in CSS; but here just for emphasis
      element.css('position', 'relative');

      $(window).on('scroll', function(event) {
          var scrollTop = $(window).scrollTop();

          element.stop(false, false).animate({
              top: scrollTop < originalY ? 0 : scrollTop - originalY + topMargin
          }, -100);
      });
    }
})(jQuery);
