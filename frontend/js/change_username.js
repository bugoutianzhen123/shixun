const changeUsernameForm = document.getElementById('changeUsernameForm');
const changeUsernameMessage = document.getElementById('changeUsernameMessage');

changeUsernameForm.addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(changeUsernameForm);
    const newUsername = formData.get('newUsername');

    const token = localStorage.getItem('token');
    if (!token) {
        changeUsernameMessage.textContent = '请先登录';
        return;
    }

    fetch('http://localhost:8088/user/changeusername', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify({ newUsername: newUsername }),
    })
        .then(response => {
            if (!response.ok) {
                return response.json().then(error => {
                    throw new Error(error.message || 'Failed to change username');
                });
            }
            return response.json();
        })
        .then(data => {
            changeUsernameMessage.textContent = data.message;
        })
        .catch(error => {
            console.error('Error changing username:', error);
            changeUsernameMessage.textContent = error.message || 'Failed to change username. Please try again.';
        });
});
