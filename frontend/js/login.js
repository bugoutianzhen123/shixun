loginForm.addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(loginForm);
    const messageDiv = document.getElementById('message');
    const formDataJSON = {
        username: formData.get('username'),
        password: formData.get('password')
    };

    console.log('Form Data:', formDataJSON);

    fetch('http://localhost:8088/user/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(formDataJSON),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('登录失败');
            }
            return response.json();
        })
        .then(data => {
            console.log('Response Data:', data);

            messageDiv.textContent = data.message;

            if (data.data && data.data.token && data.data.permission !== undefined) {
                localStorage.setItem('token', data.data.token);

                if (data.data.permission === 1) {
                    window.location.href = 'http://localhost:63342/code/html/frontend/userPanel.html';
                } else if (data.data.permission === 0) {
                    window.location.href = 'http://localhost:63342/code/html/frontend/adminPanel.html';
                } else {
                    console.error('未知权限');
                }
            } else {
                console.error('Token 或 Permission 缺失');
                messageDiv.textContent = '登录失败，请重试';
            }
        })
        .catch(error => {
            console.error('Error logging in:', error);
            messageDiv.textContent = '登录失败，请重试';
        });
});
