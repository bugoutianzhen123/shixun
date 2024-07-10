document.getElementById("registerForm").addEventListener("submit", function(event) {
    event.preventDefault();

    // 获取表单数据
    let formData = new FormData(this);
    let data = {
        username: formData.get("username"),  // 确保这里的键名和后端一致
        password: formData.get("password")  // 确保这里的键名和后端一致
    };

    // 检查数据是否为空
    if (!data.username || !data.password) {
        alert('用户名和密码不能为空');
        return;
    }

    // 在发送请求之前，确认数据
    console.log('Data to be sent:', data);

    // 发送请求到后端
    fetch('http://localhost:8088/user/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Registration successful!');
                window.location.href = 'login.html';
            } else {
                alert('Registration failed: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error registering.');
        });
});
