(function toggleMenu() {
    'use strict';

    const navigationMenu = document.getElementById('navigation-list');
    const hamburguerMenu = document.getElementById('hamburguerMenu');

    hamburguerMenu.onclick = function() {
        navigationMenu.classList.toggle('active');
    };
})();