document.addEventListener('click', async (e) => {
    const button = e.target.closest('.copy-code-button');
    
    if (!button) return;

    e.preventDefault();

    const container = button.closest('.code-container');
    if (!container) return;
    
    const codeElement = container.querySelector('code');
    if (!codeElement) return;
    
    try {
        await navigator.clipboard.writeText(codeElement.innerText);
        
        button.src = '/static/assets/copy/copied.svg';
        
        setTimeout(() => {
            button.src = '/static/assets/copy/copy.svg';
        }, 1500);
        
    } catch (err) {
        console.error('Failed to copy text: ', err);
    }
});

document.addEventListener('click', async (e) => {
    const button = e.target.closest('.category-button');
    if (!button) return;

    e.preventDefault();

    const categoryList = button.nextElementSibling;
    if (!categoryList) return;

    categoryList.classList.toggle('hidden'); 
    const isHidden = window.getComputedStyle(categoryList).display === 'none';

    const chevronImage = button.querySelector('img.chevron');
    if (!chevronImage) return;

    chevronImage.classList.toggle('rotated', isHidden); // rotate it
});