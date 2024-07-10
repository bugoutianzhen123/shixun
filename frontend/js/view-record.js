async function loadRecords() {
    const response = await fetch(`${baseUrl}/record/list`, {
        headers: {
            'Authorization': getToken()
        }
    });

    const records = await response.json();
    const recordList = document.getElementById('record-list');
    records.forEach(record => {
        const li = document.createElement('li');
        li.textContent = `Type: ${record.type}, Item ID: ${record.item_id}, Warehouse ID: ${record.warehouse_id}, Quantity: ${record.quantity}`;
        recordList.appendChild(li);
    });
}

loadRecords();
