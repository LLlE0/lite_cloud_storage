async function submit(event) {
    event.preventDefault();
    console.log("Started func")
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    try {
        const response = await fetch('/auth/try', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        });
        if (response.status === 200) {
            //window.location.href = '/index';
        } else {
            alert('Invalid username or password');
        }
    } catch (error) {
        console.error('Error:', error);
    }
}