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