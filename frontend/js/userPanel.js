// Add warehouse form handling
function showAddWarehouseForm() {
    document.getElementById('addWarehouseForm').style.display = 'block';
}

function hideAddWarehouseForm() {
    document.getElementById('addWarehouseForm').style.display = 'none';
}

function addWarehouse() {
    let formData = new FormData();
    formData.append('warehouseName', document.getElementById('warehouseName').value);
    formData.append('warehouseLocation', document.getElementById('warehouseLocation').value);
    formData.append('warehouseDescription', document.getElementById('warehouseDescription').value);

    fetch('http://localhost:8088/item/createware', {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Warehouse added successfully!');
                // Reload warehouse list or update UI as needed
            } else {
                alert('Failed to add warehouse.');
            }
            hideAddWarehouseForm();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error adding warehouse.');
            hideAddWarehouseForm();
        });
}

// Add item form handling
function showAddItemForm() {
    document.getElementById('addItemForm').style.display = 'block';
}

function hideAddItemForm() {
    document.getElementById('addItemForm').style.display = 'none';
}

function addItem() {
    let formData = new FormData();
    formData.append('itemName', document.getElementById('itemName').value);
    formData.append('itemDescription', document.getElementById('itemDescription').value);
    formData.append('itemTotalNumber', document.getElementById('itemTotalNumber').value);

    fetch('http://localhost:8088/item/createitem', {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Item added successfully!');
                // Reload item list or update UI as needed
            } else {
                alert('Failed to add item.');
            }
            hideAddItemForm();
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error adding item.');
            hideAddItemForm();
        });
}
