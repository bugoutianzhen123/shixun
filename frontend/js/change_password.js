const changePasswordForm = document.getElementById('changePasswordForm');
const changePasswordMessage = document.getElementById('changePasswordMessage');

changePasswordForm.addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = new FormData(changePasswordForm);
    const currentPassword = formData.get('currentPassword');
    const newPassword = formData.get('newPassword');

    fetch('http://localhost:8088/user/changepassword', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token') // 添加token
        },
        body: JSON.stringify({ Prepassword: currentPassword, Newpassword: newPassword }),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to change password');
            }
            return response.json();
        })
        .then(data => {
            changePasswordMessage.textContent = data.message;
        })
        .catch(error => {
            console.error('Error changing password:', error);
            changePasswordMessage.textContent = 'Failed to change password. Please try again.';
        });
});
