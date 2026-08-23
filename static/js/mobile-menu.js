function toggleMobileMenu() {
    const menu = document.getElementById('mobile-menu');
    if (!menu) return;

    const isOpen = menu.style.maxHeight !== '0px' && menu.style.maxHeight !== '';

    if (isOpen) {
        closeMobileMenu();
    } else {
        openMobileMenu();
    }
}

function openMobileMenu() {
    const menu = document.getElementById('mobile-menu');
    const content = document.getElementById('mobile-nav-content');
    if (!menu) return;

    menu.style.maxHeight = menu.scrollHeight + 'px';
    menu.style.opacity = '1';
    menu.style.pointerEvents = 'auto';

    if (content) {
        setTimeout(() => {
            content.style.opacity = '1';
            content.style.transform = 'translateY(0px)';
        }, 75);
    }

    setMenuState(true);
}

function closeMobileMenu() {
    const menu = document.getElementById('mobile-menu');
    const content = document.getElementById('mobile-nav-content');
    if (!menu) return;

    if (content) {
        content.style.opacity = '0';
        content.style.transform = 'translateY(-10px)';
    }

    menu.style.maxHeight = '0px';
    menu.style.opacity = '0';
    menu.style.pointerEvents = 'none';

    setMenuState(false);
}

function setMenuState(isOpen) {
    const burger = document.getElementById('icon-burger');
    const close = document.getElementById('icon-close');

    if (!burger || !close) return;

    if (isOpen) {
        burger.style.opacity = '0';
        burger.style.transform = 'rotate(90deg) scale(0.5)';

        close.style.opacity = '1';
        close.style.transform = 'rotate(0deg) scale(1)';
    } else {
        burger.style.opacity = '1';
        burger.style.transform = 'rotate(0deg) scale(1)';

        close.style.opacity = '0';
        close.style.transform = 'rotate(-90deg) scale(0.5)';
    }
}

document.body.addEventListener('htmx:afterSwap', function() {
    closeMobileMenu();
});
