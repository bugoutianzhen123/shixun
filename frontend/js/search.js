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
            <td>${record.innumber || ''}</td>
            <td>${record.CreatedAt || ''}</td>
            <td>${record.UpdatedAt || ''}</td>
            <td>${record.Item ? record.Item.itemname || '' : ''}</td>
            <td>${record.Item ? record.Item.CreatedAt || '' : ''}</td>
            <td>${record.Item ? record.Item.UpdatedAt || '' : ''}</td>
            <td>${record.Item ? record.Item.itemdescription || '' : ''}</td>
            <td>${record.warehousename || ''}</td>
            <td>${record.warehouselocation || ''}</td>
            <td>${record.warehousedescription || ''}</td>
            <td>${record.warehouse ? record.warehouse.CreatedAt || '' : ''}</td>
            <td>${record.warehouse ? record.warehouse.UpdatedAt || '' : ''}</td>
        `;
        tableBody.appendChild(row);
    });
}

function displayinnerResults(containerId, results) {
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
            <td>${record.innumber || ''}</td>
            <td>${record.CreatedAt || ''}</td>
            <td>${record.UpdatedAt || ''}</td>
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
    let searchType = document.querySelector('input[name="searchType"]:checked').value;
    let searchValue;
    let requestBody;

    if (searchType === "id") {
        searchValue = document.getElementById('itemId').value;
        requestBody = { itemid: parseInt(searchValue, 10) };
    } else {
        searchValue = document.getElementById('itemName').value;
        requestBody = { itemname: searchValue };
    }

    fetch('http://localhost:8088/item/getitem', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify(requestBody)
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
                let filteredData = data.data.map(record => ({
                    ID: record.ID,
                    warehouseid: record.warehouseid,
                    itemid: record.itemid,
                    Number: record.Number,
                    CreatedAt: record.CreatedAt,
                    UpdatedAt: record.UpdatedAt
                }));
                displayResults('inventoryResults', filteredData);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}

function displaywarehouseResults(containerId, results) {
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
            <td>${record.warehousename || ''}</td>
            <td>${record.warehouselocation || ''}</td>
            <td>${record.warehousedescription || ''}</td>
            <td>${record.CreatedAt || ''}</td>
            <td>${record.UpdatedAt || ''}</td>
        `;
        tableBody.appendChild(row);
    });
}

function findWarehouse() {
    let searchType = document.querySelector('input[name="searchType"]:checked').value;
    let searchValue;
    let requestBody;

    if (searchType === "id") {
        searchValue = document.getElementById('warehouseId').value;
        requestBody = { warehouseid: parseInt(searchValue, 10) };
    } else {
        searchValue = document.getElementById('warehouseName').value;
        requestBody = { warehousename: searchValue };
    }

    fetch('http://localhost:8088/item/getware', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify(requestBody)
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                displaywarehouseResults('warehouseResults', data.data);
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
                let filteredData = data.data.map(record => ({
                    ID: record.ID,
                    warehouseid: record.warehouseid,
                    itemid: record.itemid,
                    innumber: record.innumber,
                    CreatedAt: record.CreatedAt,
                    UpdatedAt: record.UpdatedAt
                }));
                displayinnerResults('inboundResults', filteredData);
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
                let filteredData = data.data.map(record => ({
                    ID: record.ID,
                    warehouseid: record.warehouseid,
                    itemid: record.itemid,
                    Number: record.Number,
                    CreatedAt: record.CreatedAt,
                    UpdatedAt: record.UpdatedAt
                }));
                displayResults('outboundResults', filteredData);
            } else {
                console.error('Error:', data.message);
            }
        })
        .catch(error => console.error('Error:', error));
}