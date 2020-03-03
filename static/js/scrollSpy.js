  (function() {
    'use strict';

  let section = document.querySelectorAll(".section");
  let sections = {};
  let i = 0;

  Array.prototype.forEach.call(section, function(e) {
  sections[e.id] = e.offsetTop;
});

window.onscroll = function() {
  let scrollPosition = document.documentElement.scrollTop;

  for (let i in sections) {
    if (sections[i] <= scrollPosition) {
      document.querySelector('.active').setAttribute('class', 'none');
      document.querySelector('a[href*=' + i + ']').setAttribute('class', 'active');
    }
  }
};
})();