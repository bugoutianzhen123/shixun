document.addEventListener('DOMContentLoaded', function() {
    // Fetch user information
    fetch('http://localhost:8088/user/getinfo', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById('username').textContent = data.username;
            document.getElementById('permission').textContent = data.permission;
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error fetching user information.');
        });
});

// Change password
document.getElementById("changePasswordForm").addEventListener("submit", function(event) {
    event.preventDefault();
    let formData = new FormData(this);
    fetch('http://localhost:8088/user/changepassword', {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Password changed successfully!');
            } else {
                alert('Failed to change password.');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error changing password.');
        });
});

// Change username
document.getElementById("changeUsernameForm").addEventListener("submit", function(event) {
    event.preventDefault();
    let formData = new FormData(this);
    fetch('http://localhost:8088/user/changeusername', {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Username changed successfully!');
                document.getElementById('username').textContent = formData.get('newUsername');
            } else {
                alert('Failed to change username.');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error changing username.');
        });
});
