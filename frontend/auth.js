async function submit(event) {
    // Prevent the form from submitting normally
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const response = await fetch('/auth', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
    });
    const result = await response.json();
    if (result.success) {
        window.location.href = '/index';
    } else {
        alert('Invalid username or password');
    }
}