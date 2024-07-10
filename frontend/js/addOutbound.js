// 监听表单提交事件
document.getElementById("searchForm").addEventListener("submit", function(event) {
    event.preventDefault(); // 阻止表单默认提交行为

    const formData = new FormData(this);
    const warehouse = formData.get('warehouse').trim();
    const item = formData.get('item').trim();
    const resultsDiv = document.getElementById('results');

    // 清空结果区域
    resultsDiv.innerHTML = '';

    if (warehouse || item) { // 使用或运算符，以避免两者均为空时显示错误
        if (warehouse && item) {
            // 查询指定仓库中的指定物品
            fetchWarehouseItemData(warehouse, item)
                .then(data => {
                    displayResults(data);
                })
                .catch(error => {
                    resultsDiv.innerHTML = `<p>查询出现错误: ${error.message}</p>`;
                });
        } else if (warehouse) {
            // 查询指定仓库中的所有物品
            fetchWarehouseData(warehouse)
                .then(data => {
                    displayResults(data);
                })
                .catch(error => {
                    resultsDiv.innerHTML = `<p>查询出现错误: ${error.message}</p>`;
                });
        } else if (item) {
            // 查询指定物品所在的所有仓库
            fetchItemData(item)
                .then(data => {
                    displayResults(data);
                })
                .catch(error => {
                    resultsDiv.innerHTML = `<p>查询出现错误: ${error.message}</p>`;
                });
        }
    } else {
        resultsDiv.innerHTML = '<p>请至少输入一个查询条件</p>';
    }
});

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

// 取消添加的函数
function cancelAdd() {
    window.location.href = 'adminPanel.html'; // 取消后跳转到管理员面板
}

function fetchWarehouseData(warehouse) {
    return fetch('http://localhost:8088/item/getware', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ warehouse })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('网络响应异常');
            }
            return response.json();
        })
        .catch(error => {
            console.error('Fetch 错误:', error);
            throw error;
        });
}

function fetchItemData(item) {
    return fetch('http://localhost:8088/item/getitem', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ item })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('网络响应异常');
            }
            return response.json();
        })
        .catch(error => {
            console.error('Fetch 错误:', error);
            throw error;
        });
}

function fetchWarehouseItemData(warehouse, item) {
    return fetch('http://localhost:8088/item/getinven', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ warehouse, item })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('网络响应异常');
            }
            return response.json();
        })
        .catch(error => {
            console.error('Fetch 错误:', error);
            throw error;
        });
}

function displayResults(data) {
    const resultsDiv = document.getElementById('results');
    if (data.length > 0) {
        if (data[0].warehouse && data[0].item) {
            // 同时包含仓库和物品信息，显示仓库标题
            data.forEach(item => {
                resultsDiv.innerHTML += `<h4>仓库: ${item.warehouse}  物品: ${item.item}  数量: ${item.quantity}   描述: ${item.description}</h4>`;
            });
        } else {
            // 其他情况处理
            data.forEach(item => {
                resultsDiv.innerHTML += `<p>${JSON.stringify(item)}</p>`;
            });
        }
    } else {
        resultsDiv.innerHTML = '<p>没有找到相关数据</p>';
    }
}
