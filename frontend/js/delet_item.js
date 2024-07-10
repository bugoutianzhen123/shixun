function deleteItem() {
    let itemId = document.getElementById('itemIdToDelete').value;

    fetch('http://localhost:8088/item/delete', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({ id: parseInt(itemId, 10) })
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                console.log('Item deleted successfully');
                // 可以根据需求处理删除成功的情况，比如跳转到其他页面或刷新列表
            } else {
                console.error('Error:', data.message);
                // 处理删除失败的情况
            }
        })
        .catch(error => console.error('Error:', error));
}