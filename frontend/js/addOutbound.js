// 添加出库记录的表单提交事件监听器
document.getElementById("addOutboundForm").addEventListener("submit", function(event) {
    event.preventDefault(); // 阻止表单默认提交行为

    const formData = new FormData(this);
    const data = {
        warehouseid: parseInt(formData.get("warehouseid"), 10),
        itemid: parseInt(formData.get("itemid"), 10),
        outnumber: parseInt(formData.get("outnumber"), 10)
    };

    fetch('http://localhost:8088/item/createoutb', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                alert('出库记录添加成功！');
                window.location.href = 'adminPanel.html'; // 成功后跳转到管理员面板
            } else {
                alert('添加出库记录失败: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('添加出库记录失败.');
        });
});

