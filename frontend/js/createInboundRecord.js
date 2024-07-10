document.getElementById("inboundRecordForm").addEventListener("submit", function(event) {
    event.preventDefault();

    let formData = new FormData(this);
    let data = {
        itemId: formData.get("itemId"),
        quantity: formData.get("quantity"),
        warehouseId: formData.get("warehouseId")
    };

    fetch('http://localhost:8088/item/createinb', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('Inbound record created successfully!');
                window.location.href = 'admin_panel.html';
            } else {
                alert('Failed to create inbound record: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error creating inbound record.');
        });
});
