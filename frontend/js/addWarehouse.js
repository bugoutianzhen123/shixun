document.getElementById("addWarehouseForm").addEventListener("submit", function(event) {
    event.preventDefault();

    let formData = new FormData(this);
    let data = {
        warehousename: formData.get("warehouseName"),
        warehouselocation: formData.get("warehouseLocation"),
        warehousedescription: formData.get("warehouseDescription")
    };

    console.log("发送的数据: ", data); // 调试信息，确保数据正确

    fetch('http://localhost:8088/item/createware', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token') // 假设 token 存储在 localStorage 中
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            console.log("响应数据: ", data); // 调试信息，确保响应数据正确
            if (data.code === 200) {
                alert('Warehouse added successfully!');
                //window.location.href = 'adminPanel.html'; // 成功后跳转到管理员面板
            } else {
                alert('Failed to add warehouse: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Fetch Error:', error);
            alert('Error adding warehouse: ' + error.message);
        });
});

function cancelAdd() {
    window.location.href = 'admin_panel.html'; // 取消操作后跳转到管理员面板
}