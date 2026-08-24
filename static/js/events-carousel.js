function initEventsCarousel() {
    const container = document.getElementById('events-carousel');
    if (!container || container.dataset.carouselInitialized === "true") return;

    container.dataset.carouselInitialized = "true";

    const leftBtn = document.getElementById('scroll-left');
    const rightBtn = document.getElementById('scroll-right');
    const wrappers = container.querySelectorAll('.event-card-wrapper');

    const getStep = () => {
        const first = wrappers[0];
        return first ? first.offsetWidth : 320;
    };

    if (leftBtn) {
        leftBtn.addEventListener('click', () => {
            container.scrollBy({ left: -getStep(), behavior: 'smooth' });
        });
    }

    if (rightBtn) {
        rightBtn.addEventListener('click', () => {
            container.scrollBy({ left: getStep(), behavior: 'smooth' });
        });
    }

    const activeClasses = ['opacity-100', 'scale-100', 'ring-2', 'ring-amber-400', 'shadow-2xl', 'z-10'];
    const inactiveClasses = ['opacity-40', 'scale-90'];

    const updateActiveCard = () => {
        const containerCenter = container.getBoundingClientRect().left + container.offsetWidth / 2;
        let closestWrapper = null;
        let minDistance = Infinity;

        wrappers.forEach(wrapper => {
            const rect = wrapper.getBoundingClientRect();
            const wrapperCenter = rect.left + rect.width / 2;
            const distance = Math.abs(containerCenter - wrapperCenter);

            if (distance < minDistance) {
                minDistance = distance;
                closestWrapper = wrapper;
            }
        });

        wrappers.forEach(wrapper => {
            const card = wrapper.querySelector('.event-card');
            if (!card) return;

            if (wrapper === closestWrapper) {
                card.classList.remove(...inactiveClasses);
                card.classList.add(...activeClasses);
            } else {
                card.classList.remove(...activeClasses);
                card.classList.add(...inactiveClasses);
            }
        });
    };

    let ticking = false;
    container.addEventListener('scroll', () => {
        if (!ticking) {
            window.requestAnimationFrame(() => {
                updateActiveCard();
                ticking = false;
            });
            ticking = true;
        }
    });

    window.addEventListener('resize', updateActiveCard);
    updateActiveCard();
}

if (typeof htmx !== 'undefined') {
    htmx.onLoad(() => initEventsCarousel());
} else {
    document.addEventListener('DOMContentLoaded', () => initEventsCarousel());
}
