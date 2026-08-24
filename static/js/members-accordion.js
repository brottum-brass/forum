function openCard(card) {
    card.style.width = "340px";
    
    const title = card.querySelector(".vertical-title");
    if (title) title.style.opacity = "0";

    const content = card.querySelector(".expanded-content");
    if (content) {
        content.style.visibility = "visible";
        content.style.opacity = "1";
        content.style.transform = "translateY(0)";
    }
}

function closeCard(card) {
    card.dataset.pinned = "false";
    clearTimeout(card.cardTimer);
    
    const content = card.querySelector(".expanded-content");
    if (content) {
        content.style.opacity = "0";
        content.style.transform = "translateY(10px)";
        content.style.visibility = "hidden";
    }
    
    card.style.width = "76px";
    
    const title = card.querySelector(".vertical-title");
    if (title) title.style.opacity = "1";
}

function closeAllCards(exceptCard = null) {
    const cards = document.querySelectorAll(".member-accordion-card");
    cards.forEach((card) => {
        if (card !== exceptCard) {
            closeCard(card);
        }
    });
}

function setupAccordionCard(card) {
    if (card.dataset.initialized === "true") return;
    card.dataset.initialized = "true";
    card.dataset.pinned = "false";

    card.addEventListener("click", () => {
        const isPinned = card.dataset.pinned === "true";

        closeAllCards(card);

        if (isPinned) {
            closeCard(card);
        } else {
            card.dataset.pinned = "true";
            openCard(card);
        }
    });

    card.addEventListener("mouseenter", () => {
        if (card.dataset.pinned === "true") return;
        
        card.style.width = "340px";
        const title = card.querySelector(".vertical-title");
        if (title) title.style.opacity = "0";

        clearTimeout(card.cardTimer);
        const content = card.querySelector(".expanded-content");
        
        if (content) {
            card.cardTimer = setTimeout(() => {
                if (card.dataset.pinned === "true") return;
                content.style.visibility = "visible";
                content.style.opacity = "1";
                content.style.transform = "translateY(0)";
            }, 250);
        }
    });

    card.addEventListener("mouseleave", () => {
        if (card.dataset.pinned === "true") return;

        closeCard(card);
    });
}

function initMemberAccordion() {
    const cards = document.querySelectorAll(".member-accordion-card");
    cards.forEach(setupAccordionCard);
}

document.addEventListener("click", (e) => {
    if (!e.target.closest(".member-accordion-card")) {
        closeAllCards();
    }
});

if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", initMemberAccordion);
} else {
    initMemberAccordion();
}

document.body.addEventListener("htmx:afterSwap", initMemberAccordion);
