function displayResults(containerId, results) {
    let container = document.getElementById(containerId);
    if (!container) {
        console.error(`Element with id ${containerId} not found`);
        return;
    }

    let tableBody = container.querySelector('tbody');
    if (!tableBody) {
        console.error('Table body element not found');
        return;
    }

    tableBody.innerHTML = ''; // Clear previous results

    results.forEach(record => {
        let row = document.createElement('tr');
        row.innerHTML = `
            <td>${record.ID || ''}</td>
            <td>${record.warehouseid || ''}</td>
            <td>${record.itemid || ''}</td>
            <td>${record.Number || ''}</td>
            <td>${record.CreatedAt || ''}</td>
            <td>${record.UpdatedAt || ''}</td>
            <td>${record.Item ? record.Item.itemname || '' : ''}</td>
            <td>${record.Item ? record.Item.CreatedAt || '' : ''}</td>
            <td>${record.Item ? record.Item.UpdatedAt || '' : ''}</td>
            <td>${record.Item ? record.Item.itemdescription || '' : ''}</td>
            <td>${record.Warehouse ? record.Warehouse.warehousename || '' : ''}</td>
            <td>${record.Warehouse ? record.Warehouse.warehouselocation || '' : ''}</td>
            <td>${record.Warehouse ? record.Warehouse.warehousedescription || '' : ''}</td>
            <td>${record.Warehouse ? record.Warehouse.CreatedAt || '' : ''}</td>
            <td>${record.Warehouse ? record.Warehouse.UpdatedAt || '' : ''}</td>
        `;
        tableBody.appendChild(row);
    });
}

function displayItemResults(containerId, results) {
    let container = document.getElementById(containerId);
    if (!container) {
        console.error(`Element with id ${containerId} not found`);
        return;
    }

    let tableBody = container.querySelector('tbody');
    if (!tableBody) {
        console.error('Table body element not found');
        return;
    }

    tableBody.innerHTML = ''; // Clear previous results

    results.forEach(record => {
        let row = document.createElement('tr');
        row.innerHTML = `
            <td>${record.ID || ''}</td>
            <td>${record.itemname || ''}</td>
            <td>${record.TotalNumber || ''}</td>
            <td>${record.itemdescription || ''}</td>
            <td>${record.CreatedAt || ''}</td>
            <td>${record.UpdatedAt || ''}</td>
        `;
        tableBody.appendChild(row);
    });
}
function findItem() {
    let itemId = document.getElementById('itemId').value;

    fetch('http://localhost:8088/item/getitem', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ itemid: parseInt(itemId, 10) })
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                console.log(data.data);  // 添加此行来调试数据结构
                displayItemResults('itemResults', data.data);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}
function findInventory() {
    let warehouseId = document.getElementById('inventoryWarehouseId').value;
    let itemId = document.getElementById('inventoryItemId').value;

    fetch('http://localhost:8088/item/getinven', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ warehouseid: parseInt(warehouseId, 10), itemid: parseInt(itemId, 10) })
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                displayResults('inventoryResults', data.data);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}

function findWarehouse() {
    let warehouseId = document.getElementById('warehouseId').value;

    fetch('http://localhost:8088/item/getware', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ warehouseid: parseInt(warehouseId, 10) }) // 确保字段名为 warehouseid
    })
        .then(response => response.json())
        .then(data => {
            console.log('Response data:', data);
            if (data.code === 200) {
                displayResults('warehouseResults', data.data);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}

function findInboundRecords() {
    let warehouseId = document.getElementById('inboundWarehouseId').value;
    let itemId = document.getElementById('inboundItemId').value;

    fetch('http://localhost:8088/item/getinb', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ warehouseid: parseInt(warehouseId, 10), itemid: parseInt(itemId, 10) })
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                displayResults('inboundResults', data.data);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}

function findOutboundRecords() {
    let warehouseId = document.getElementById('outboundWarehouseId').value;
    let itemId = document.getElementById('outboundItemId').value;

    fetch('http://localhost:8088/item/getoutb', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify({ warehouseid: parseInt(warehouseId, 10), itemid: parseInt(itemId, 10) })
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                displayResults('outboundResults', data.data);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}
