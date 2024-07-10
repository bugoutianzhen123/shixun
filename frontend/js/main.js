document.addEventListener('DOMContentLoaded', function() {
    const apiUrl = 'http://localhost:8088'; // 请替换为您的后端API地址

    // 登录表单提交事件
    document.getElementById('loginForm').addEventListener('submit', function(event) {
        event.preventDefault();

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        fetch(`${apiUrl}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
            .then(response => response.json())
            .then(data => {
                if (data.code === 200) {
                    alert('登录成功！');
                    // 可以根据需要进行页面重定向或其他操作
                    showUserDashboard(username); // 显示用户仪表板
                } else {
                    alert(`登录失败：${data.message}`);
                }
            })
            .catch(error => console.error('登录请求失败：', error));
    });

    // 注册表单提交事件
    document.getElementById('registerForm').addEventListener('submit', function(event) {
        event.preventDefault();

        const username = document.getElementById('regUsername').value;
        const password = document.getElementById('regPassword').value;

        fetch(`${apiUrl}/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
            .then(response => response.json())
            .then(data => {
                if (data.code === 200) {
                    alert('注册成功！');
                    // 可以根据需要进行页面重定向或其他操作
                } else {
                    alert(`注册失败：${data.message}`);
                }
            })
            .catch(error => console.error('注册请求失败：', error));
    });

    // 创建仓库表单提交事件
    document.getElementById('warehouseForm').addEventListener('submit', function(event) {
        event.preventDefault();

        const warehouseName = document.getElementById('warehouseName').value;

        fetch(`${apiUrl}/createWarehouse`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name: warehouseName })
        })
            .then(response => response.json())
            .then(data => {
                if (data.code === 200) {
                    alert('创建仓库成功！');
                    // 可以根据需要进行页面重定向或其他操作
                    fetchWarehouses(); // 重新获取并显示仓库列表
                } else {
                    alert(`创建仓库失败：${data.message}`);
                }
            })
            .catch(error => console.error('创建仓库请求失败：', error));
    });

    // 创建物品表单提交事件
    document.getElementById('itemForm').addEventListener('submit', function(event) {
        event.preventDefault();

        const itemName = document.getElementById('itemName').value;
        const itemQuantity = document.getElementById('itemQuantity').value;

        fetch(`${apiUrl}/createItem`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name: itemName, quantity: itemQuantity })
        })
            .then(response => response.json())
            .then(data => {
                if (data.code === 200) {
                    alert('创建物品成功！');
                    // 可以根据需要进行页面重定向或其他操作
                    fetchItems(); // 重新获取并显示物品列表
                } else {
                    alert(`创建物品失败：${data.message}`);
                }
            })
            .catch(error => console.error('创建物品请求失败：', error));
    });

    // 获取并显示仓库列表
    function fetchWarehouses() {
        fetch(`${apiUrl}/warehouses`)
            .then(response => response.json())
            .then(data => {
                const warehouseList = document.getElementById('warehouseList');
                warehouseList.innerHTML = '';
                data.forEach(warehouse => {
                    const li = document.createElement('li');
                    li.textContent = warehouse.name;
                    warehouseList.appendChild(li);
                });
            })
            .catch(error => console.error('获取仓库列表失败：', error));
    }

    // 获取并显示物品列表
    function fetchItems() {
        fetch(`${apiUrl}/items`)
            .then(response => response.json())
            .then(data => {
                const itemList = document.getElementById('itemList');
                itemList.innerHTML = '';
                data.forEach(item => {
                    const li = document.createElement('li');
                    li.textContent = `${item.name} - 数量: ${item.quantity}`;
                    itemList.appendChild(li);
                });
            })
            .catch(error => console.error('获取物品列表失败：', error));
    }

    // 根据用户角色显示相应的仪表板
    function showUserDashboard(username) {
        // 模拟用户角色（可以根据实际情况调整）
        const isAdmin = username === 'admin'; // 假设用户名为 admin 的是管理员

        // 根据角色显示不同的表单和操作
        if (isAdmin) {
            document.getElementById('warehouse-form').style.display = 'block';
            document.getElementById('item-form').style.display = 'block';
        } else {
            document.getElementById('warehouse-form').style.display = 'none';
            document.getElementById('item-form').style.display = 'none';
        }

        // 显示仓库和物品列表
        fetchWarehouses();
        fetchItems();
    }
});
