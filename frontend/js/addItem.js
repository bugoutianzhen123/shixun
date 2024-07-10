document.getElementById("addItemForm").addEventListener("submit", function(event) {
    event.preventDefault();

    let formData = new FormData(this);
    let data = {
        itemname: formData.get("itemName"),
        totalnumber: formData.get("itemTotalNumber"),
        itemdescription: formData.get("itemDescription")
    };

    console.log('发送的数据:', data); // 添加调试信息

    fetch('http://localhost:8088/item/createitem', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token') // 添加token
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            console.log('服务器响应:', data); // 添加调试信息
            if (data.success) {
                alert('Item added successfully!');
                window.location.href = 'adminPanel.html';
            } else {
                alert('Failed to add item: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error adding item.');
        });
});

function cancelAdd() {
    window.location.href = 'adminPanel.html';
}
