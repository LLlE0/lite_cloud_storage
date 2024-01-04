async function submit(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    try {
        const response = await fetch('/registration/new', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password }),
            redirect: 'manual'
        });
        
        if (response.ok) {
            const data = await response.json();
            if (data.redirect) {
                window.location.href = data.redirect;
            }
        } else {
            const errorText = await response.text();
            alert(`Registration failed: ${errorText}`);
        }
    } catch (error) {
        console.error('Error:', error);
    }
}